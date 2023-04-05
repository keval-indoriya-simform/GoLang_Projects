package main

import "fmt"

func main() {
	fmt.Println("How many numbers you want to enter?")
	var length_of_arr int
	fmt.Scan(&length_of_arr)

	if length_of_arr == 0 {
		fmt.Println("You did something wrong")
	} else {
		var array = make([]int, length_of_arr)
		for i := 1; i <= length_of_arr; i++ {
			fmt.Printf("Enter %dth element: ", i)
			fmt.Scan(&array[i-1])
		}
		max := 0
		for i := 0; i < length_of_arr; i++ {
			if array[i] > max {
				max = array[i]
			}
		}
		fmt.Println("Maximum of Numbers in array is: ", max)
	}
}
