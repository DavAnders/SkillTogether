package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteInterest(c *gin.Context) {
    interestID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interest ID format"})
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

    // Fetch the interest to check the owner before deletion
    interest, err := h.Queries.GetInterest(c.Request.Context(), int32(interestID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve interest"})
        return
    }

    // Check if the interest's user ID matches the logged-in user ID
    if interest.UserID != user.ID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this interest"})
        return
    }

    // Delete if authorized
    if err := h.Queries.DeleteInterest(c.Request.Context(), int32(interestID)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete interest"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"success": "Interest deleted successfully"})
}