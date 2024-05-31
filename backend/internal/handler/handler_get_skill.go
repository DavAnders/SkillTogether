package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetSkill(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	skill, err := h.Queries.GetSkill(c.Request.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get skill"})
		}
		return
	}

	c.JSON(http.StatusOK, skill)
}