package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetInterest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	interest, err := h.Queries.GetInterest(c.Request.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Interest not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get interest"})
		}
		return
	}

	c.JSON(http.StatusOK, interest)
}