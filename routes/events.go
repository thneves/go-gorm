package routes

import (
	"fmt"
	"go-gorm/models"
	event "go-gorm/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getEvents(context *gin.Context, db *gorm.DB) {
	//events, err := event.GetAllEvents(db)
	events, err := event.GetAllEvents(db)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve events",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "All Events Retrieved",
		"events":  events,
	})
}

func createEvent(context *gin.Context, db *gorm.DB) {
	var event event.Event

	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	event, err := event.Save(db)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save event",
		})

		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event Created",
		"event":   event,
	})
}

func getEvent(context *gin.Context, db *gorm.DB) {
	id := context.Param("id")
	parsedId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Error parsing string")
	}

	event, err := event.GetEventById(parsedId, db)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to retrieve evenbt",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "succesfully retrieved event",
		"event":   event,
	})
}

func deleteEvent(context *gin.Context, db *gorm.DB) {
	id := context.Param("id")
	parsedId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Error parsing string")
	}

	deleted_event, err := event.Delete(parsedId, db)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to delete!",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "event succesfully deleted!",
		"event":   deleted_event,
	})
}

func updateEvent(context *gin.Context, db *gorm.DB) {
	// Extract the event ID from the URL params
	id := context.Param("id")
	parsedID, _ := strconv.Atoi(id)

	// Find the existing event by its ID
	event, err := models.GetEventById(parsedID, db)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Bind the updated data from the JSON request body to the event struct (without modifying the ID)
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the update function to update the record in the database
	updatedEvent, err := event.Update(db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	// Return the updated event as JSON
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent})
}
