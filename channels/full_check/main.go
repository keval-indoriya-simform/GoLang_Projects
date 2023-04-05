package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	send(2, ch)
	send(3, ch)
	send(4, ch)
	send(5, ch)
	send(1, ch)

	fmt.Println("Channel content")
	for i := 0; i < cap(ch); i++ {
		fmt.Println(<-ch)
	}
}

func send(num int, ch chan int) {
	if len(ch) == cap(ch) {
		fmt.Println("Channel is full")
	} else {
		fmt.Println(num, "Inserted in channel")
		ch <- num
	}
}
