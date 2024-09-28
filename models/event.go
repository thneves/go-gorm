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
	// query := `
	// INSERT INTO events(name, description, location, dateTime, user_id)
	// VALUES(?, ?, ?, ?, ?)
	//`
	result := db.Create(&e)

	if result.Error != nil {
		return Event{}, result.Error
	}

	return e, nil
}

func Delete(id int, db *gorm.DB) (Event, error) {
	event := Event{}

	result := db.Delete(&event, id)

	if result.Error != nil {
		return Event{}, result.Error
	}

	return event, nil
}

func GetAllEvents(db *gorm.DB) ([]Event, error) {
	events := []Event{}
	result := db.Find(&events)

	if result.Error != nil {
		return []Event{}, result.Error
	}

	return events, nil
}

func GetEventById(id int, db *gorm.DB) (Event, error) {
	event := Event{}

	result := db.First(&event, id)
	// SELECT * FROM events WHERE id = id;

	if result.Error != nil {
		return Event{}, result.Error
	}

	return event, nil
}
