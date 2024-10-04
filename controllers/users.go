package controllers

import (
	"go-gorm/models"
	"go-gorm/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var jwtSecret = []byte("bigmistery")

func SignUp(context *gin.Context, db *gorm.DB) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Password, err = utils.HashPassword(user.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
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

func Login(context *gin.Context, db *gorm.DB) {
	var userInput struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := context.ShouldBindJSON(&userInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	// Find User model by email
	err := db.Where("email = ?", userInput.Email).First(&user).Error

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid email",
		})
		return
	}

	// Check pass

	isValidPassword := utils.CheckPassword(userInput.Password, user.Password)

	if !isValidPassword {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to create token",
			"message": err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "login succesful!",
		"token":   token,
	})
}
