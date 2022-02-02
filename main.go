package main

import (
	"log"

	"github.com/pallab-gogoi/Go_Rest_API/database"
	"github.com/pallab-gogoi/Go_Rest_API/routers"
)

func main() {
	database.Setup()                    // establish the database connection
	engine := routers.Router()          //get the customized engine
	err := engine.Run("127.0.0.1:8080") // start the engine
	if err != nil {
		log.Fatal(err)
	}
}
