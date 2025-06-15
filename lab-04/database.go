package database

import (
	"fmt"
)

type Database struct {
	data map[string]string
}

func InitializeTestEnvironment() {
	fmt.Println("InitializeTestEnvironment called")
}

func CleanupTestEnvironment() {
	fmt.Println("CleanupTestEnvironment called")
}

func NewTestDatabase() *Database {
	fmt.Println("NewTestDatabase called")
	return &Database{
		data: make(map[string]string),
	}
}

func (db *Database) Close() {
	fmt.Println("Database connection closed")
}

func (db *Database) Insert(key, value string) error {
	fmt.Printf("Inserted key: %s, value: %s\n", key, value)
	db.data[key] = value
	return nil
}

func (db *Database) Get(key string) (string, error) {
	if value, exists := db.data[key]; exists {
		fmt.Printf("Retrieved key: %s, value: %s\n", key, value)
		return value, nil
	}
	fmt.Printf("Key: %s not found\n", key)
	return "", fmt.Errorf("key not found")
}

func (db *Database) Delete(key string) error {
	if _, exists := db.data[key]; !exists {
		return fmt.Errorf("key not found")
	}
	delete(db.data, key)
	fmt.Printf("Deleted key: %s\n", key)
	return nil
}