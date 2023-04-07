package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Println("enter 1st number")
	fmt.Scan(&num1)
	isZero(&num1)

	fmt.Println("enter 2nd number")
	fmt.Scan(&num2)
	isZero(&num2)

	fmt.Println("Before swap :\n number 1 :", num1, "| number 2 :", num2)

	num1 += num2
	num2 = num1 - num2
	num1 -= num2

	fmt.Println("After swap :\n number 1 :", num1, "| number 2 :", num2)
}

func isZero(num *int) {
	if *num == 0 {
		fmt.Println("You Entered Zero please enter something else")
		fmt.Scan(num)
	}
}
