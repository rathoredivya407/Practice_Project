package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"patient/config"
	"patient/models"
)

// get the patient details
func GetPatientDetails(w http.ResponseWriter, r *http.Request) {
	var patients []models.Patient
	config.Db.Find(&patients)
	fmt.Println("the patient details:", patients)
	//parsing the html file
	tmpl, err := template.ParseFiles("templates/patients.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//executing and writing data
	err = tmpl.Execute(w, patients)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
