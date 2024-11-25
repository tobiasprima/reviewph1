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

}
