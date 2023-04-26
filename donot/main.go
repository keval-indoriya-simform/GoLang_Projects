package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Address struct {
	ID        int
	CountryID int
	Country   Country
}

type Country struct {
	ID   int
	Name string
}

type Org struct {
	PostalAddress   Address `gorm:"embedded;embeddedPrefix:postal_address_"`
	VisitingAddress Address `gorm:"embedded;embeddedPrefix:visiting_address_"`
	AddressID       int
	Address         Address
}

func main() {

	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=donot sslmode=disable password=Ke^@l081001 port=5432")

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

	err = db.AutoMigrate(&Country{}, &Address{}, &Org{})
	errorCheck(err)

	org := Org{
		PostalAddress: Address{
			Country: Country{
				Name: "India",
			},
		},
		VisitingAddress: Address{
			Country: Country{
				Name: "Canada",
			},
		},
	}

	db.Create(&org)

	var orgs []Org
	// Only preload Org.Address and Org.Address.Country
	db.Preload("Address.Country").Find(&orgs) // "Address" is has_one, "Country" is belongs_to (nested association)
	fmt.Println(orgs)
	// Only preload Org.VisitingAddress
	db.Preload("PostalAddress.Country").Find(&orgs) // "PostalAddress.Country" is belongs_to (embedded association)
	fmt.Println(orgs)

	// Only preload Org.NestedAddress
	db.Preload("NestedAddress.Address.Country").Find(&orgs) // "NestedAddress.Address.Country" is belongs_to (embedded association)
	fmt.Println(orgs)

	// All preloaded include "Address" but exclude "Address.Country", because it won't preload nested associations.
	db.Preload(clause.Associations).Find(&orgs)
	fmt.Println(orgs)

}
