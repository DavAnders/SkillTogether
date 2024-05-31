package handler

import (
	"net/http"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/gin-gonic/gin"
)

// This is temporary, create user should be done through discord oauth
func (h *Handler) AddUser(c *gin.Context) {
	var req db.AddUserParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data," + err.Error()})
		return
	}

	id, err := h.Queries.AddUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}