package main

import (
	"go-gorm/database"
	event "go-gorm/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	server.GET("/events", getEvents)
	server.POST("/create_event", createEvent)

	server.Run()
}

func getEvents(context *gin.Context) {
	events := event.GetAllEvents()
	context.JSON(http.StatusOK, gin.H{
		"message": "All Events Retrieved",
		"events":  events,
	})
}

func createEvent(context *gin.Context) {
	var event event.Event

	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event Created",
		"event":   event,
	})
}
