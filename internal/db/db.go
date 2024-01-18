package db

import (
	"fmt"
	"log"

	"github.com/sha1sof/Effective_Mobile_test/internal/config"
	"github.com/sha1sof/Effective_Mobile_test/internal/migrations"
	"github.com/sha1sof/Effective_Mobile_test/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Get().DBHost,
		config.Get().DBPort,
		config.Get().DBUser,
		config.Get().DBPassword,
		config.Get().DBName)

	var err error
	db, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	if err := migrations.Migrate(db); err != nil {
		log.Fatal("Error applying migrations:", err)
	}
}

func GetPeople() ([]model.Person, error) {
	var people []model.Person

	result := db.Find(&people)
	if result.Error != nil {
		return nil, result.Error
	}

	return people, nil
}

func DeletePerson(personID int) error {
	result := db.Where("id = ?", personID).Delete(&model.Person{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func CreatePerson(person *model.Person) error {
	result := db.Create(person)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdatePerson(updatedPerson *model.Person) error {
	result := db.Save(updatedPerson)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func FilterPeople(name, surname, patronymic string, age int, gender, nationality string, page, pageSize int) ([]model.Person, error) {
	var people []model.Person

	query := db
	if name != "" {
		query = query.Where("name = ?", name)
	}
	if surname != "" {
		query = query.Where("surname = ?", surname)
	}
	if patronymic != "" {
		query = query.Where("patronymic = ?", patronymic)
	}
	if age > 0 {
		query = query.Where("age = ?", age)
	}
	if gender != "" {
		query = query.Where("gender = ?", gender)
	}
	if nationality != "" {
		query = query.Where("nationality = ?", nationality)
	}

	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&people).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return people, nil
}
