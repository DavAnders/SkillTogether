package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/gin-gonic/gin"
)

// NewHandler creates a new Handler struct.
func NewHandler(queries *db.Queries) *Handler {
	return &Handler{
		Queries: queries,
	}
}

// Handler is the struct that contains the database queries.
type Handler struct {
	Queries *db.Queries
}

// AddSkill handles the addition of a skill to the database.
func (h *Handler) AddSkill(c *gin.Context) {
	var req struct {
		Description string `json:"description"`
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

	// Now, use this userID to insert the skill into the database
	skillID, err := h.Queries.AddSkill(c.Request.Context(), db.AddSkillParams{
		UserID:           sql.NullInt32{Int32: userID, Valid: true},
		SkillDescription: req.Description,
	})
	if err != nil {
		log.Printf("Error adding skill: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add skill"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"skill_id": skillID})
}
