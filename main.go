package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Email        string
	PhoneNumbers []string `gorm:"type:JSONB;serializer:json"`
}

func main() {
	// PostgreSQL connection URI
	uri := "postgres://developer:password123@127.0.0.1:30432/mydb?sslmode=disable"

	// Connect to the database
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Error migrating the schema: %v", err)
	}

	// Create a new user
	user := User{Name: "asafssfsfsfsfsfs", Email: "asdaadadadad@example.com", PhoneNumbers: []string{"1234567890", "0987654321"}}
	db.Create(&user)
	if db.Error != nil {
		log.Fatalf("Error creating user: %v", db.Error)
	}

	log.Println("User created successfully:", user)
}
