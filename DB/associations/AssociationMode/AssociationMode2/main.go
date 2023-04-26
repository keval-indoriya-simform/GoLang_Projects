package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=AssociationMode2 sslmode=disable password=Ke^@l081001 port=5432")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
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
	//var person []Person
	//db.Find(&person)
	//var Books []Book
	//for i := range person {
	//	db.Model(&person[i]).Association("Books").Find(&Books)
	//	person[i].Books = Books
	//	fmt.Println(person[i])
	//}

	//var person []Person
	//person = []Person{
	//	{
	//		ID:    1,
	//		Name:  "Keval",
	//		Email: "keval.i@simformsolutions.com",
	//		Books: []Book{
	//			{
	//				ID:       1,
	//				Title:    "basic of golang",
	//				Author:   "google",
	//				PersonID: 1,
	//			},
	//			{
	//				ID:       2,
	//				Title:    "hello",
	//				PersonID: 1,
	//			},
	//		},
	//	},
	//	{
	//		ID:    2,
	//		Name:  "Kishan",
	//		Email: "kishan.m@simformsolutions.com",
	//		Books: []Book{
	//			{
	//				ID:       3,
	//				Title:    "basic",
	//				Author:   "basic",
	//				PersonID: 2,
	//			},
	//			{
	//				ID:       4,
	//				Title:    "advance",
	//				Author:   "advance",
	//				PersonID: 2,
	//			},
	//		},
	//	},
	//}
	var person Person
	db.Preload("Books").Find(&person, "id = 2")
	var Books Book
	db.Find(&Books, "id = 3")
	//fmt.Println(db.Model(&person).Where("person_id = 2").Association("Books").Count())

	//db.Model(&person).Association("Books").Append(&Book{
	//	Title:  "hello",
	//	Author: "gooo",
	//})
	db.Model(&person).Association("Books").Replace(&Book{Title: "replaced"}, &Book{Title: "this is replaced"})
	//db.Model(&person).Association("Books").Clear()
	//db.Model(&person).Association("Books").Delete(&Books)

	//count := db.Model(&Person{}).Association("Books").Count()
	//fmt.Println(count)

	fmt.Println(person)
}
