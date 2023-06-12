package main

import (
	"fmt"
)

func main() {
	//Println outputs many values separated by space and new line at the end
	fmt.Println("hey this is tory", "form usa")
	fmt.Println("hello World")
	// fmt.Printf("hellow world")
	// fmt.Printf(23) //error

	//for loops

	for i := 0; i < 6; i++ {

		fmt.Println(i)
	}
	//infinite for loop
	// for {
	// 	fmt.Print(1)
	// }

	// _ blank identifier for unused var

	print("******************************* variable declarations **************************** \n")

	// var name type = value;
	var username, passwd string = "anjali", "2354434"
	fmt.Println(username, " ", passwd)

	//unintialized var then default values as 0 for int, "" for strings

	var n int
	var s string

	a := 20
	fmt.Println(n, " ", s)
	fmt.Printf("the type of a is %T ", a)
}
