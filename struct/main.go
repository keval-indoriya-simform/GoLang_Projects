package main

import (
	parent "struct/father"
	child "struct/father/son"

	"fmt"
)

func main() {
	f := new(parent.Father)
	fmt.Println(f.Data("ABC"))

	c := new(child.Son)
	fmt.Println(c.Data("XYZ"))
}
