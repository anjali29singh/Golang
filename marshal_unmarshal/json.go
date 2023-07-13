package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//marshaling is encoding  i.e conversion to json
//unmarshaling is decoding i.e conversion to go data structure

type Users struct {
	Users []User
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func main() {

	user1 := User{"anjali29singh", "11111", "anjali@gmail.com"}
	user2 := User{"joe", "2222jd", "joetrib@gmail.com"}

	// var users Users

	data, err := json.Marshal([]User{user1, user2})
	if err != nil {
		log.Fatal(err)
	}
	//here data is JSON string represented as bytes
	fmt.Println(string(data))
}
