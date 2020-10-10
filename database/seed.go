package database

import (
	"fmt"
	"gechoplate/helpers"
	"gechoplate/models"
)

// Seed seed the database
func Seed() {
	UserSeeder("Doe", "John", "admin@gechoplate.com", "root")
}

// UserSeeder seed users table.
func UserSeeder(lastName string, firstName string, email string, password string) {
	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		fmt.Printf("Seeding error\n")
	}
	Gorm.Create(&models.User{LastName: lastName, FirstName: firstName, Email: email, Password: hashedPassword})
}
