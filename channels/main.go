package main

import (
	"fmt"
	"sync"
)

func main() {
	mych := make(chan int, 3)
	mych2 := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(4)
	go func(ch chan int, wg *sync.WaitGroup) {
		x1, isChanelOpen := <-mych
		fmt.Println(isChanelOpen)
		fmt.Println(x1)
		// x2, isChanelOpen := <-mych
		// fmt.Println(isChanelOpen)
		// fmt.Println(x2)
		// x3, isChanelOpen := <-mych
		// fmt.Println(isChanelOpen)
		// fmt.Println(x3)
		wg.Done()
	}(mych, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		mych <- 4
		mych <- 6
		mych <- 5
		// close(mych)
		wg.Done()
	}(mych, wg)

	// read ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		x := <-mych
		fmt.Println(x)
		wg.Done()
	}(mych2, wg)

	// write ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		mych <- 10
		wg.Done()
	}(mych2, wg)

	wg.Wait()
}

// output :-
// true
// 10
// 4
