package main

import "fmt"

type (
	bookcase [5]int
	cabinet  [5]int
)

func main() {
	book := bookcase{6, 9, 3, 2, 1}
	cab := cabinet{6, 9, 3, 2, 1}

	fmt.Println("Are they equal? ")

	if book == bookcase(cab) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	fmt.Printf("book: %#v\n", book)
	fmt.Printf("cab: %#v\n", cab)
	fmt.Printf("cab: %#v\n", bookcase(cab))
}
