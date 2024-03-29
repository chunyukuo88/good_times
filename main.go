package main

import (
	"fmt"
	"net/http"

	"github.com/chunyukuo88/good_times/db"
	"github.com/chunyukuo88/good_times/models"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("About to create tables...")
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
	fmt.Println("Events obtained.")
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
