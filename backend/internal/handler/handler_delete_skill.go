package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteSkill(c *gin.Context) {
    skillID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID format"})
        return
    }

    // Retrieve the Discord ID from the session
    discordID, exists := c.Get("discord_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User authentication failed"})
        return
    }

	// Expect the Discord ID to be a string
    discordIDStr, ok := discordID.(string)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Discord ID format is invalid"})
        return
    }

    // Fetch the user's ID from the Discord ID
    user, err := h.Queries.GetUserByDiscordID(c.Request.Context(), discordIDStr)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user information"})
        return
    }

    // Fetch the skill to check the owner before deletion
    skill, err := h.Queries.GetSkill(c.Request.Context(), int32(skillID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve skill"})
        return
    }

    // Check if the skill's user ID matches the logged-in user ID
    if skill.UserID.Valid && skill.UserID.Int32 != user.ID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this skill"})
        return
    }

    // Delete if authorized
    if err := h.Queries.DeleteSkill(c.Request.Context(), int32(skillID)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete skill"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"success": "Skill deleted successfully"})
}