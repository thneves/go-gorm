package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	NAME        string    `json:"name" gorm:"size:255;not null"`
	Description string    `json:"description" gorm:"size:255;not null"`
	Location    string    `json:"location" gorm:"default: 'planet earth'"`
	DateTime    time.Time `json:"date_time"`
	UserID      uint      `json:"user_id" gorm:"foreignKey:UserID"`
}

var events = []Event{}

func (e Event) Save(db *gorm.DB) (Event, error) {
	result := db.Create(&e)

	if result.Error != nil {
		return Event{}, result.Error
	}

	return e, nil
}

func GetAllEvents(db *gorm.DB) ([]Event, error) {
	events := []Event{}
	result := db.Find(&events)

	if result.Error != nil {
		return []Event{}, result.Error
	}

	return events, nil
}
