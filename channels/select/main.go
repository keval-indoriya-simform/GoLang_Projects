package main

import (
	"fmt"
	"time"
)

func main() {
	cha1 := make(chan string)
	cha2 := make(chan string)

	go func() {
		// After 2 sec message is passed through channel
		time.Sleep(2 * time.Second)
		cha1 <- "Hello"
	}()

	go func() {
		// After 2 sec message is passed through channel
		time.Sleep(2 * time.Second)
		cha2 <- "everyone"
	}()

	for i := 0; i < 2; i++ {
		// Select wait for any go routine to come back whhich ever
		// comes back first will get served first
		select {
		case msg1 := <-cha1:
			fmt.Println(msg1)
		case msg2 := <-cha2:
			fmt.Println(msg2)
		}
	}
}

// output :-
// Hello
// everyone
