package main

import (
	"GORM/connection"
	"GORM/models"
	"GORM/routers"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var (
	//person = &models.Person{Name: "Keval", Email: "kevalindoriya5@gmail.com"}
	//books  = []models.Book{
	//	{Title: "Basics Of Golang", Author: "Google", CallNumber: 1234, PersonID: 1},
	//	{Title: "Advance Of Golang", Author: "Google", CallNumber: 5678, PersonID: 1},
	//}
	db  *gorm.DB
	err error
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db := connection.GetConnection()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully Connected to Database")
	}

	defer func() {
		db, err := db.DB()
		errorCheck(err)
		err = db.Close()
		errorCheck(err)
		fmt.Println("Closing Database Connection")
	}()

	err = db.AutoMigrate(&models.Person{})
	errorCheck(err)
	err = db.AutoMigrate(&models.Book{})
	errorCheck(err)

	var people []models.Person
	var books []models.Book
	db.Select("id", "name", "email").Find(&people)

	for i := range people {
		err = db.Model(&people[i]).Select("title", "author", "call_number, person_id").Association("Books").Find(&books)
		errorCheck(err)
		people[i].Books = books
	}

	log.Fatal(http.ListenAndServe(":8080", routers.Router))
}
