package models

import (
	"gorm.io/gorm"
)

// User defines the structure of the user entity
type User struct {
	gorm.Model
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string
}
