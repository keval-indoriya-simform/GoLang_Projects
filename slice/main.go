package main

import "fmt"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < len(slice) && index >= 0 {
		return slice[index]
	} else {
		return -1
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < len(slice) && index >= 0 {
		slice[index] = value
	} else if index == len(slice) {
		slice = append(slice, value)
	} else if index < 0 {
		slice = append(slice, value)
	}
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	var s1 = []int{}
	s1 = append(s1, values...)
	s1 = append(s1, slice...)
	return s1
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < len(slice) && index >= 0 {
		slice = append(slice[:index], slice[index+1:]...)
	}
	return slice
}

func main() {
	slice := FavoriteCards()
	fmt.Println(GetItem(slice, 1))
	fmt.Println(slice)
	slice = SetItem(slice, -1, 10)
	fmt.Println(slice)
	slice = SetItem(slice, 1, 7)
	fmt.Println(slice)
	
	slice = PrependItems(slice, 5, 3, 1)
	fmt.Println(slice)
	slice = RemoveItem(slice, 5)
	fmt.Println(slice)
}
