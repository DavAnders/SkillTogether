package handler

import (
	"log"
	"net/http"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/gin-gonic/gin"
)

// AddInterest adds an interest to the database.
func (h *Handler) AddInterest(c *gin.Context) {
	var req struct {
		Interest string `json:"interest"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the Discord ID from the context
	discordID, exists := c.Get("discord_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User authentication failed"})
		return
	}

	// Use the Discord ID to get the actual user ID from the database
	userID, err := h.Queries.GetUserIDByDiscordID(c.Request.Context(), discordID.(string))
	if err != nil {
		log.Printf("Error retrieving user ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user ID"})
		return
	}

	// Now, use this userID to insert the interest into the database
	interestID, err := h.Queries.AddInterest(c.Request.Context(), db.AddInterestParams{
		UserID:   userID,
		Interest: req.Interest,
	})
	if err != nil {
		log.Printf("Error adding interest: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add interest"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"interest_id": interestID})
}
