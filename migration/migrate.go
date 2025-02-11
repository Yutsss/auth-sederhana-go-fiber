package migration

import (
	"auth-sederhana-go-fiber/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entities.User{},
	)

	if err != nil {
		return err
	}

	return nil
}
