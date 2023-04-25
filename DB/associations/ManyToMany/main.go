package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID        int
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	ID    int
	Name  string
	Users []User `gorm:"many2many:user_languages;"`
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=ManyToMany sslmode=disable password=Ke^@l081001 port=5432")

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

	err = db.AutoMigrate(&Language{}, &User{})
	errorCheck(err)

	//db.Create(&User{
	//	Name: "Keval",
	//	Languages: []Language{
	//		{Name: "English"},
	//		{Name: "Hindi"},
	//		{Name: "Gujarati"},
	//	},
	//})
	//db.Create(&User{
	//	Name: "Meet",
	//	Languages: []Language{
	//		{ID: 1, Name: "English"},
	//		{ID: 2, Name: "Hindi"},
	//	},
	//})
	//db.Create(&User{
	//	Name: "Hari",
	//	Languages: []Language{
	//		{ID: 1, Name: "English"},
	//		{ID: 3, Name: "Gujarati"},
	//	},
	//})

	fmt.Println("users")
	GetAllUsers(db)
	fmt.Println("language")
	GetAllLanguage(db)
}

func GetAllUsers(db *gorm.DB) {
	var users []User
	db.Model(&User{}).Preload("Languages").Find(&users)

	for i := range users {
		fmt.Println(users[i])
	}
}

func GetAllLanguage(db *gorm.DB) {
	var languages []Language
	db.Model(&Language{}).Preload("Users").Find(&languages)

	for i := range languages {
		fmt.Println(languages[i])
	}
}
