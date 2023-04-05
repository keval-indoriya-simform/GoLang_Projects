package main

import "fmt"

func main() {
	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mukemmel",
	}

	var word string
	fmt.Println("enter word from [good, great, perfact] : ")
	fmt.Scan(&word)

	if value, ok := dict[word]; ok {
		fmt.Printf("%q means %#v\n", word, value)
	} else {
		fmt.Printf("%q not found\n", word)
	}
}
