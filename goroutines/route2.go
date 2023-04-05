package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex
var signals = []string{}

func main() {
	websitelist := []string{
		"https://lco.dev",
		"https://google.com",
		"https://go.dev",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("endpoint error")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
