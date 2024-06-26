package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/gin-gonic/gin"
)

// EnhancedInterest combines a simple interest row with user information.
type EnhancedInterest struct {
	Interest  db.SimpleSearchInterestRow `json:"interest"`
	User      UserSimple                 `json:"user"`
	CreatedAt time.Time                  `json:"created_at"`
}

// SimpleSearchInterest handles the search for interests based on a query parameter.
// An ILIKE search is used.
func (h *Handler) SimpleSearchInterest(c *gin.Context) {
	searchQuery := c.Query("q")
	if searchQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	// Convert the string to sql.NullString
	nullSearchQuery := sql.NullString{String: searchQuery, Valid: true}

	interests, err := h.Queries.SimpleSearchInterest(c, nullSearchQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, interests)
}

// SearchInterestsWithUserInfo handles the search for interests based on a query parameter.
// An ILIKE search is used. This function also retrieves user information for each interest.
func (h *Handler) SearchInterestsWithUserInfo(c *gin.Context) {
	searchQuery := c.Query("q")
	if searchQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	nullSearchQuery := sql.NullString{String: searchQuery, Valid: true}
	interests, err := h.Queries.SimpleSearchInterest(c, nullSearchQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed: " + err.Error()})
		return
	}

	enhancedInterests := make([]EnhancedInterest, 0, len(interests))
	for _, interest := range interests {
		userID := interest.UserID

		user, err := h.Queries.GetUserByID(c, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
			return
		}

		avatarURL := ""
		if user.AvatarUrl.Valid {
			avatarURL = user.AvatarUrl.String
		}

		enhancedInterests = append(enhancedInterests, EnhancedInterest{
			Interest: interest,
			User: UserSimple{
				ID:        user.ID,
				Username:  user.Username,
				AvatarURL: avatarURL,
				DiscordID: user.DiscordID,
			},
			CreatedAt: interest.CreatedAt.Time,
		})
	}

	c.JSON(http.StatusOK, enhancedInterests)
}
