package database

import (
	"fmt"
	"gechoplate/helpers"
	"gechoplate/models"
)

// Seed seed the database
func Seed() {
	UserSeeder("Doe", "John", "admin@gechoplate.com", "root", true)
}

// UserSeeder seed users table.
func UserSeeder(lastName string, firstName string, email string, password string, admin bool) {
	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		panic(fmt.Errorf("Seeding error: %s", err.Error()))
	}

	Gorm.Where(models.User{Email: email}).FirstOrCreate(&models.User{LastName: lastName, FirstName: firstName, Email: email, Password: hashedPassword, Admin: admin})
}
