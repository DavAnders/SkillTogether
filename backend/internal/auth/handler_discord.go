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
	"net/http"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/gin-gonic/gin"
)

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

    if err := processDiscordUser(ctx, h.Queries, user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    loginSessionHandler(c, user)
}


func getDiscordUser(data []byte) (*DiscordUser, error) {
	var user DiscordUser
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func processDiscordUser(ctx context.Context, q *db.Queries, user *DiscordUser) error {
	avatarURL := user.AvatarURL()
    // Check if the user exists
    _, err := q.GetUser(ctx, user.ID)
    if err != nil {
        if err == sql.ErrNoRows {
            // Create new user
            _, err := q.AddUser(ctx, db.AddUserParams{
                DiscordID: user.ID,
                Username:  user.Username,
                Email:     user.Email,
                AvatarUrl: sql.NullString{String: avatarURL, Valid: avatarURL != ""},
            })
            if err != nil {
                return err
            }
            return nil
        }
        return err
    }

    // Update user information if not exists
    err = q.UpdateUser(ctx, db.UpdateUserParams{
        DiscordID: user.ID,
        Username:  user.Username,
        Email:     user.Email,
        AvatarUrl: sql.NullString{String: avatarURL, Valid: avatarURL != ""},
    })
    return err
}


func loginSessionHandler(c *gin.Context, user *DiscordUser) {
    // Generate a session token
	sessionToken, err := generateSessionToken(user.ID)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session token"})
        return
    }

	c.SetCookie("session_token", sessionToken, 86400, "/", "", true, true)

    // Redirect to dashboard
	c.Redirect(http.StatusSeeOther, "/dashboard")
}

func generateSessionToken(userID string) (string, error) {
    // Generate random bytes
    randomBytes := make([]byte, 32)
    if _, err := rand.Read(randomBytes); err != nil {
        return "", err
    }

    // Combine user ID with random bytes
    tokenData := fmt.Sprintf("%s:%x", userID, randomBytes)

    // Hash the combined data to get the token
    hash := sha256.Sum256([]byte(tokenData))

    // Return the base64 URL encoded hash
    return base64.URLEncoding.EncodeToString(hash[:]), nil
}