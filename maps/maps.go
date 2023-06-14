package main

import (
	"fmt"
)

func main() {

	mp := map[int]string{1: "a",
		2: "b",
		3: "c",
		4: "d"}

	fmt.Println(mp[1])

	//map declaration using make function
	var pets = make(map[int]string)
	pets[1] = "Dogs"
	pets[2] = "Cats"
	pets[3] = "Birds"
	pets[4] = "Rabbits"
	pets[5] = "Cows"
	//iterate over maps
	for id, pet := range pets {
		fmt.Println(id, pet)
	}

	//to check key is present or not
	_, present := mp[9]
	if present {
		fmt.Println("key is present in map")
	} else {
		fmt.Println("key is absent")
	}

	value, is_present := pets[3]

	fmt.Println(value, is_present)

	//delete key
	delete(pets, 4)
	fmt.Println(pets)
}
