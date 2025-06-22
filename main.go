package main

import (
	"fmt"
	"log"
	"net/http"
	"patient/config"
	"patient/routes"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("This is Patient Management System!!")

	r := mux.NewRouter()

	fmt.Println("Starting the Db connection!!")
	config.DbConnect()
	//serving css file
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("assets/"))))

	fmt.Println("starting the Routes!!")
	routes.PatientRoute(r)

	//making the server connection
	log.Fatal(http.ListenAndServe(":8080", r))

}
