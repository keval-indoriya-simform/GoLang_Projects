package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	ID          int
	Name        string
	CreditCards []CreditCard `gorm:"constraint:onUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreditCard struct {
	ID     int
	Number string
	UserID int
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=HasMany sslmode=disable password=Ke^@l081001 port=5432")

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

	err = db.AutoMigrate(&User{}, &CreditCard{})
	errorCheck(err)

	db.Create(&[]User{
		{
			Name: "Keval",
			CreditCards: []CreditCard{
				{
					Number: "555544443333",
				},
				{
					Number: "123456789123",
				},
			},
		},
		{
			Name: "Meet",
			CreditCards: []CreditCard{
				{
					Number: "111144443333",
				},
			},
		},
		{
			Name: "Hari",
			CreditCards: []CreditCard{
				{
					Number: "222244443333",
				},
			},
		},
		{
			Name: "Juhi",
			CreditCards: []CreditCard{
				{
					Number: "333344443333",
				},
			},
		},
	})

	GetAll(db)
	//var user User
	//db.First(&user, "id = 2").Delete(&user)
	//fmt.Println("deleted")
	//GetAll(db)
}

func GetAll(db *gorm.DB) {
	var users []User
	err := db.Model(&User{}).Preload("CreditCards").Find(&users).Error
	errorCheck(err)
	for i := range users {
		fmt.Println(users[i])
	}
}
