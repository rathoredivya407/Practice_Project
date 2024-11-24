package models

//book structure
type BookStore struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var Book BookStore
