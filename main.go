package main

import (
	"fmt"
)

func main() {
	//Println outputs many values separated by space and new line at the end
	/*fmt.Println("hey this is tory", "form usa")
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
	fmt.Println(username, passwd)

	//unintialized var then default values as 0 for int, "" for strings

	var n int
	var s string

	a := 20
	fmt.Println(n, s)
	fmt.Printf("the type of a is %T \n", a)

	*/

	//pointers

	x := 2
	var p *int = &x
	fmt.Println("address of int x is", &x)
	fmt.Println("value of pointer p is", p)
	fmt.Println(*p)

	//indirect change value of x
	*p = 5

	fmt.Println(x, *p)
	x = 20
	fmt.Println(x, *p)
	//another way to create pointer

	pt := new(int) //address is located

	qt := new(int)

	fmt.Println(*pt)

	fmt.Println(qt == pt)

	*pt = 30

	fmt.Println(*pt)
	//each new (type) is different

	num1 := 24
	num2 := 45
	num1, num2 = num2, num1     //tuple assignment
	num1, num2 = num2+6, num1+5 //first RHS is evaluated and then assigned to LHS variable

	fmt.Println(num1, num2)
}
