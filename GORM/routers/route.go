package routers

import (
	"GORM/controllers"
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/people", controllers.GetPeople).Methods("GET")
	Router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	Router.HandleFunc("/person/{ID}", controllers.GetPerson).Methods("GET")
	Router.HandleFunc("/book/{ID}", controllers.GetBook).Methods("GET")
	Router.HandleFunc("/create/people", controllers.CreatePerson).Methods("POST")
	Router.HandleFunc("/create/book", controllers.CreateBook).Methods("POST")
	Router.HandleFunc("/delete/person/{ID}", controllers.DeletePerson).Methods("DELETE")
	Router.HandleFunc("/delete/book/{ID}", controllers.DeleteBook).Methods("DELETE")
	Router.HandleFunc("/update/person", controllers.UpdatePerson).Methods("PUT")
	Router.HandleFunc("/update/book", controllers.UpdateBook).Methods("PUT")
}
