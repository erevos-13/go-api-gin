package main

import (
	"net/http"

	"example.com/gin-api/db"
	"example.com/gin-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Code
	db.InitDB()
	server := gin.Default()

	server.GET("/event", getEvent)
	server.POST("/event", createEvent)

	server.Run(":3000")

}

func getEvent(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	newEvent := models.Event{}
	err := context.ShouldBindBodyWithJSON(&newEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newEvent.ID = 1
	newEvent.UserId = 55
	err = newEvent.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"status": "created",
	})

}
