package main

import (
	"fmt"
	"go-gorm/database"
	event "go-gorm/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db_conn, err := database.NewConnection(config)

	if err != nil {
		log.Fatal(err)
	}

	DB, err := db_conn.DB()

	DB.SetConnMaxIdleTime(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatal(err)
	}

	server := gin.Default()

	//	server.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//	"message": "pong",
	// })
	// })

	server.GET("/events", func(c *gin.Context) {
		getEvents(c, db_conn)
	})
	server.POST("/create_event", func(c *gin.Context) {
		createEvent(c, db_conn)
	})
	server.GET("/event/:id", func(c *gin.Context) {
		getEvent(c, db_conn)
	})

	server.Run()
}

func getEvents(context *gin.Context, db *gorm.DB) {
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
