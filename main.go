package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// dsn means Data Source Name, it's a string to specify the information requiired to connect to a database
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	pgDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	pgDB.SetConnMaxIdleTime(10)
	pgDB.SetMaxOpenConns(100)
	pgDB.SetConnMaxLifetime(time.Hour)
}
