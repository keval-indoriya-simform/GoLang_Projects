package main

import (
	"fmt"
)

func main() {
	fmt.Println("How many lines you want in pyramid : ")
	var end, temp, temp1, k int
	fmt.Scan(&end)
	if end != 0 {
		for i := 1; i <= end; i++ {
			for j := 1; j <= end-i; j++ {
				fmt.Print("  ")
				temp++
			}
			for {
				if temp <= end-1 {
					fmt.Printf(" %d", i+k)
					temp++
				} else {
					temp1++
					fmt.Printf(" %d", (i + k - 2*temp1))
				}
				k++

				if k == 2*i-1 {
					break
				}

			}
			temp = 0
			temp1 = 0
			k = 0
			fmt.Println()
		}
	} else {
		fmt.Println("You Entered something wrong")
	}
}
