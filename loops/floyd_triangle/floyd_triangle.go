package main

import "fmt"

func main() {
	var end int
	fmt.Print("Enter how many lines you want : ")
	fmt.Scan(&end)
	num := 1
	for i := 1; i <= end; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(num, " ")
			num++
		}
		fmt.Println("")
	}
}
