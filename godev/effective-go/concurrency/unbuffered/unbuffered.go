package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from go routine"

	}()

	msg := <-ch
	fmt.Println(msg)

}
