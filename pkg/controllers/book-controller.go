package controllers

import (
	"bookstore/pkg/config"
	"bookstore/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create a book list info!!")
	w.Header().Set("Content-Type", "application/json")
	var book models.BookStore
	json.NewDecoder(r.Body).Decode(&book) //json.Unmarshallinh
	config.Db.Create(&book)
	json.NewEncoder(w).Encode(book) //json.Marshalling
	w.WriteHeader(http.StatusOK)
}
func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get the book list!!")
	w.Header().Set("Content-Type", "application/json")
	var book []models.BookStore //array of all the entries in the database
	config.Db.Find(&book)       //find the book data
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)

}
func GetBookbyBookId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get the book by bookid")
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"] // extracting the id from path parameter
	var book models.BookStore
	err := config.Db.First(&book, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "the book details not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)

}
func UpdateBookbyBookId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update the book data by bookid")
	w.Header().Set("content-type", "application/json")
	id := mux.Vars(r)["id"]
	conId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("unable to convert id in integer value")
		return
	}
	var Book models.BookStore
	// find the book in the database
	err = config.Db.First(&Book, conId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "the book details not found", http.StatusNotFound)
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	//create empty structue variable
	var updatedBook models.BookStore

	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		fmt.Println("error while decoding the new updated book data")
		return
	}

	config.Db.Save(&updatedBook)
	json.NewEncoder(w).Encode(&updatedBook)
	w.WriteHeader(http.StatusOK)
}
func DeleteBookbyBookId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete the book by book id")
	w.Header().Set("content-type", "application/json")
	id := mux.Vars(r)["id"]
	var book models.BookStore
	config.Db.Delete(&book, id)
	w.WriteHeader(http.StatusOK)

}
