package main

import "fmt"

func adder() func(int) int{
	x := 0
	return func(y int) int {
		x += y
		return x
		// fmt.Println(x)
	}
}

func main() {
	a := adder()
	fmt.Println(a(10))
	fmt.Println(a(20))
	fmt.Println(a(40))
	
}

// output :-
// 10
// 30
// 70