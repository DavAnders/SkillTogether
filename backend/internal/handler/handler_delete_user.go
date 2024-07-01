package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteUser deletes a user based on the provided Discord ID.
func (h *Handler) DeleteUser(c *gin.Context) {
	user := c.Param("discord_id")
	if err := h.Queries.DeleteUser(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
