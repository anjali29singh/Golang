package main

import (
	"fmt"
)

// bidirectional channel is converted to unidirectional channel
// jobs is reciever channel and results channel is used for sending data
func worker(jobs <-chan int, results chan<- int) {

	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {

	if n <= 1 {

		return 1
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	/*c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)

		}
	}()
	go func() {
		for {
			c2 <- "Every two seconds "
			time.Sleep(time.Second * 2)

		}
	}()

	// for {
	// 	fmt.Println(<-c1)
	// 	//channel one waits until channel 2 recieve message
	// 	fmt.Println(<-c2)
	// }

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)

		case msg2 := <-c2:
			fmt.Println(msg2)

		}
	}
	*/

	jobs := make(chan int, 20)
	results := make(chan int, 20)

	go worker(jobs, results)
	for i := 0; i < 20; i++ {
		jobs <- i
	}
	close(jobs)
	for j := 0; j < 20; j++ {
		fmt.Println(<-results)
	}
}
