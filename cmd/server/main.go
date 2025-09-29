package main

import (
    "fmt"
    "log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "jackdodev"
	dbPassword = "wornr123"
	dbName = "webpage"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err	)
	}

	server := NewServer(db)

	server.start()
	defer server.stop()
}

func connectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database successfully")

	return db, nil
}
