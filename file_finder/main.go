package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var dir string
	fmt.Println("enter directory : ")
	fmt.Scan(&dir)

	files, err := os.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
		return
	}

	var names []byte

	for _, file := range files {
		name := file.Name()
		names = append(names, name...)
		names = append(names, '\n')
	}

 	ioutil.WriteFile("new.txt", names, 0644)

	fmt.Printf("%s", names)
}
