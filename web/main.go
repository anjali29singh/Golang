package main

import (
	"fmt"
	"net/http"
)

type person struct {
	name      string
	emailId   string
	scholarNo int
	percents  float32
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	//structs
	person_a := person{name: "anjali", emailId: "anjalisingh@gmail.com", scholarNo: 211114041, percents: 87.7}
	fmt.Fprintf(w, "Whoa,Go is neat! %v", person_a)

}
func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about me!")
}
func main() {
	/*http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000", nil)
	*/
	//value receiver
}
