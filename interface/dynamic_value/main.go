package main

import "fmt"

type printer interface {
	print()
}

type book struct {
	title string
	price float64
}

func (b book) print() {
	fmt.Println(b.title, ":", b.price)
}

type game struct {
	title string
	price float64
}

func (b game) print() {
	fmt.Println(b.title, ":", b.price)
}

func (b *game) discount(ratio float64) {
	b.price = b.price * ratio
}

type puzzle struct {
	title string
	price float64
}

func (b puzzle) print() {
	fmt.Println(b.title, ":", b.price)
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("sorry, we are waiting for delivery.")
		return
	}

	for _, it := range l {
		it.print()
	}
}

func (l list) discount(ratio float64) {
	type discounter interface {
		discount(float64)
	}

	for _, it := range l {
		if it, ok := it.(discounter); ok {
			it.discount(ratio)
		}
	}
}

func main() {
	var (
		books     = book{title: "book1", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		GOW       = game{title: "gow", price: 15}
		rubik     = puzzle{title: "rubik", price: 5}
	)

	var store list
	store = append(store, &minecraft, &GOW, books, rubik)

	store.discount(.5)
	minecraft.print()
	store.print()
}
