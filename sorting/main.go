package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.First, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func main() {
	xi := []int{1, 12, 8, 80, 50, 40}
	xs := []string{"abc", "aba", "xyz", "x"}

	p1 := Person{"keval", 25}
	p2 := Person{"abc", 25}
	p3 := Person{"aba", 21}
	p4 := Person{"abc", 25}
	p5 := Person{"keval", 21}
	people := []Person{p1, p2, p3, p4, p5}

	sort.Ints(xi)
	fmt.Println(xi)
	sort.Strings(xs)
	fmt.Println(xs)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)
}
