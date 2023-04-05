package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mut := &sync.Mutex{}
	mut1 := &sync.RWMutex{}

	var score = []int{0}
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("routeine 1")
		mut.Lock()
		defer mut.Unlock()
		score = append(score, 1)
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("routeine 2")
		mut.Lock()
		defer mut.Unlock()
		score = append(score, 2)
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("routeine 3")
		mut.Lock()
		defer mut.Unlock()
		score = append(score, 3)
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		time.Sleep(20*time.Millisecond)
		fmt.Println("routeine 4")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut1)

	wg.Wait()
	defer fmt.Println(score)
}
