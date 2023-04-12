package main

import "fmt"

func main() {
	var empty interface{}

	empty = map[int]bool{1: true, 2: false}
	empty = 3

	empty = empty.(int) * 2
	fmt.Println(empty)

	empty = []int{1, 2, 3, 4, 5}
	length := len(empty.([]int))
	fmt.Println(length)

	//var any []interface{}

	//for _, n := range empty.([]int) {
	//	any = append(any, n)
	//}
	//fmt.Println(any)

	empty = "hello"
	empty = []int{1, 2, 3, 4, 5}

	switch empty.(type) {
	case int:
		fmt.Println("Int ->", empty)
	case string:
		fmt.Println("String ->", empty)
	case map[int]bool:
		fmt.Println("Map[int]bool ->", empty)
	case []int:
		fmt.Println("Slice Int ->", empty)

	}

}
