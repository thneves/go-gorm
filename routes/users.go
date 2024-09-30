package routes

import (
	"go-gorm/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var jwtSecret = []byte("This is super secret")

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

func login(context *gin.Context, db *gorm.DB) {
	var userInput struct {
		Email    string `json:"email" binding:"required, email"`
		Password string `json:"password" binding:"required"`
	}

	if err := context.ShouldBindJSON(&userInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	// Find User model by email
	err := db.Where("email = ?", userInput.Email).First(&user)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid email",
		})
		return
	}

	// Check pass

	if err := user.CheckPassword(userInput.Password); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create token",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
