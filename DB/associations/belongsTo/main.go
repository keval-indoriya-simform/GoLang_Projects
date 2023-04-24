package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model

	Name      string
	CompanyID int
	Company   Company `gorm:"references:id"`
}

type Company struct {
	gorm.Model

	Name string
}

var (
	user = User{
		Name:      "meet",
		CompanyID: 2,
	}

	//cmp = Company{
	//	Name: "tcs",
	//}
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=belongsTo sslmode=disable password=Ke^@l081001 port=5432")

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

	err = db.AutoMigrate(&Company{}, &User{})
	errorCheck(err)

	//db.Create(&cmp)
	//fmt.Println(cmp.ID)
	//db.Create(&user)
	//fmt.Println(user.ID)

	db.Create(&User{
		Name: "Manish",
		Company: Company{
			Name: "isro",
		},
	})

	var getUser []User
	db.InnerJoins("Company").Find(&getUser)
	if err != nil {
		log.Fatal(err)
	}
	println("USER")
	for i := range getUser {
		fmt.Println(getUser[i])
	}

	db.Where("id = ")
	//db.Where("id = 3").Delete(&Company{})
}
