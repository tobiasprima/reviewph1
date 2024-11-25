package main

import (
	"fmt"
	"sync"
)

// 1. Functions to print numbers and letters concurrently
func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func printLetters() {
	for ch := 'a'; ch <= 'j'; ch++ {
		fmt.Println(string(ch))
	}
}

// 3. Produce and consume functions
func produce(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consume(ch chan int) {
	for num := range ch {
		fmt.Printf("Consumed: %d\n", num)
	}
}


// 2. Use sync.WaitGroup to wait for goroutines to complete
func main() {
	var wg sync.WaitGroup

	// Task 1.3 & 2.1, 2.2
	wg.Add(2)
	go func() {
		defer wg.Done()
		printNumbers()
	}()
	go func() {
		defer wg.Done()
		printLetters()
	}()
	wg.Wait()

	// Task 3. Produce and consume using a channel
	numbers := make(chan int)
	go produce(numbers)
	go consume(numbers)


	// Task 4. Buffered channel
	bufferedChannel := make(chan int, 5)
	go produce(bufferedChannel)
	go consume(bufferedChannel)
}
