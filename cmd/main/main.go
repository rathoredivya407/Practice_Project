package main

import (
	"bookstore/pkg/config"
	"bookstore/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("HelloWorld!!!")

	r := mux.NewRouter()

	fmt.Println("the db connection start")
	config.DbConnect()

	fmt.Println("registerd routers!!!")
	routes.RegisterRoutes(r)

	//making the server connection
	log.Fatal(http.ListenAndServe(":8080", r))
}
