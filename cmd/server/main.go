package main

import (
    "fmt"
    "log"
		"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database successfully")

	return db, nil
}
