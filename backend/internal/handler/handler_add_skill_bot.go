package handler

import (
	"database/sql"
	"net/http"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/gin-gonic/gin"
)

// AddSkillByBot handles the addition of a skill by the Discord bot.
func (h *Handler) AddSkillByBot(c *gin.Context) {
	var input struct {
		DiscordID        string `json:"discord_id"`
		SkillDescription string `json:"skill_description"`
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

	arg := db.AddSkillParams{
		UserID:           sql.NullInt32{Int32: userID, Valid: true}, // Valid since we got a valid response
		SkillDescription: input.SkillDescription,
	}
	_, err = h.Queries.AddSkill(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add skill: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Skill added successfully"})
}
