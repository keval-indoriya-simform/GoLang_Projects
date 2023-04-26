package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type User struct {
	ID              int
	Name            string
	BillingAddress  Address
	ShippingAddress Address
	Emails          []Email
	Languages       []Language `gorm:"many2many:user_languages;"`
}

type Email struct {
	ID     int
	Mail   string
	UserID int
}

type Address struct {
	ID          int
	AddressType string
	Address     string
	UserID      int
}

type Language struct {
	ID   int
	Name string
	//Users []User `gorm:"many2many:user_languages;"`
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=AssociationMode sslmode=disable password=Ke^@l081001 port=5432")

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

	err = db.AutoMigrate(&Language{}, &User{}, &Email{}, &Address{})
	errorCheck(err)

	//var users []User
	//
	//users = []User{
	//	{
	//		Name: "Keval",
	//		BillingAddress: Address{
	//			AddressType: "Billing",
	//			Address:     "Ahmedabad",
	//		},
	//		ShippingAddress: Address{
	//			AddressType: "Shipping",
	//			Address:     "Rajkot",
	//		},
	//		Emails: []Email{
	//			{Mail: "keval.i@simformsolutions.com"},
	//			{Mail: "kevalindoriya5@gmail.com"},
	//		},
	//		Languages: []Language{
	//			{Name: "English"},
	//			{Name: "Hindi"},
	//			{Name: "Gujarati"},
	//		},
	//	},
	//	{
	//		Name: "Meet",
	//		BillingAddress: Address{
	//			AddressType: "Billing",
	//			Address:     "Rajkot",
	//		},
	//		ShippingAddress: Address{
	//			AddressType: "Shipping",
	//			Address:     "Ahmedabad",
	//		},
	//		Emails: []Email{
	//			{Mail: "meetindoriya5@gmail.com"},
	//		},
	//		Languages: []Language{
	//			{ID: 1, Name: "English"},
	//			{ID: 2, Name: "Hindi"},
	//			{ID: 3, Name: "Gujarati"},
	//		},
	//	},
	//	{
	//		Name: "Hari",
	//		BillingAddress: Address{
	//			AddressType: "Billing",
	//			Address:     "Ahmedabad",
	//		},
	//		ShippingAddress: Address{
	//			AddressType: "Shipping",
	//			Address:     "Jamnagar",
	//		},
	//		Emails: []Email{
	//			{Mail: "harishindoriya5@gmail.com"},
	//		},
	//		Languages: []Language{
	//			{ID: 1, Name: "English"},
	//			{ID: 2, Name: "Hindi"},
	//			{ID: 3, Name: "Gujarati"},
	//		},
	//	},
	//	{
	//		Name: "Juhi",
	//		BillingAddress: Address{
	//			AddressType: "Billing",
	//			Address:     "Ahmedabad",
	//		},
	//		ShippingAddress: Address{
	//			AddressType: "Shipping",
	//			Address:     "Ahmedabad",
	//		},
	//		Emails: []Email{
	//			{Mail: "juhi.m@simformsolutions.com"},
	//		},
	//		Languages: []Language{
	//			{ID: 1, Name: "English"},
	//			{ID: 2, Name: "Hindi"},
	//			{ID: 3, Name: "Gujarati"},
	//		},
	//	},
	//	{
	//		Name: "Abhishek",
	//		BillingAddress: Address{
	//			AddressType: "Billing",
	//			Address:     "Ahmedabad",
	//		},
	//		ShippingAddress: Address{
	//			AddressType: "Shipping",
	//			Address:     "Surat",
	//		},
	//		Emails: []Email{
	//			{Mail: "abhishek.m@simformsolutions.com"},
	//		},
	//		Languages: []Language{
	//			{ID: 1, Name: "English"},
	//			{ID: 2, Name: "Hindi"},
	//			{ID: 3, Name: "Gujarati"},
	//		},
	//	},
	//	{
	//		Name: "Kishan",
	//		BillingAddress: Address{
	//			AddressType: "Billing",
	//			Address:     "Ahmedabad",
	//		},
	//		ShippingAddress: Address{
	//			AddressType: "Shipping",
	//			Address:     "Ahmedabad",
	//		},
	//		Emails: []Email{
	//			{Mail: "kishan.m@simformsolutions.com"},
	//		},
	//		Languages: []Language{
	//			{ID: 1, Name: "English"},
	//			{ID: 2, Name: "Hindi"},
	//			{ID: 3, Name: "Gujarati"},
	//		},
	//	},
	//}
	//
	//db.Create(&users)

	var users []User
	//db.Preload("Languages").Preload("Emails").Preload("BillingAddress", "address_type = ?", "Billing").Preload("ShippingAddress", "address_type = ?", "Shipping").Find(&users)
	db.Preload("BillingAddress", "address_type = ?", "Billing").Preload("ShippingAddress", "address_type = ?", "Shipping").Preload(clause.Associations).Find(&users)
	for i := range users {
		fmt.Println(users[i])
	}

}
