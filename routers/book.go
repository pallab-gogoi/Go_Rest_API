package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pallab-gogoi/Go_Rest_API/database"
	"github.com/pallab-gogoi/Go_Rest_API/handler"
)

func Router() *gin.Engine {
	router := gin.Default() // get the default engine for further customization
	api := handler.Handler{ //
		DB: database.GetDB(), //  set the handler DB
	}
	router.GET("/books", api.GetBooks) // set the function for http://localhost:8080/books : Get request
	// here api.GetBooks is used to pass function ,NOT to call function
	// while calling handler function gin will pass *gin.Context in the handler function

	router.POST("/books", api.SaveBook)
	return router
}
