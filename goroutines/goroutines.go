package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("step2: ", url)
	defer response.Body.Close()
	fmt.Println("step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("step4: ", len(body))
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)

}

// channels are the way for goroutines to communicate with each other

func main() {
	//when main goroutine finishes the program exits regardless of any goroutine running in the background.

	/*go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	go responseSize("https://stackoverflow.com")

	*/
	// time.Sleep(10 * time.Second)
	// fmt.Scanln() //prevent main goroutine from exit before finishing the goroutines
	/*var wg sync.WaitGroup
	wg.Add(1)
	//
	go func() {
		count("sheep")
		wg.Done()
		//
	}()
	wg.Wait()
	*/

	c := make(chan string)
	go count("sheep", c)

	// for {

	// 	msg, open := <-c
	// 	if !open {
	// 		break

	// 	}
	// 	fmt.Println(msg)
	// }
	//same thing can also be done like this

	for msg := range c {
		fmt.Println(msg)
	}
}
