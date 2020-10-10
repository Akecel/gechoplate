package database

import "gechoplate/models"

// Migrate migrate the database schema.
func Migrate() {
	Gorm.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.User{})
}
