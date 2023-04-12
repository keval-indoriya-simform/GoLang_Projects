package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mut      sync.Mutex
	counters map[string]int
}

func (c *Container) incr(name string) {
	// it will lock the variable while this goroutine is writing something on that
	c.mut.Lock()
	// with defer at last it will unlock the variable
	defer c.mut.Unlock()
	c.counters[name]++
}

func main() {

	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	wg := &sync.WaitGroup{}

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.incr(name)
		}
		// this will tell wait group that I am Responded and my task is over
		wg.Done()
	}

	// three goroutines are added in wait group
	wg.Add(3)
	go doIncrement("a", 100)
	go doIncrement("b", 200)
	go doIncrement("c", 100)

	// it will ask main method to wait for goroutines to came back and respond
	wg.Wait()
	fmt.Println(c.counters)
}
