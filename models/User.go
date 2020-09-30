package model

import (
	"database/sql"
	"errors"

	db "database"
)

// User defines the structure of the user entity
type User struct {
	ID        int    `json:"id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// GetUserByEmail returns data of a user.
func GetUserByEmail(userEmail string) (User, error) {
	var user User

	const query = `SELECT password FROM "user" WHERE email = ?`
	err := db.DB.QueryRow(query, userEmail).Scan(&user.Password)

	if err == sql.ErrNoRows {
		return user, errors.New("User is not found")
	}

	if err != nil {
		return user, err
	}

	return user, nil
}

