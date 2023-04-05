package main

import (
	"fmt"
)

func main() {

	count := 5
	var result string

	if count < 10 {
		result = "Less than 10"
	} else if count > 10 {
		result = "Greater than 10"
	} else {
		result = "Exectly 10"
	}
	fmt.Println(result)

	if num := 5; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greeater than 10")
	}

	day := 1
	var name_of_day string
	switch day {
	case 1:
		name_of_day = "sunday"
	case 2:
		name_of_day = "monday"
	case 3:
		name_of_day = "tuesday"
	case 4:
		name_of_day = "wednesday"
	case 5:
		name_of_day = "thursday"
	case 6:
		name_of_day = "friday"
	case 7:
		name_of_day = "saterday"
	}

	var isWeekend bool
	switch day {
	case 1, 7:
		isWeekend = true
	case 2, 3, 4, 5, 6:
		isWeekend = false
	}

	fmt.Println(name_of_day)
	fmt.Println(isWeekend)

}
