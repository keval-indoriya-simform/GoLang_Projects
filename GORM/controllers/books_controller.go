package controllers

import (
	"GORM/connection"
	"GORM/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	var books []models.Book

	fmt.Println("Book")
	db := connection.GetConnection()
	db.Find(&books)

	fmt.Println(books)
	b, err := json.MarshalIndent(&books, "", "\t")
	ErrorCheck(err)
	fmt.Fprint(writer, string(b))
}

func GetBook(writer http.ResponseWriter, request *http.Request) {
	var books []models.Book
	params := mux.Vars(request)

	fmt.Println("Book With ID :", params["ID"])
	db := connection.GetConnection()
	db.Find(&books, params["ID"])

	fmt.Println(books)
	b, err := json.MarshalIndent(&books, "", "\t")
	ErrorCheck(err)
	fmt.Fprintln(writer, "Book With ID :", params["ID"])
	fmt.Fprint(writer, string(b))
}
