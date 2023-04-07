package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numString string
	fmt.Scanln(&numString)
	numbers := strings.Split(numString, ",")
	var numSlice []int
	var primeSlice []int

	for _, number := range numbers {
		num, err := strconv.Atoi(number)
		if err == nil {
			numSlice = append(numSlice, num)
		}
	}

	for _, num := range numSlice {
		if isPrime(num) {
			primeSlice = append(primeSlice, num)
		}
	}
	fmt.Println("list of prime numbers:", primeSlice)
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
