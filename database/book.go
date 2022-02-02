package database

import (
	"github.com/jinzhu/gorm"
	"github.com/pallab-gogoi/Go_Rest_API/models"
)

// GetBooks is creating connection and interecting from golang app to db server via db variable
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
	book := models.Book{}
	return models.Book{}, nil // will populate later 1:35:47
}

func DeleteBookByID(db *gorm.DB, id string) error {
	return nil // will populate later
}

func UpdateBookByID(db *gorm.DB, book *models.Book) error {
	return nil // will populate later
}
func SaveBook(db *gorm.DB, book *models.Book) error {
	err := db.Save(book).Error
	if err != nil {
		return err
	}
	return nil
}
