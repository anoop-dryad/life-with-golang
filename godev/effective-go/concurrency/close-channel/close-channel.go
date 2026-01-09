package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := range 10 {
			ch <- i
		}

		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
