package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/gin-gonic/gin"
)

// UserSimple represents a simplified user structure with basic information.
type UserSimple struct {
	ID        int32  `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	DiscordID string `json:"discord_id"`
}

// EnhancedSkill represents a skill along with associated user information.
type EnhancedSkill struct {
	Skill     db.SimpleSearchSkillRow `json:"skill"`
	User      UserSimple              `json:"user"`
	CreatedAt time.Time               `json:"created_at"`
}

// SimpleSearchSkill performs a simple search for skills based on a query parameter.
// An ILIKE search is used.
func (h *Handler) SimpleSearchSkill(c *gin.Context) {
	searchQuery := c.Query("q")
	if searchQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	// Convert the string to sql.NullString
	nullSearchQuery := sql.NullString{String: searchQuery, Valid: true}

	skills, err := h.Queries.SimpleSearchSkill(c, nullSearchQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, skills)
}

// SearchSkillsWithUserInfo performs a search for skills based on a query parameter.
// An ILIKE search is used. This function also retrieves user information for each skill.
func (h *Handler) SearchSkillsWithUserInfo(c *gin.Context) {
	searchQuery := c.Query("q")
	if searchQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	nullSearchQuery := sql.NullString{String: searchQuery, Valid: true}
	skills, err := h.Queries.SimpleSearchSkill(c, nullSearchQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed: " + err.Error()})
		return
	}

	enhancedSkills := make([]EnhancedSkill, 0, len(skills))
	for _, skill := range skills {
		userID := skill.UserID.Int32

		user, err := h.Queries.GetUserByID(c, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
			return
		}

		avatarURL := ""
		if user.AvatarUrl.Valid {
			avatarURL = user.AvatarUrl.String
		}

		enhancedSkills = append(enhancedSkills, EnhancedSkill{
			Skill: skill,
			User: UserSimple{
				ID:        user.ID,
				Username:  user.Username,
				AvatarURL: avatarURL,
				DiscordID: user.DiscordID,
			},
			CreatedAt: skill.CreatedAt.Time,
		})
	}

	c.JSON(http.StatusOK, enhancedSkills)
}
