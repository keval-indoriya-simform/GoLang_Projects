package main

import "fmt"

func main() {
	var continents map[int]string = make(map[int]string)
	continents[1] = "Asia"
	continents[2] = "Africa"
	continents[3] = "North America"
	continents[4] = "South America"
	continents[5] = "Antarctica"
	continents[6] = "Europe"
	continents[7] = "Oceania"

	for i, j := range continents {
		fmt.Printf("Key: %d Value: %s\n", i, j)
	}
}
