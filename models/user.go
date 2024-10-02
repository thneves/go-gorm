package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Email     string `json:"email" gorm:"unique;not null" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) Save(db *gorm.DB) (User, error) {

	result := db.Create(&u)

	if result.Error != nil {
		return User{}, result.Error
	}

	return u, nil
}
