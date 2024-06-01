package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) GetUserFromSession(c *gin.Context) {
    // Retrieve session token from cookie
    sessionToken, err := c.Cookie("session_token")
    if err != nil {
        log.Printf("No session token provided: %v", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "No session token provided"})
        return
    }
    log.Printf("Retrieved session token: %s", sessionToken)

    parts := strings.Split(sessionToken, ":")
    if len(parts) != 2 {
        log.Printf("Invalid session token format: %s", sessionToken)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session token format"})
        return
    }
    tokenHash := parts[1]

    // Retrieve Discord ID using the tokenHash part of the session token
    discordID, err := h.Queries.GetUserIDFromSessionToken(c.Request.Context(), tokenHash)
    if err != nil {
        log.Printf("Failed to retrieve Discord ID from session token hash: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Discord ID from session token"})
        return
    }
    log.Printf("Retrieved Discord ID: %s", discordID)

    // Retrieve user details using Discord ID
    user, err := h.Queries.GetUserByDiscordID(c.Request.Context(), discordID)
    if err != nil {
        log.Printf("Failed to retrieve user details for Discord ID %s: %v", discordID, err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details"})
        return
    }
    log.Printf("Retrieved user details: %+v", user)

    c.JSON(http.StatusOK, user)
}


