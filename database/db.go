// initializing connection to the database postgres
// code which will be responsible to create the connection  between  the RESTAPI golang application and th database -
// - apllication and gice the control to the DB global variable
package database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pallab-gogoi/Go_Rest_API/models"
)

var DB *gorm.DB //we use this global variable to interect with the database
func GetDB() *gorm.DB {
	return DB
}
func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	userName := "postgres"
	passWord := "postgres"
	arg := "host=" + host + "port=" + port + " user=" + userName + "dbName=" + dbName + "sslmode=disable password=" + passWord
	db, err := gorm.Open("postgres", arg) //this line of code creates the connection b/w this application and the database server
	if err != nil {
		log.Fatal(err)
	}
	// if connection established automigrated tables
	db.AutoMigrate(models.Book{}) //it will automatically create table
	DB = db
}
