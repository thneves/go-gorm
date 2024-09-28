package database

import (
	"go-gorm/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	// dsn means Data Source Name, it's a string to specify the information requiired to connect to a database
	dsn := "host=localhost user=postgres password=postgres dbname=go_auth_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Event{}) // must pass the pointer of the struct type, which is required by GORM.

	if err != nil {
		log.Fatal("failed to migrate databse")
	}

	return db, nil
}

/*
func createTables(db *gorm.DB) {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			datetime DATETIME NOT NULL,
			user_id INTEGER
		)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		log.Fatal(err)
	}
}
*/
