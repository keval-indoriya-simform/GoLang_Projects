package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxTurns = 5

func main() {
	var num int

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(11)
	for turn := 0; turn != maxTurns; turn++ {
		// n := 5
		fmt.Println("enter a random number : ")
		fmt.Scan(&num)
		if n == num {
			if turn == 0 {
				fmt.Println("congo!! You got right in first guess")
				return
			}
			switch num {
			case 0, 1, 2:
				fmt.Println("You Win!!")
			case 3, 4, 5:
				fmt.Println("Great!!")
			case 6, 7, 8:
				fmt.Println("Awesone...")
			case 9, 10:
				fmt.Println("Bravo")
			}
			fmt.Println("")
			return
		} else if n > num {
			fmt.Println("Your guess is lesser than number")
		} else if n < num {
			fmt.Println("Your guess is higher than number")
		}

		switch num {
		case 0, 1, 2:
			fmt.Println("You Lost???")
		case 3, 4, 5:
			fmt.Println("Oops!!")
		case 6, 7, 8:
			fmt.Println("Loooser")
		case 9, 10:
			fmt.Println("Try Again")
		}
		fmt.Println("")
	}

}
