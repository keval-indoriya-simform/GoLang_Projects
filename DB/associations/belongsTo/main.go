package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID   int `gorm:"primaryKey,autoIncrement"`
	Name string
	//CompaniesID int `gorm:"foreignKey:companies.id"`
	Companies Company
}

type Company struct {
	ID   int `gorm:"primaryKey,autoIncrement"`
	Code string
	Name string
}

var (
	user = User{
		Name: "Keval",
		//CompanyID: 1,
	}

	cmp = Company{
		Code: "SSL",
		Name: "Simform Software LLP",
	}
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

	db.Create(&cmp)
	fmt.Println(cmp.ID)
	db.Create(&user)
	fmt.Println(user.ID)
}
