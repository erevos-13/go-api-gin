package routes

import (
	"net/http"
	"strconv"

	"example.com/gin-api/models"
	"github.com/gin-gonic/gin"
)

func GetEvent(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func GetEventById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	eventId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"event": event,
	})

}

func CreateEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	newEvent := models.Event{}
	err := context.ShouldBindBodyWithJSON(&newEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEvent.UserId = userId
	err = newEvent.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"status": "created",
	})

}

func UpdateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "you not eligible to update this event"})
		return
	}
	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateEvent.ID = id
	err = updateEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": "updated",
	})
}

func DeleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "you not eligible to update this event"})
		return
	}
	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": "deleted",
	})

}
