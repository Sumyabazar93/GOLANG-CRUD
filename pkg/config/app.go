package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	// Replace "your-username", "your-password", "your-host", "your-port", and "your-dbname" with your actual database credentials.
	dsn := "user=postgres password=03220359 dbname=my_pgdb host=127.0.0.1 port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
