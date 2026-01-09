package main

import "fmt"

func worker(done chan bool) {
	fmt.Println("Worker....")
	done <- true
}

func main() {
	ch := make(chan bool)

	go worker(ch)

	<-ch
	fmt.Println("Finished")
}
