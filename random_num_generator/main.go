package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var num int
	fmt.Println("enter a random number : ")
	fmt.Scan(&num)

	var uniques []int

	loop:
		for len(uniques) < num {
			n := rand.Intn(num) + 1
			fmt.Print(n, " ")

			for _, u := range uniques {
				if u == n {
					continue loop
				}
			}

			uniques = append(uniques, n)
		}
		fmt.Println("\nuniques:", uniques)
		
		sort.Ints(uniques)
		fmt.Println("sorted uniques:", uniques)
}
