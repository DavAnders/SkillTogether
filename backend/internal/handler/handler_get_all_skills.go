package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllSkills retrieves all skills.
func (h *Handler) GetAllSkills(c *gin.Context) {
	skills, err := h.Queries.GetAllSkills(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all skills"})
		return
	}
	c.JSON(http.StatusOK, skills)
}
