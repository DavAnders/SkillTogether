package auth

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(q *db.Queries) gin.HandlerFunc {
    return func(c *gin.Context) {
        token, err := c.Cookie("session_token")
        log.Printf("Retrieved session token: %s", token)
        if err != nil {
            log.Printf("No session token provided: %v", err)
            c.JSON(401, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        valid, discordID, err := validateToken(c.Request.Context(), q, token)
        if err != nil || !valid {
            log.Printf("Invalid session token: %v, %v", err, token)
            c.JSON(401, gin.H{"error": "Unauthorized: Invalid session token"})
            c.Abort()
            return
        }

        // Set the Discord ID in the Gin context, not the split ID from token
        c.Set("discord_id", discordID)

        c.Next()
    }
}

func validateToken(ctx context.Context, q *db.Queries, token string) (bool, string, error) {
    // Split the token to separate the Discord ID and hash part
    parts := strings.Split(token, ":")
    if len(parts) != 2 {
        log.Printf("Invalid token format: %s", token)
        return false, "", fmt.Errorf("invalid token format")
    }
    
    tokenHash := parts[1]  // This is what we store and validate against

    discordID, err := q.GetUserIDFromSessionToken(ctx, tokenHash)
    if err != nil {
        log.Printf("Failed to retrieve Discord ID from session token: %v", err)
        return false, "", fmt.Errorf("failed to validate session token: %v", err)
    }

    if discordID == "" {
        log.Printf("No Discord ID associated with the provided token hash")
        return false, "", fmt.Errorf("token validation failed")
    }

    log.Printf("Token successfully validated for Discord ID: %s", discordID)
    return true, discordID, nil
}