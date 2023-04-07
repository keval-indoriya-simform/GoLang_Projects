package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var num int
	fmt.Print("Enter number that you want to check if it is armstrong or not : ")
	fmt.Scan(&num)
	if isAramstrong(num) {
		fmt.Println("Yes!!, It is Armstrong number")
	} else {
		fmt.Println("No!!, it is not Armstrong number")
	}
}

func isAramstrong(num int) bool {
	if num >= 100 {
		var str string = strconv.Itoa(num)
		sum := 0
		length := len(str)
		for i := 0; i < length; i++ {
			tempNum, _ := strconv.ParseInt(string(str[i]), 10, 64)
			sum += int(math.Pow(float64(tempNum), float64(length)))
		}
		if sum == num {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
