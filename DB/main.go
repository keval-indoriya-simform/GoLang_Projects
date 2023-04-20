package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Person struct {
	gorm.Model

	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

type APIPerson struct {
	Name  string
	Email string
}

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber int `gorm:"unique_index"`
	PersonID   int
}

var (
// person = &Person{Name: "Abhishek", Email: "abhishek.m@simformsolutions.com"}
//
//	books  = []Book{
//		{Title: "Basics Of Golang", Author: "Google", CallNumber: 1234, PersonID: 2},
//		{Title: "Advance Of Golang", Author: "Google", CallNumber: 5678, PersonID: 2},
//	}
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=DB sslmode=disable password=Ke^@l081001 port=5432")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

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

	err = db.AutoMigrate(&Person{})
	errorCheck(err)
	err = db.AutoMigrate(&Book{})
	errorCheck(err)

	//var person = &Person{Name: "keval", Email: "keval.i@simformsolutions.com"}
	//Create(person, db)

	//person := []*Person{
	//	{Name: "Juhi", Email: "juhi.m@simformsolutions.com"},
	//	{Name: "Shashwat", Email: "shashwat.m@simformsolutions.com"},
	//}
	//CreateMultiple(person, db)

	//db.Model(&Person{}).Create(map[string]interface{}{
	//	"Name": "abhishek", "Email": "abhishek.m@simformsolutions.com",
	//})
	//var person []Person
	//db.Take(&person)
	//db.Select("name").Find(&person)
	//db.Last(&person)
	//for i := range person {
	//	fmt.Println(person[i])
	//}
	//books := []Book{
	//	{Title: "Basics Of Golang", Author: "Google", CallNumber: 7123, PersonID: 1},
	//	{Title: "Advance Of Golang", Author: "Google", CallNumber: 6454, PersonID: 1},
	//	{Title: "The India Story", Author: "Bimal Jalal", CallNumber: 1456, PersonID: 2},
	//	{Title: "Hear Yourself", Author: "Prem Rawat", CallNumber: 4684, PersonID: 3},
	//}
	//db.Create(&books)

	//var person Person
	//db.Where("name = ?", "keval").Find(&person)
	//db.Find(&person, "id = 1")
	//fmt.Println(person)

	//var books []Book
	//db.Where("person_id = 1").Find(&books)
	//db.Where("person_id = 1").Or("person_id = 2").Find(&books)
	//db.Where("person_id = 1").Or("person_id = 2").Order("person_id desc").Find(&books)
	//db.Where("person_id = 1").Or("person_id = 2").Order("person_id desc").Limit(1).Find(&books)
	//db.Where("person_id = 1").Or("person_id = 2").Order("person_id desc").Offset(2).Find(&books)
	//db.Not("person_id = 1").Find(&books)
	//for _, book := range books {
	//	fmt.Println(book)
	//}
	//var result []int
	//db.Model(&Book{}).Select("conut(*)").Group("person_id").Find(&result)
	//fmt.Println(result)
	//var person Person
	//var books Book
	//person := db.Joins("p").Find(&books)
	//fmt.Println(person)

	//var apiPerson []APIPerson
	//var books []Book
	//db.Model(&Person{}).Find(&apiPerson)
	//db.Where("id = (?)", db.Table("books").Select("id").First(&books)).Find(&person)
	//db.Where("(name, email) IN ?", [][]interface{}{{"keval", "keval.i@simformsolutions.com"}, {"abhishek", "abhishek.m@simformsolutions.com"}}).Find(&person)
	//db.Where("name = @name OR name = @name", sql.Named("name", "keval")).Find(&person)

	//var result []map[string]interface{}
	//db.Model(&person).Find(&result)
	//fmt.Println(result)
	//fmt.Println(person)
	//db.FirstOrInit(&person, Person{
	//	Name:  "kishan",
	//	Email: "kishan.m@simformsolutions.com",
	//})
	//
	//db.FirstOrCreate(&person, Person{
	//	Name:  "keval",
	//	Email: "keval.i@simformsolutions.com",
	//})

	//fmt.Println(person)

	//rows, err := db.Model(&person).Rows()
	//
	//for rows.Next() {
	//	db.ScanRows(rows, &person)
	//	fmt.Println(person)
	//}
	//
	//defer rows.Close()

	//var names []string
	//db.Model(&person).Pluck("name", &names)
	//fmt.Println(names)
	//
	//var people []Person
	//db.Model(&person).Pluck("name", &names)
	//db.Select("name", "email").Scan(&people)

	//db.Scopes(IdLessThan4, NameIn).Find(&people)
	//fmt.Println(people)

	//var count int64
	//fmt.Println(count)

	//db.First(&person)
	//
	//person.Name = "meet"
	//
	//db.Save(&person)

	//db.Model(&person).Where("id = 1").Update("name", "Keval")
	//db.Model(&person).Clauses(clause.Returning{}).Where("id = 4").Update("name", "Abhishek")
	//fmt.Println(person)

	//var person Person
	//stmt := db.Session(&gorm.Session{DryRun: true}).First(&person, 1).Statement
	//fmt.Println(stmt.SQL.String())
	//fmt.Println(person)

	//var name string
	//row := db.Table("people").Where("id = 1").Select("name").Row()
	//row.Scan(&name)
	//fmt.Println(name)

	//var name string
	//var names []string
	var person []Person
	rows, err := db.Table("people").Rows()
	if err != nil {
		log.Fatal(err)
	}
	//for rows.Next() {
	//	rows.Scan(&name)
	//	names = append(names, name)
	//}
	//fmt.Println(names)

	for rows.Next() {
		db.ScanRows(rows, &person)
	}
	fmt.Println(person)
	defer rows.Close()
}

func IdLessThan4(db *gorm.DB) *gorm.DB {
	return db.Where("id < 4")
}

func NameIn(db *gorm.DB) *gorm.DB {
	return db.Where("name IN (?)", []string{"keval", "Juhi"})
}

//func Create(person *Person, db *gorm.DB) {
//	result := db.Create(person)
//	errorCheck(result.Error)
//	fmt.Println(person.ID)
//}

//func CreateMultiple(person []*Person, db *gorm.DB) {
//	result := db.Create(person)
//	errorCheck(result.Error)
//	for i := range person {
//		fmt.Println(person[i].ID)
//	}
//}

//func first(db *gorm.DB) {
//	var person Person
//	db.First(&person)
//	fmt.Println(person)
//}
