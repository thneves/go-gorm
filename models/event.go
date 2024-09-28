package event

import "time"

type Event struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	NAME        string    `json:"name" gorm:"size:255;not null"`
	Description string    `json:"description" gorm:"size:255;not null"`
	Location    string    `json:"location" gorm:"default: 'planet earth'"`
	DateTime    time.Time `json:"date_time"`
	UserID      uint      `json:"user_id" gorm:"foreignKey:UserID"`
}

var events = []Event{}

func (e Event) Save() {
	//to do: add it to database

	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
