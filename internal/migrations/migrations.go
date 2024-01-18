package migrations

import (
	"github.com/sha1sof/Effective_Mobile_test/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.Person{})
}
