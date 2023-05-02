package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
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
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=Session sslmode=disable password=Ke^@l081001 port=5432")
	db1URI := fmt.Sprintf("host=localhost user=postgres dbname=Hooks sslmode=disable password=Ke^@l081001 port=5432")
	db2URI := fmt.Sprintf("host=localhost user=postgres dbname=AssociationMode2 sslmode=disable password=Ke^@l081001 port=5432")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	db.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{
			postgres.New(postgres.Config{
				DSN:                  db1URI,
				PreferSimpleProtocol: true,
			}),
		},
		Replicas: []gorm.Dialector{
			postgres.New(postgres.Config{
				DSN:                  db2URI,
				PreferSimpleProtocol: true,
			}),
		},
		Policy:            dbresolver.RandomPolicy{},
		TraceResolverMode: true,
	}))

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
	//
	//err = db.AutoMigrate(&Person{}, &Book{})
	//errorCheck(err)

	var person Person
	db.Preload("Books").First(&person, 3)

	fmt.Println(person)

	db.Create(&Person{
		Name:  person.Name,
		Email: person.Email,
		Books: person.Books,
	})
}
