package main

import (
	"fmt"
)

//if variable name starts with Capital letter then it is exported global
/*var (
	username   string = "anjali"
	profession string = "student"
	percent    int    = 80
)*/

func main() {
	//Println outputs many values separated by space and new line at the end
	fmt.Println("hey this is tory", "form usa")
	fmt.Println("hello World")
	// fmt.Printf("hellow world")
	// fmt.Printf(23) //error

	//for loops

	/*for i := 0; i < 6; i++ {

		fmt.Println(i)
	}
	// infinite for loop
	for {
		fmt.Println(1)
	}

	// _ blank identifier for unused var

	fmt.Printf("******************************* variable declarations **************************** \n")

	*/
	// var name type = value;
	var username, passwd string = "anjali", "2354434"
	fmt.Println(username, passwd)

	//unintialized var then default values as 0 for int, "" for strings

	var n int
	var s string

	a := 20
	fmt.Println(n, s)
	fmt.Printf("the type of a is %T \n", a)

	// fmt.Println(username, profession, percent)

	/*fmt.Println("******************************** Pointers ***************************************")
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


	//data type conversion
	a := strconv.Itoa(num1)
	b := strconv.Itoa(num2)
	fmt.Printf(a+b, '\n')
	*/
	fmt.Printf("*********************************** primitives *********************** \n")
	//bool, int,int8,int16,int64 , unsigned uint8,uint16 , string,float32,float64
	var present bool //default value for boolean is false
	fmt.Println(present)
	fmt.Printf("data type of %v is %T \n", 45, 45)

	//bit operators &,|,^ (XOR), &^  NOR

}
