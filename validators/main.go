package main

import (
	"fmt"
	"github.com/go-playground/validator"
)

type User struct {
	FName string `validate:"alpha"`
	LName string `validate:"alpha"`
	Age   int    `validate:"gte=20,lte=65"`
	Email string `validate:"required,email"`
	Pwd   string `validate:"required"`
	Repwd string `validate:"required,eqfield=Pwd"`
	//ImgUrl string `validate:"URI"`
	//JoiningDate string `validate:"datetime"`
	//JoiningDate string `validate:"ltecsfield=InnerStructField.StartDate"`

}

var validate *validator.Validate

func main() {
	validate = validator.New()

	user := &User{
		FName: "abc",
		LName: "xyz",
		Age:   45,
		Email: "abc.xyz@gmail.com",
		Pwd:   "abc123",
		Repwd: "abc123",
		//ImgUrl: "https://www.w3schools.com/",
		//JoiningDate: "02-02-2000",
	}

	err := validate.Struct(user)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		fmt.Println("------- List of tag fields with error -------")

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.StructField())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println("---------------")
		}
		return
	}
}
