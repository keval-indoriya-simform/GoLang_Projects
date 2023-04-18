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
}
