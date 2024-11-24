package routes

import (
	"bookstore/pkg/controllers"
	"fmt"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	/*initializing gorila mux that implements a request router and
	dispatcher for matching incoming requests to their respective handler.*/
	//r := mux.NewRouter()
	fmt.Println("routers intiated for request!!!!")
	r.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.GetBookbyBookId).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.UpdateBookbyBookId).Methods("PUT")
	r.HandleFunc("/book/{id}", controllers.DeleteBookbyBookId).Methods("DELETE")
}
