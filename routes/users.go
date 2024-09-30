package routes

import (
	"go-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func signUp(context *gin.Context, db *gorm.DB) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err = user.Save(db)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save user.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "new user created",
		"user":    user,
	})
}
