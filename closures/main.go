package main

import "fmt"

func counter() func() int {
	// when this function is called for the first time it will initialize count
	// and return function to assigned variable now when this variable function called
	// it will increment counter and don't initialize count again
	// if you make counter funcation calll to some other variable it will again initilize
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	a := counter()
	fmt.Println(a()) //1
	fmt.Println(a()) //2
	fmt.Println(a()) //3
	fmt.Println(a()) //4
	fmt.Println(a()) //5

	b := counter()
	fmt.Println(b()) //1
	fmt.Println(b()) //2

	fmt.Println(a()) //6
}
