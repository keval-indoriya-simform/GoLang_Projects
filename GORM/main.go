package main

import (
	"GORM/connection"
	"GORM/models"
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
)

var (
	//person = &models.Person{Name: "Keval", Email: "kevalindoriya5@gmail.com"}
	//books  = []models.Book{
	//	{Title: "The India Story", Author: "Bimal Jalal", CallNumber: 7123, PersonID: 1},
	//	{Title: "Hear Yourself", Author: "Prem Rawat", CallNumber: 6454, PersonID: 1},
	//}
	ctx context.Context
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
	//db.Select("id", "name", "email").Find(&people)
	db.WithContext(ctx).Find(&people)

	for i := range people {
		err = db.Model(&people[i]).Select("title", "author", "call_number, person_id").Association("Books").Find(&books)
		errorCheck(err)
		people[i].Books = books
	}

	fmt.Println(people)
	//log.Fatal(http.ListenAndServe(":8080", routers.Router))
}
