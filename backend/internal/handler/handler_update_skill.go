package handler

import (
	"net/http"
	"strconv"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateSkill(c *gin.Context) {
	var req db.UpdateSkillParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	req.ID = int32(id)

	err = h.Queries.UpdateSkill(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update skill"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Skill updated successfully"})
}