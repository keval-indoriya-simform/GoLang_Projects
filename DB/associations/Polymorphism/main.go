package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Cat struct {
	ID   int
	Name string
	Toys Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
	ID   int
	Name string
	Toys Toy `gorm:"polymorphic:Owner;polymorphicValue:master"`
}

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=Polymorphism sslmode=disable password=Ke^@l081001 port=5432")

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

	err = db.AutoMigrate(&Toy{}, &Cat{}, &Dog{})
	errorCheck(err)

	db.Create(&Dog{Name: "Dog1", Toys: Toy{Name: "toy1"}})
	db.Create(&Cat{Name: "Cat1", Toys: Toy{Name: "toy2"}})
	db.Create(&Dog{Name: "Dog2", Toys: Toy{Name: "toy3"}})
	db.Create(&Cat{Name: "Cat2", Toys: Toy{Name: "toy4"}})

	//db.Create(&Dog{Name: "Dog1", Toys: []Toy{{Name: "toy1"}, {Name: "toy2"}}})
	//db.Create(&Cat{Name: "Cat1", Toys: []Toy{{Name: "toy3"}, {Name: "toy4"}}})
	//db.Create(&Dog{Name: "Dog2", Toys: []Toy{{Name: "toy5"}, {Name: "toy6"}}})
	//db.Create(&Cat{Name: "Cat2", Toys: []Toy{{Name: "toy7"}, {Name: "toy8"}}})

	var dogs []Dog
	var cats []Cat
	db.Joins("Toys").Find(&dogs)
	db.Joins("Toys").Find(&cats)
	//var dogs []Dog
	//var cats []Cat
	//err = db.Model(&Dog{}).Preload("Toys").Find(&dogs).Error
	//errorCheck(err)
	//err = db.Model(&Cat{}).Preload("Toys").Find(&cats).Error
	//errorCheck(err)

	fmt.Println("dogs")
	fmt.Println(dogs)
	fmt.Println("cats")
	fmt.Println(cats)

}
