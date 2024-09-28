package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"size:255;not null"`
	Description string    `json:"description" gorm:"size:255;not null"`
	Location    string    `json:"location" gorm:"default: 'planet earth'"`
	DateTime    time.Time `json:"date_time"`
	UserID      uint      `json:"user_id" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
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

func (e *Event) Update(db *gorm.DB) (Event, error) {
	// Ensure that we update an existing event by its primary key (ID)
	result := db.Model(e).Updates(map[string]interface{}{
		"Name":        e.Name,
		"Description": e.Description,
		"Location":    e.Location,
		"DateTime":    e.DateTime, // Only update if it's provided
	})

	// Check for errors during the update operation
	if result.Error != nil {
		return Event{}, result.Error
	}

	// Return the updated event
	return *e, nil
}
