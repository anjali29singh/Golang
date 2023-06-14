package main

import (
	"fmt"
)

// In golang in structs there is no concept of inheritance, parent, child

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

// methods
// getAge() method will act on u of type User
func (u User) getAge() int {
	return u.Age

}

func (u *User) setAge(newage int) {
	u.Age = newage
}
func main() {

	anjali := User{"anjali", "anjali@gmail.com", 20, true,
		Address{
			"Bhopal", "MP", "Wallstreet", 34,
		},
	}

	fmt.Println(anjali)
	fmt.Printf("detail of anjali is %+v \n", anjali) //+v when structs are used
	fmt.Printf("detail info of anjali is %+v \n", anjali)
	fmt.Println(anjali.getAge())

	anjali.setAge(21)
	fmt.Println(anjali.getAge())
}
