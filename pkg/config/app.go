package config

/*In this file logic for database is added*/
import (
	"bookstore/pkg/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "testdb"
)

func DbConnect() {
	var err error
	fmt.Println("database connection!!!!")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("failed to connect to the database!!")
	}
	//this is used to create the table and automigrate the table contents
	err = Db.AutoMigrate(&models.BookStore{})
	if err != nil {
		log.Fatal(err)
		fmt.Println("unable to migrate the data to models")
	}
}
