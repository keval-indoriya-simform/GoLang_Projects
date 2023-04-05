package main

import "fmt"

func main() {
	t1 := 0
	t2 := 1
	nextTerm := 0

	fmt.Println("Enter the number of terms you want : ")
	var end int
	fmt.Scan(&end)
	if end != 0 {
		fmt.Println("Fibonacci Series : ")
		for i := 1; i <= end; i++ {
			if i == 1 {
				fmt.Print(" ", t1)
				continue
			}
			if i == 2 {
				fmt.Print(" ", t2)
				continue
			}
			nextTerm = t1 + t2
			t1 = t2
			t2 = nextTerm
			fmt.Print(" ", nextTerm)
		}
		fmt.Println()
	} else {
		fmt.Println("You Entered something wrong")
	}
}
