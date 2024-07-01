package handler

import (
	"net/http"
	"strconv"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/gin-gonic/gin"
)

// UpdateInterest updates an interest's information based on the provided JSON request.
func (h *Handler) UpdateInterest(c *gin.Context) {
	var req db.UpdateInterestParams
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

	err = h.Queries.UpdateInterest(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update interest"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Interest updated successfully"})
}
