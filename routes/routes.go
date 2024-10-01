package routes

import (
	"go-gorm/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(server *gin.Engine, db *gorm.DB) {
	server.GET("/events", func(c *gin.Context) {
		controllers.GetEvents(c, db)
	})
	server.POST("/create_event", func(c *gin.Context) {
		controllers.CreateEvent(c, db)
	})
	server.GET("/event/:id", func(c *gin.Context) {
		controllers.GetEvent(c, db)
	})
	server.DELETE("/event/:id", func(c *gin.Context) {
		controllers.DeleteEvent(c, db)
	})
	server.PUT("/event/:id", func(c *gin.Context) {
		controllers.UpdateEvent(c, db)
	})
	server.POST("/signup", func(c *gin.Context) {
		controllers.SignUp(c, db)
	})

	server.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})
}
