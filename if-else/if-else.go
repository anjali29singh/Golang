package main

import (
	"fmt"
)

const num int = 55

func main() {
	//if condition{

	// }
	if num < 75 {

		fmt.Println("Work hard")
	} else { //else should be at the ending bracket of if statement

		fmt.Println("congrats you have passed exam")
	}

	//switch case

	day := 4
	switch day {
	case 1:
		fmt.Println("Today is Monday")

	case 2:
		fmt.Println("Today is Tuesday")
	case 3:
		fmt.Println("Today is Wednesday")
	case 4:
		fmt.Println("Today is Thursday")
	case 5:
		fmt.Println("Today is Frinday")
	case 6:
		fmt.Println("Today is Saturday")
	case 7:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Please check the day")

	}
	//fallthrough force the execution to fall through the successive code blocks

	n := 3
	switch n {
	case 1, 3:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
		fallthrough
	case 4:
		fmt.Println(3)
	}
}
