package controllers

import (
	"GORM/models"
	_ "context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	var books []models.Book
	models.GetAllBooks(&books)
	b, err := json.MarshalIndent(&books, "", "\t")
	ErrorCheck(err)
	fmt.Fprint(writer, string(b))
}

func GetBook(writer http.ResponseWriter, request *http.Request) {
	var books models.Book
	params := mux.Vars(request)
	models.GetBookByID(&books, params["ID"])
	b, err := json.MarshalIndent(&books, "", "\t")
	ErrorCheck(err)
	fmt.Fprintln(writer, "Book With ID :", params["ID"])
	fmt.Fprint(writer, string(b))
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	var book models.Book
	err := json.NewDecoder(request.Body).Decode(&book)
	ErrorCheck(err)
	models.AddBook(&book)
	p, err := json.MarshalIndent(&book, "", "\t")
	ErrorCheck(err)
	fmt.Fprint(writer, string(p))
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	var books models.Book
	params := mux.Vars(request)
	models.DeleteBookByID(&books, params["ID"])
	fmt.Fprintln(writer, "Book With ID :", params["ID"], "Deleted")
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	var books models.Book
	err := json.NewDecoder(request.Body).Decode(&books)
	ErrorCheck(err)
	models.UpdateBookByID(&books)
	fmt.Fprintln(writer, "Book With ID :", books.ID, "Updated")
}
