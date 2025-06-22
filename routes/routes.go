package routes

import (
	"fmt"
	controlllers "patient/controllers"

	"github.com/gorilla/mux"
)

func PatientRoute(r *mux.Router) {
	fmt.Println("initiating the routers for the method!!")
	r.HandleFunc("/patients", controlllers.GetPatientDetails).Methods("GET")
}
