package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllInterests retrieves all interests.
func (h *Handler) GetAllInterests(c *gin.Context) {
	interests, err := h.Queries.GetAllInterests(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all interests"})
		return
	}
	c.JSON(http.StatusOK, interests)
}
