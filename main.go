package main

import (
	"errors"
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

func distributeNumbers(even, odd chan int, errChan chan error) {
	defer close(even)
	defer close(odd)
	defer close(errChan)

	for i := 1; i <= 25; i++ { // intentionally going beyond 20 for error demonstration
		if i > 20 {
			errChan <- errors.New(fmt.Sprintf("Invalid number: %d", i))
		} else if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
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

	// Task 5. Separate even and odd channels
	even := make(chan int)
	odd := make(chan int)
	errorChannel := make(chan error)
	go distributeNumbers(even, odd, errorChannel)

	for {
		select {
		case num, ok := <-even:
			if !ok {
				even = nil
			} else {
				fmt.Printf("Received even: %d\n", num)
			}
		case num, ok := <-odd:
			if !ok {
				odd = nil
			} else {
				fmt.Printf("Received odd: %d\n", num)
			}
		case err, ok := <-errorChannel:
			if ok {
				fmt.Printf("Error received: %v\n", err)
			} else {
				errorChannel = nil
			}
		}
		if even == nil && odd == nil && errorChannel == nil {
			break
		}
	}
}
