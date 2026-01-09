package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sem := make(chan struct{}, 2)
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)

		go func(id int) {

			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			fmt.Println("Processing", id)
			time.Sleep(time.Second)
		}(i)
	}

	wg.Wait()
	fmt.Println("All done")
}
