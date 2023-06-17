package main

import (
	"fmt"
	"sort"
)

func double(slice []int) {

	for key := range slice {
		slice[key] = 2 * slice[key]

	}
}

func main() {

	//slice is the part of array , its size can increase or decrease
	/* three components of slice
	1.Pointer pointer to first element of array
	2. len :no. elements present in slice
	3.capacity : maximum elements it can occupy
	*/
	Arr := [6]int{23, 45, 12, 89, 45, 10}
	my_slice := Arr[1:4]
	fmt.Println(my_slice, len(my_slice), cap(my_slice))

	//modifying slice will also modify array
	my_slice[0] = 1
	my_slice = append(my_slice, 4, 6, 7, 8)
	fmt.Println(my_slice, Arr)

	slice2 := []int{1, 5, 22, 9, 3}
	slice2 = append(slice2, 10)
	fmt.Println(len(slice2), cap(slice2))
	//sort slice
	sort.Ints(slice2)
	fmt.Println(slice2)

	double(slice2)
	fmt.Println(slice2)
	var slice3 = []int{1, 2, 3, 4, 5, 6, 10}

	//copy(des, source) // minlen(des,source) is copied to des
	x := copy(slice2, slice3)
	fmt.Println(x, slice2, slice3)

}
