package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Person struct {
	ID    int
	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

type Book struct {
	ID       int
	Title    string
	Author   string
	PersonID int
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=Tutorials sslmode=disable password=Ke^@l081001 port=5432")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully Connected to Database")
	}

	defer func() {
		db, err := db.DB()
		if err != nil {
			fmt.Println(err)
		}
		err = db.Close()
		errorCheck(err)
		fmt.Println("Closing Database Connection")
	}()

	err = db.AutoMigrate(&Person{}, &Book{})
	errorCheck(err)

	//person := []Person{
	//	{
	//		Name:  "Keval",
	//		Email: "keval.i@simformsolutions.com",
	//		Books: []Book{
	//			{
	//				Title:  "basic of golang",
	//				Author: "google",
	//			},
	//		},
	//	},
	//	{
	//		Name:  "Kishan",
	//		Email: "kishan.m@simformsolutions.com",
	//		Books: []Book{
	//			{
	//				Title:  "basic",
	//				Author: "basic",
	//			},
	//			{
	//				Title:  "advance",
	//				Author: "advance",
	//			},
	//		},
	//	},
	//}
	//
	//db.Create(&person)

	var people []Person
	var person1 Person
	stmts := db.Session(&gorm.Session{PrepareStmt: true})
	stmts.First(&person1)
	stmts.Find(&people)

	stmtManager, _ := stmts.ConnPool.(*gorm.PreparedStmtDB)

	for _, stmt := range stmtManager.PreparedSQL {
		fmt.Println(stmt)
	}
	fmt.Println(person1, "\n", people)

}
