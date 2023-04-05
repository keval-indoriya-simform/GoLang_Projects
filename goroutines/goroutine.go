package main

import (
	"fmt"
	"time"
)

func main() {
	// just write go in front of functions to make them goroutines
	go greet("Hello") 
	greet("world")
}

func greet(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}
