package main

import "fmt"

func main() {

	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4

	//shorthand declaration

	nums := [6]int{2, 4, 6, 8, 10, 12}
	fmt.Println(arr, nums)
	//multi-dimensional array

	matrix := [3][3]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	fmt.Println(matrix)

	var chars [2][3]string
	chars[0][0] = "a"
	chars[0][1] = "b"
	chars[0][2] = "c"
	chars[1][0] = "d"
	chars[1][1] = "e"
	chars[1][2] = "f"
	fmt.Println(chars)

	//... is ellipsis
	lang := [...]string{"c++", "python", "go", "javascript"}
	fmt.Println(len(lang), lang)

}
