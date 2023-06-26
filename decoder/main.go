package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

//convert json into 
func main() {

	jsonData := `{"name":"anjali", "email":"anjalisinghdev2@gmail.com","age":22}`

	reader := strings.NewReader(jsonData)
	dec := json.NewDecoder(reader)
	var person Person
	err := dec.Decode(&person)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(person)

	}

}
