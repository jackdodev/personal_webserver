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
		log.Fatalf("Failed to connect to database: %v", err)
	}

	server := NewServer(db)

	server.start()
	defer server.stop()
}

func connectDB() (*gorm.DB, error) {
	// Read environment variables and provide sensible defaults where appropriate.
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Basic validation to help surface configuration problems early.
	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		return nil, fmt.Errorf("database environment variables not fully set: DB_HOST=%q DB_USER=%q DB_NAME=%q", dbHost, dbUser, dbName)
	}

	// Use %s for port (we keep it as a string from the environment).
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	
	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database successfully")

	return db, nil
}
