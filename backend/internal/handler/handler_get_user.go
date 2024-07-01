package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User represents a user in the database, omitting some fields.
type User struct {
	ID        int    `json:"id"`
	DiscordID string `json:"discord_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// GetUser handles the retrieval of a user based on their Discord ID.
func (h *Handler) GetUser(c *gin.Context) {
	discordID := c.Param("discord_id")
	if discordID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid discord id"})
		return
	}

	user, err := h.Queries.GetUser(c, discordID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
