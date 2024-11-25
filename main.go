package main

import "fmt"

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

func main(){
	// 1
	printNumbers()
	printLetters()
}