package handler

import (
	"net/http"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddInterestByBot(c *gin.Context) {
	var input struct {
		DiscordID string `json:"discord_id"`
		Interest  string `json:"interest"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: " + err.Error()})
		return
	}

	userID, err := h.Queries.GetUserIDByDiscordID(c, input.DiscordID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user ID: " + err.Error()})
		return
	}

	arg := db.AddInterestParams{
		UserID:   userID,
		Interest: input.Interest,
	}
	_, err = h.Queries.AddInterest(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add interest: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Interest added successfully"})
}