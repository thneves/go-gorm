package routes

import (
	"go-gorm/controllers"
	"go-gorm/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(server *gin.Engine, db *gorm.DB) {
	server.POST("/signup", func(c *gin.Context) {
		controllers.SignUp(c, db)
	})
	server.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})

	protected := server.Group("/api")
	protected.Use(middleware.Authorization())

	{
		protected.GET("/events", func(c *gin.Context) {
			controllers.GetEvents(c, db)
		})
		protected.POST("/create_event", func(c *gin.Context) {
			controllers.CreateEvent(c, db)
		})
		protected.GET("/event/:id", func(c *gin.Context) {
			controllers.GetEvent(c, db)
		})
		protected.DELETE("/event/:id", func(c *gin.Context) {
			controllers.DeleteEvent(c, db)
		})
		protected.PUT("/event/:id", func(c *gin.Context) {
			controllers.UpdateEvent(c, db)
		})
	}
}
