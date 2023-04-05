package main

import (
	"fmt"
	"math"
)

func main() {
	// var num int
	// fact := 1
	// fmt.Println("Enter the number for finding factorial")
	// fmt.Scan(&num)

	// for i := 1; i <= num; i++ {
	// 	fact *= i
	// }

	// fmt.Printf("Factorial of %v is %v\n", num, fact)

	var num float64
	fmt.Println("Enter the number for finding factorial")
	fmt.Scan(&num)

	fmt.Println(math.Gamma(num + 1))
}
