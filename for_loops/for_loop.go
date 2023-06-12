package main

import (
	"fmt"
)

func main() {

	//for init,condition,post{}

	for i := 0; i < 6; i++ {
		fmt.Println(i)
	}

	//infinite for loop

	// for {
	// 	fmt.Println(1)
	// }

	//iterating map ,array,slice,string

	var str string = "anjali"

	for key, char := range str {
		fmt.Printf("%v %c \n", key, char)
	}
	//if no need of second param

	for key := range str {
		fmt.Println(key)
	}
	//blank identifier : if no nedd of 1st param

	for _, char := range str {
		fmt.Printf("%c \n", char)
	}

}
