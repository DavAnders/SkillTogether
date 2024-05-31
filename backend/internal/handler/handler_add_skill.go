package handler

import (
	"log"
	"net/http"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/gin-gonic/gin"
)

func NewHandler (queries *db.Queries) *Handler {
	return &Handler{
		Queries: queries,
	}
}

type Handler struct {
	Queries *db.Queries
}

func (h *Handler) AddSkill(c *gin.Context) {
	var req db.AddSkillParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.Queries.AddSkill(c, req)
	if err != nil {
		log.Printf("Error adding skill: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add skill"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}