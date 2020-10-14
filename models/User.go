package models

import (
	"time"

	"gorm.io/gorm"
)

// User defines the structure of the user entity
type User struct {
	gorm.Model
	LastName  string `form:"last_name" gorm:"not null; type:varchar(55)"`
	FirstName string `form:"first_name" gorm:"not null; type:varchar(55)"`
	Email     string `form:"email" gorm:"unique; not null; type:varchar(55)"`
	Password  string `form:"password" gorm:"not null; type:varchar(255)"`
}

// APIUser defines the structure of the user api repsponse entity
type APIUser struct {
	ID        uint      `json:"id"`
	LastName  string    `json:"last_name"`
	FirstName string    `json:"first_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
