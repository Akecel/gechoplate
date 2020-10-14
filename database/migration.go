package database

import (
	"fmt"
	"gechoplate/models"
)

// Migrate migrate the database schema.
func Migrate() {
	if err := Gorm.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.User{}); err != nil {
		panic(fmt.Errorf("Migration error: %s", err.Error()))
	}
}
