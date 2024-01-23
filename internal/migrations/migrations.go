package migrations

import (
	"log"

	"github.com/sha1sof/Effective_Mobile_test/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	log.Println("Starting database migration...")

	err := db.AutoMigrate(&model.Person{})
	if err != nil {
		log.Printf("Error during migration: %v\n", err)
		return err
	}

	log.Println("Database migration completed successfully.")
	return nil
}
