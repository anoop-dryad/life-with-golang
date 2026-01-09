package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	// Producer: generates numbers
	go func() {
		for i := range 10 {
			fmt.Println("Producing: ", i)
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
		close(ch)
	}()

	// Consumer: receives numbers
	go func() {
		for val := range ch {
			fmt.Println("Consumed: ", val)
			time.Sleep(time.Second)
		}
		fmt.Println("Consumer done")
	}()

	time.Sleep(15 * time.Second)
}
