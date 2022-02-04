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
func GetBookByID(db *gorm.DB, id string) (*models.Book, error) {
	book := models.Book{}
	err := db.Select("books.*").Group("books.id").Where("books.id= ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func DeleteBookByID(db *gorm.DB, id string) error {
	var book models.Book
	err := db.Where("id=?", id).Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateBookByID(db *gorm.DB, book *models.Book) error {
	err := db.Save(book).Error
	if err != nil {
		return err
	}
	return nil
}
func SaveBook(db *gorm.DB, book *models.Book) error {
	err := db.Save(book).Error
	if err != nil {
		return err
	}
	return nil
}
