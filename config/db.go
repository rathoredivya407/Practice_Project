package config

import (
	"fmt"
	"log"
	"os"
	"patient/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

// this function to initialise database connection
func DbConnect() {
	fmt.Println("connection to database!!")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found or failed to load")
	}
	//loading db credentials from env file
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}
	//this is used to create table and automigrate the contents in db
	err = Db.AutoMigrate(&models.Patient{})
	if err != nil {
		log.Fatal(err)
		fmt.Println("unable to migrate data to models ")
	}
	fmt.Println("entering the data in db!!")
	patientDummyData()
}

func patientDummyData() {
	var count int64
	Db.Model(&models.Patient{}).Count(&count)
	if count == 0 {
		patients := []models.Patient{
			{FullName: "Adam White", Age: 50, Gender: "Male", InsuranceProvider: "XYZ", PolicyNumber: "POL123"},
			{FullName: "Smith Pat", Age: 45, Gender: "Male", InsuranceProvider: "GHJ", PolicyNumber: "POL865"},
			{FullName: "Lucy Johnson", Age: 30, Gender: "Female", InsuranceProvider: "LMN", PolicyNumber: "POL987"},
			{FullName: "Diana Lee", Age: 56, Gender: "Female", InsuranceProvider: "OIU", PolicyNumber: "POL098"},
			{FullName: "Karl Max", Age: 66, Gender: "Male", InsuranceProvider: "DSA", PolicyNumber: "POL443"},
		}
		for _, p := range patients {
			//create the table
			Db.Create(&p)
		}
	}
	fmt.Println("Successfully entered the data in db!!")
}
