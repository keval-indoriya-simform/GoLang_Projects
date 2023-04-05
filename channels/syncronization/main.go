package main

import (
	"fmt"
	"time"
)

// This goroutine send true in channel after execution
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	// This will make main function wait till go routine send back the responce
	<-done
}

// done will be printed after one second of working is printed
// output :-
// Working...done
