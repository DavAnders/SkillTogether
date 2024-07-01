package handler

import (
	"net/http"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/gin-gonic/gin"
)

// UpdateUser updates a user's information based on the provided JSON request.
func (h *Handler) UpdateUser(c *gin.Context) {
	var req db.UpdateUserParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	req.DiscordID = c.Param("discord_id")

	if err := h.Queries.UpdateUser(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
