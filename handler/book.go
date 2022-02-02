package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pallab-gogoi/Go_Rest_API/database"
	"github.com/pallab-gogoi/Go_Rest_API/models"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetBooks(in *gin.Context) { //pointer receiver
	books, err := database.GetBooks(h.DB) // collect books, h.DB is fully configured and can give the access of book table
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, books) //  return books
}

/*
h:=Handler{}
h.DB=GetDB()
h.GetBooks()
*/

func (h *Handler) SaveBook(in *gin.Context) { //pointer receiver
	book := models.Book{}
	err := in.BindJSON(&book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}

	err = database.SaveBook(h.DB, &book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, book)
}
