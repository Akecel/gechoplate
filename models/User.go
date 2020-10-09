package models

import (
	"gorm.io/gorm"
)

// User defines the structure of the user entity
type User struct {
	gorm.Model
	LastName  string `json:"last_name" form:"last_name" gorm:"not nul; type:varchar(55)"`
	FirstName string `json:"first_name" form:"first_name" gorm:"not null; type:varchar(55)"`
	Email     string `json:"email" form:"email" gorm:"unique; not null; type:varchar(55)"`
	Password  string `form:"password" gorm:"not null; type:varchar(255)"`
}
