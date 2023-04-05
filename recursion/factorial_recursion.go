package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number for finding factorial")
	fmt.Scan(&num)
	fmt.Println("Factorial of", num, "is:",factorial(num))
}

func factorial(num int) int {
	if num == 1 {
		return 1
	} else {
		num *= factorial(num - 1)
	}
	return num
}
