package main

import (
	"fmt"
)

func main() {
	//new keyword allocate memory and returns pointer , pointing to type
	//T and has value =0 (int ) value="" (strings)
	p := new(int)
	q := new(string)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(*q)

	//make allocate memory and initializes object
	//of type slice,map . return type is not pointer
	x := make([]int, 5)
	fmt.Println(x)

}
