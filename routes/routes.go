package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(server *gin.Engine, db *gorm.DB) {
	server.GET("/events", func(c *gin.Context) {
		getEvents(c, db)
	})
	server.POST("/create_event", func(c *gin.Context) {
		createEvent(c, db)
	})
	server.GET("/event/:id", func(c *gin.Context) {
		getEvent(c, db)
	})
	server.DELETE("/event/:id", func(c *gin.Context) {
		deleteEvent(c, db)
	})
	server.PUT("/event/:id", func(c *gin.Context) {
		updateEvent(c, db)
	})
}
