package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *EventMemberHandler) SaveEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	err = h.Service.SaveEvent(uint(id), c.GetUint("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event saved"})
}

func (h *EventMemberHandler) UnsaveEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	err = h.Service.UnsaveEvent(uint(id), c.GetUint("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to unsave event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event unsaved"})
}

func (h *EventMemberHandler) GetSavedEvents(c *gin.Context) {
	events, err := h.Service.GetSavedEvents(c.GetUint("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get saved events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}
