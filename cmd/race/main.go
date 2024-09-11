package main

import (
	"fmt"
	"sync"
)

func main() {
	numbersChannel := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go numberPrinter(numbersChannel, &wg, 0)
	go numberPrinter(numbersChannel, &wg, 1)

	go func() {
		wg.Wait()
		close(numbersChannel)
	}()

	for data := range numbersChannel {
		fmt.Printf("Received: %d\n", data)
	}
}

func numberPrinter(data chan int, wg *sync.WaitGroup, startNumber int) {
	defer wg.Done()
	for i := startNumber; i < 100; i += 2 {
		data <- i
	}
}
