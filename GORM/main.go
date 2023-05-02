package main

import (
	"GORM/connection"
	"GORM/models"
	"GORM/routers"
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
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

	log.Fatal(http.ListenAndServe(":8080", routers.Router))
}
