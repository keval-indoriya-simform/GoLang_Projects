package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model

	Number string
	UserID uint
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=HasOne sslmode=disable password=Ke^@l081001 port=5432")

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
	//users := []User{
	//	{
	//		Name: "Keval",
	//		CreditCard: CreditCard{
	//			Number: "5555222244443636",
	//		},
	//	},
	//	{
	//		Name: "Meet",
	//		CreditCard: CreditCard{
	//			Number: "2345234523452345",
	//		},
	//	},
	//	{
	//		Name: "Hari",
	//		CreditCard: CreditCard{
	//			Number: "8989898989898989",
	//		},
	//	},
	//}
	//
	//db.Create(&users)

	//db.Create(&CreditCard{
	//	Number: "8989898989898989",
	//	UserID: 3,
	//})
	var users []User
	//db.Find(&users)
	db.Joins("CreditCard").Find(&users)
	for i := range users {
		//db.Model(&users[i]).Association("CreditCard").Find(&users[i].CreditCard)
		fmt.Println(users[i])
	}

	//db.Where("id = 3").Delete(&CreditCard{})
}
