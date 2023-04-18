package controllers

import (
	"GORM/connection"
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
	var books []models.Book

	fmt.Println("People")
	db := connection.GetConnection()
	db.Find(&people)

	for i := range people {
		err := db.Model(&people[i]).Association("Books").Find(&books)
		ErrorCheck(err)
		people[i].Books = books
	}
	fmt.Println(people)
	p, err := json.MarshalIndent(&people, "", "\t")
	ErrorCheck(err)
	fmt.Fprint(writer, string(p))
}

func GetPerson(writer http.ResponseWriter, request *http.Request) {
	var people []models.Person
	var books []models.Book
	params := mux.Vars(request)

	fmt.Println("People With ID :", params["ID"])
	db := connection.GetConnection()
	db.Find(&people, params["ID"])

	for i := range people {
		err := db.Model(&people[i]).Association("Books").Find(&books)
		ErrorCheck(err)
		people[i].Books = books
	}
	fmt.Println(people)
	p, err := json.MarshalIndent(&people, "", "\t")
	ErrorCheck(err)
	fmt.Fprintln(writer, "People With ID :", params["ID"])
	fmt.Fprint(writer, string(p))
}
