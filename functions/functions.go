package main

import (
	"fmt"
)

func rectangle(a int, b int) (area int) {

	area = a * b
	return //retun area

}

func double(a *int, b *int) int {
	*a = 2 * (*a)
	*b = 2 * (*b)

	//closure function is a anonymous function which have access to variables defined outside it
	return func() int {
		var perimeter int = *a * (*b)
		return perimeter
	}()

}

// passing function as an argument
func avg(sum int, n int) int {

	average := sum / n
	return average

}
func sum(arr [5]int) int {
	cnt := 0
	for _, val := range arr {
		cnt = cnt + val
	}
	return cnt
}

/*
defer functions are used to cause delay in function execution .It is executed only

when nearby functions completes its execution. defer function or method call arguments evaluate instantly, but they

don’t execute until the nearby functions returns
. It is used to check files are closed or not

*/

func greet(s string) int {
	fmt.Printf("hello my name is %v \n", s)
	return 0
}
func meet() {
	fmt.Println("How have you been ")
}

func main() {
	fmt.Println(rectangle(23, 15))

	a := 2
	b := 5
	fmt.Println(double(&a, &b))
	var arr = [5]int{1, 2, 3, 4, 5}
	sum := sum(arr)
	fmt.Println(avg(sum, 5))

	//here greet function act as defer
	//if more than one defer function than LIFO rule
	defer greet("anshita")

	meet()

}
