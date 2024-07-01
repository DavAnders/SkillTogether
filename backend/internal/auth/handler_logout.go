package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogoutHandler handles the logout of a user.
func (h *AuthHandler) LogoutHandler(c *gin.Context) {
	sessionToken, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No session token found"})
		return
	}

	err = h.Queries.DeleteSession(c.Request.Context(), sessionToken)
	if err != nil {
		log.Printf("Failed to delete session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log out"})
		return
	}

	c.SetCookie("session_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
