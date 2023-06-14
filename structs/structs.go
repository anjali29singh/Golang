package main

import (
	"fmt"
)

// In golang in structs there is no concept of inheritance, parent, child

//function in structs

func main() {

	anjali := User{"anjali", "anjali@gmail.com", 20, true,
		Address{
			"Bhopal", "MP", "Wallstreet", 34,
		},
	}

	fmt.Println(anjali)
	fmt.Printf("detail of anjali is %+v \n", anjali) //+v when structs are used
	fmt.Printf("detail info of anjali is %+v \n", anjali)

}

type Address struct {
	City    string
	State   string
	Street  string
	HouseNo int
}
type User struct {
	Name    string
	Email   string
	Age     int
	Status  bool
	Address Address
}
