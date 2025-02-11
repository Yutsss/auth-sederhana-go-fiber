package command

import (
	"auth-sederhana-go-fiber/migration"
	"gorm.io/gorm"
	"log"
	"os"
)

func Commands(db *gorm.DB) bool {
	migrate := false
	run := true

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}
	}

	if migrate {
		if err := migration.Migrate(db); err != nil {
			log.Fatalf("error migration: %v", err)
		}
		log.Println("migration completed successfully")
		return false
	}

	if run {
		return true
	}

	return false
}
