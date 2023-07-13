package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("./../functions/functions.go")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(f)
	}
}
