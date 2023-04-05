package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter number : ")
	var num int
	fmt.Scanln(&num)
	if num == 0 {
		fmt.Println("You Entered something wrong!")
	} else {
		if num%2 == 0 {
			fmt.Println(num, "is Even number")
		} else {
			fmt.Println(num, "is Odd number")
		}
	}

}
