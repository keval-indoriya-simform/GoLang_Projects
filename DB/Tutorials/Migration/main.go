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

func main() {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=Hooks sslmode=disable password=Ke^@l081001 port=5432")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		//Logger: logger.Default.LogMode(logger.Info),
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

	//err = db.Migrator().CreateTable(&Name{})
	//fmt.Println(db.Migrator().HasTable(&Name{}))
	//fmt.Println(db.Migrator().HasTable("hello"))
	//db.Migrator().RenameTable("names", "user_names")
	//db.Migrator().RenameTable("user_names", "names")
	//db.Migrator().DropTable(&Name{})
	//db.Migrator().AddColumn(&Name{}, "Address")
	//fmt.Println(db.Migrator().HasColumn(&Name{}, "Address"))
	//db.Migrator().RenameColumn(&Name{}, "Address", "New_Address")
	//db.Migrator().RenameColumn(&Name{}, "New_Address", "Address")
	//db.Migrator().DropColumn(&Name{}, "Address")
	//db.Migrator().AlterColumn(&Name{}, "Name")

	//db.Debug().Create(&[]Name{
	//	{Name: "keval"},
	//	{Name: "meet"},
	//	{Name: "abhishek"},
	//	{Name: "juhi"},
	//})

	var Names []Name
	db.Debug().Find(&Names)
	//fmt.Println(Names)
	//query := db.Model(&Name{}).Where("ID > 2")
	//db.Debug().Migrator().CreateView("user_view", gorm.ViewOption{Query: query})
	//db.Migrator().DropView("user_view")

	//db.Migrator().DropConstraint(&Name{}, "idx_names_name")

	//db.Migrator().CreateConstraint(&Name{}, "name_checker")
	//fmt.Println(db.Migrator().HasConstraint(&Name{}, "name_checker"))
	//db.Migrator().DropConstraint(&Name{}, "name_checker")

	//db.Migrator().CreateIndex(&Name{}, "Name")
	//db.Migrator().CreateIndex(&Name{}, "idx_name")
	//fmt.Println(db.Migrator().HasIndex(&Name{}, "Name"))
	//fmt.Println(db.Migrator().HasIndex(&Name{}, "idx_name"))
	//db.Migrator().DropIndex(&Name{}, "Name")
	//db.Migrator().DropIndex(&Name{}, "idx_name")
	//db.Migrator().DropIndex(&Name{}, "new_index")
	//db.Migrator().RenameIndex(&Name{}, "idx_name", "new_index")

	//errorCheck(err)
}

type Name struct {
	ID   int
	Name string `gorm:"index:idx_name,unique"`
}
