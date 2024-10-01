package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Email     string `json:"email" gorm:"unique;not null" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u User) Save(db *gorm.DB) (User, error) {

	result := db.Create(&u)

	if result.Error != nil {
		return User{}, result.Error
	}

	return u, nil
}
