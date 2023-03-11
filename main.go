package main

import (
	"fmt"
	"strconv" //for conversion to strings
)

// error if a variable is defined but not used
// cannot redeclare same variable in the same scope
// In go uninitialized value is 0
var i int = 34
var num float32 = 45
var a string = "abc"
var b int

func main() {
	var i int = 60
	var j float32 = 78
	m := "anjali"
	j = float32(i)
	j = i //error
	var I int = 56
	a = strconv.Itoa(i)
	fmt.Println(num, i, I)
	fmt.Println(j, a)
	fmt.Printf("%v,%T\n", m, m)
	fmt.Printf("%v,%T\n", a, a)

	//primitives bool,int8 ,int16,int32,int64 ,
	// var n bool = false
	// var n uint16 = 34
	var n int16 = -23

	fmt.Printf("%v,%T\n", n, n)
	fmt.Println(b)

	//operation of same data types

	var num1 int = 23
	var num2 int = 45
	// fmt.Println(num1+num2, num1*num2, num2-num1, num1-num2, num2/num1, num2%num1)

	//bit operators
	num1 = 10
	num2 = 3
	fmt.Println(num1&num2, num1|num2, num1^num2, num1&^num2)
	fmt.Println(num2<<3, num1>>2) //num2 * (2^3) num1/(2^2)
}
