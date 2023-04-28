package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Person struct {
	ID    int
	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

type Book struct {
	ID       int
	Title    string
	Author   string
	PersonID int
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=Hooks sslmode=disable password=Ke^@l081001 port=5432")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: true,
	})

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

	err = db.AutoMigrate(&Person{}, &Book{})
	errorCheck(err)

	//var person Person
	//person = Person{
	//	Name:  "Abhishek",
	//	Email: "abhishek.m@simformsolutions.com",
	//	Books: []Book{
	//		{
	//			Title:  "basic",
	//			Author: "basic",
	//		},
	//		{
	//			Title:  "advance",
	//			Author: "advance",
	//		},
	//	},
	//
	//	//Name:  "Keval",
	//	//Email: "keval.i@simformsolutions.com",
	//	//Books: []Book{
	//	//	{
	//	//		Title:  "basic of golang",
	//	//		Author: "google",
	//	//	},
	//	//},
	//
	//	//Name:  "Kishan",
	//	//Email: "kishan.m@simformsolutions.com",
	//	//Books: []Book{
	//	//	{
	//	//		Title:  "basic",
	//	//		Author: "basic",
	//	//	},
	//	//	{
	//	//		Title:  "advance",
	//	//		Author: "advance",
	//	//	},
	//	//},
	//}
	//CreateUser(db, &person)
	//DeleteBook(db)
	var user Person
	db.Find(&user, 2)
	//fmt.Println(user)
	db.Model(&user).Update("name", "abhi4")
}

func (p *Person) AfterSave(tx *gorm.DB) error {
	fmt.Println(*p)
	return nil
	//return gorm.ErrRecordNotFound
}

func (p *Person) AfterFind(db *gorm.DB) error {
	var books []Book
	db.Model(p).Association("Books").Find(&books)
	p.Books = books
	fmt.Println(*p)

	//db.Model(p).Preload("Books").Find(p)
	return nil
}

//	func (p *Person) AfterSave(tx *gorm.DB) error {
//		fmt.Println(*p)
//		return nil
//	}
//
//	func (p *Person) BeforeSave(tx *gorm.DB) error {
//		fmt.Println(*p)
//		return nil
//	}

//func (p *Person) BeforeCreate(tx *gorm.DB) error {
//	fmt.Println(*p)
//	return nil
//	//return errors.ErrKeyIncorrect
//}

func (b *Book) BeforeDelete(db *gorm.DB) error {
	var Books Book
	db.Where("ID = ?", b.ID).Find(&Books)
	fmt.Println("After Delete")
	fmt.Println(Books)
	return nil
	//return gorm.ErrRecordNotFound
}

func (b *Book) AfterDelete(db *gorm.DB) error {
	var Books Book
	db.Where("ID = ?", b.ID).Find(&Books)
	fmt.Println("After Delete")
	fmt.Println(Books)
	return nil
	//return gorm.ErrRecordNotFound
}

func DeleteBook(db *gorm.DB) error {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	tx.Delete(&Book{ID: 3})

	return tx.Commit().Error
}

func CreateUser(db *gorm.DB, p *Person) error {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(p).Error; err != nil {
		return err
	}

	//tx.SavePoint("user1")
	//tx.RollbackTo("user1")

	return tx.Commit().Error
}
