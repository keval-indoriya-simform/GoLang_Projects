package main

import "log"

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

var (
	person = &Person{Name: "Abhishek", Email: "abhishek.m@simformsolutions.com"}

	books = []Book{
		{Title: "Basics Of Golang", Author: "Google", PersonID: 2},
		{Title: "Advance Of Golang", Author: "Google", PersonID: 2},
	}
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

}
