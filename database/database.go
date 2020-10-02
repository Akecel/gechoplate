package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL package
	"github.com/spf13/viper"
)

var (
	// DB var
	DB *sql.DB
)

// GetMySQLDataSourceName returns environment variable for database connection.
func GetMySQLDataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
	)
}

// Connect connect to the database.
func Connect() {
	var err error

	dsn := GetMySQLDataSourceName()

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := DB.Ping(); err != nil {
		panic(err)
	}
}
