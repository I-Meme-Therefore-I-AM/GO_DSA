package main

import "fmt"

// a simple hello world program to open up my 57 exercersice
// for a programmer

func main() {
	fmt.Printf("What is your name ?: ")
	var input string
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Please enter a valid name")
		return
	}

	fmt.Printf("Hello, %s nice to meet you!\n", input)

}
