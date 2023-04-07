package main

import (
	"fmt"
)

func main() {
	fmt.Println("How many lines you want in pyramid : ")
	var end int
	fmt.Scan(&end)

	if end != 0 {
		for i := 0; i < end; i++ {
			for k := 1; k < end-i; k++ {
				fmt.Print(" ")
			}
			for j := 0; j < i+1; j++ {
				fmt.Print("* ")
			}
			fmt.Println()
		}
		fmt.Println()
	} else {
		fmt.Println("You Entered something wrong")
	}
}
