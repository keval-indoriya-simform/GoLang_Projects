package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "abc",
		Last:  "abc",
		Age:   20,
	}
	p2 := person{
		First: "xyz",
		Last:  "xyz",
		Age:   20,
	}

	people := []person{p1, p2}
	// people_from_json := []person{}
	var people_from_json []person
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	} else {
		jsonstr := string(bs)
		fmt.Println(jsonstr)
	}

	err = json.Unmarshal(bs, &people_from_json)
	if err != nil {
		fmt.Println(err)
	} else {
		for i, v := range people_from_json {
			fmt.Println("Person Number :", i)
			fmt.Println(v.First, v.Last, v.Age)
		}		
	}
}
