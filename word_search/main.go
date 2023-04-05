package main

import (
	"fmt"
	"strings"
)

const corpos = "lazy cat jumps again and again and again"

func main() {
	var query string
	fmt.Scanln(&query)
	query = strings.ToLower(query)
	queries := strings.Split(query, ",")
	words := strings.Fields(corpos)

	for _, q := range queries {
		index := 0
		for i, word := range words {
			if q == word {
				index = i + 1
				fmt.Printf("%q is at position %d\n", word, index)
				break
			}
		}

		if index == 0 {
			fmt.Printf("%q is not found\n", q)
			continue
		}
	}

}
