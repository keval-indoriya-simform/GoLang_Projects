package models

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model

	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

func GetAllPeople(People *[]Person) {
	db.Preload("Books").Find(&People)
}

func GetPeopleByID(People *Person, id string) {
	db.Preload("Books").Find(&People, id)
}

func AddPerson(people *Person) {
	db.Create(&people)
}

func DeletePeopleByID(People *Person, id string) {
	db.Delete(&People, id)
}

func UpdatePeopleByID(People *Person) {
	db.Select("ID", "Name", "Email", "Books").Save(&People)
}
