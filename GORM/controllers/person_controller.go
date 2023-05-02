package controllers

import (
	"GORM/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetPeople(writer http.ResponseWriter, request *http.Request) {
	var people []models.Person
	models.GetAllPeople(&people)
	p, err := json.MarshalIndent(&people, "", "\t")
	ErrorCheck(err)
	fmt.Fprint(writer, string(p))
}

func GetPerson(writer http.ResponseWriter, request *http.Request) {
	var people models.Person
	params := mux.Vars(request)
	models.GetPeopleByID(&people, params["ID"])
	p, err := json.MarshalIndent(&people, "", "\t")
	ErrorCheck(err)
	fmt.Fprintln(writer, "People With ID :", params["ID"])
	fmt.Fprint(writer, string(p))
}

func CreatePerson(writer http.ResponseWriter, request *http.Request) {
	var people models.Person
	err := json.NewDecoder(request.Body).Decode(&people)
	ErrorCheck(err)
	models.AddPerson(&people)
	p, err := json.MarshalIndent(&people, "", "\t")
	ErrorCheck(err)
	fmt.Fprint(writer, string(p))
}

func DeletePerson(writer http.ResponseWriter, request *http.Request) {
	var people models.Person
	params := mux.Vars(request)
	models.DeletePeopleByID(&people, params["ID"])
	fmt.Fprintln(writer, "People With ID :", params["ID"], "Deleted")
}

func UpdatePerson(writer http.ResponseWriter, request *http.Request) {
	var people models.Person
	err := json.NewDecoder(request.Body).Decode(&people)
	ErrorCheck(err)
	models.UpdatePeopleByID(&people)
	fmt.Fprintln(writer, "Person With ID :", people.ID, "Updated")
}
