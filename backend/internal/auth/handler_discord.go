package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/gin-gonic/gin"
)

// DiscordCallbackHandler handles the OAuth2 callback from Discord.
func (h *AuthHandler) DiscordCallbackHandler(c *gin.Context) {
	ctx := c.Request.Context()
	code := c.Query("code")
	token, err := DiscordOAuth2Config.Exchange(ctx, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token," + err.Error()})
		return
	}

	// Use token to get user info
	client := DiscordOAuth2Config.Client(ctx, token)
	resp, err := client.Get("https://discord.com/api/users/@me")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	defer resp.Body.Close()
	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read user info"})
		return
	}

	user, err := getDiscordUser(userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
		return
	}

	sessionToken, tokenHash, err := generateSessionToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session token"})
		return
	}

	if err := processDiscordUser(ctx, h.Queries, user, tokenHash); err != nil {
		log.Printf("processDiscordUser failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.handleLogin(ctx, user, tokenHash); err != nil {
		log.Printf("handleLogin failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to handle login"})
		return
	}

	// Set the session token as a cookie
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("session_token", sessionToken, 86400, "/", "", true, true)

	// Redirect to the dashboard
	frontendDashboardURL := os.Getenv("FRONTEND_URL") + "/dashboard"
	c.Redirect(http.StatusFound, frontendDashboardURL)
}

func getDiscordUser(data []byte) (*DiscordUser, error) {
	var user DiscordUser
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func processDiscordUser(ctx context.Context, q *db.Queries, user *DiscordUser, tokenHash string) error {
	avatarURL := user.AvatarURL()
	_, err := q.GetUser(ctx, user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("User does not exist, creating new user")
			// Create new user with just the hash part of the token
			_, err = q.AddUser(ctx, db.AddUserParams{
				DiscordID:    user.ID,
				Username:     user.Username,
				Email:        user.Email,
				AvatarUrl:    sql.NullString{String: avatarURL, Valid: avatarURL != ""},
				SessionToken: sql.NullString{String: tokenHash, Valid: tokenHash != ""},
			})
			if err != nil {
				log.Printf("Failed to create new user: %v", err)
				return err
			}
		} else {
			log.Printf("Failed to get user: %v", err)
			return err
		}
	} else {
		// Update user information, storing only the hash part
		err = q.UpdateUser(ctx, db.UpdateUserParams{
			DiscordID:    user.ID,
			Username:     user.Username,
			Email:        user.Email,
			AvatarUrl:    sql.NullString{String: avatarURL, Valid: avatarURL != ""},
			SessionToken: sql.NullString{String: tokenHash, Valid: tokenHash != ""},
		})
		if err != nil {
			log.Printf("Failed to update user: %v", err)
			return err
		}
	}

	return nil
}

func generateSessionToken(userID string) (string, string, error) {
	// Generate random bytes
	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", "", err
	}

	// Combine user ID with random bytes
	tokenData := fmt.Sprintf("%s:%x", userID, randomBytes)

	// Hash the combined data to get the token
	hash := sha256.Sum256([]byte(tokenData))

	// Encode token to create a token string, combining with user ID
	tokenHash := base64.URLEncoding.EncodeToString(hash[:])
	finalToken := fmt.Sprintf("%s:%s", userID, tokenHash)

	// Return the final token and the token hash part
	return finalToken, tokenHash, nil
}

func (h *AuthHandler) handleLogin(ctx context.Context, user *DiscordUser, tokenHash string) error {
	expiresAt := time.Now().Add(24 * time.Hour)

	// Store or update the session information in the database
	err := h.Queries.CreateOrUpdateUserSession(ctx, db.CreateOrUpdateUserSessionParams{
		DiscordID:    user.ID,
		SessionToken: tokenHash,
		ExpiresAt:    expiresAt,
	})
	if err != nil {
		log.Printf("Failed to store session token: %v", err)
		return err
	}

	return nil
}
