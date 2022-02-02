package database

import (
	"github.com/jinzhu/gorm"
	"github.com/pallab-gogoi/Go_Rest_API/models"
)

func GetBooks(db *gorm.DB) ([]models.Book, error) {
	books := []models.Book{} // slice of book
	query := db.Select("books.*")
	err := query.Find(&books).Error // .Error is from gorm
	if err != nil {
		return nil, err
	}
	return books, nil
}
func GetBookByID(db *gorm.DB, id string) (models.Book, error) {
	return models.Book{}, nil // will populate later
}

func DeleteBookByID(db *gorm.DB, id string) error {
	return nil // will populate later
}

func UpdateBookByID(db *gorm.DB, book *models.Book) error {
	return nil // will populate later
}
