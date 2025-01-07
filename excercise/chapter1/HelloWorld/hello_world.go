package main

import (
	"bufio"
	"fmt"
	"os"
)

// a simple hello world program to open up my 57 exercersice
// for a programmer

func main() {
	fmt.Printf("What is your name?: ")
	var input string
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Please enter a valid name")
		return
	}
	fmt.Printf("Hello, %s nice to meet you!\n", input)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Press 0")
	var text string

	// just a stupid pro
outerLoop:
	for scanner.Scan() {
		text = scanner.Text()
		switch {
		case text == "0":
			break outerLoop
		default:
			fmt.Println("Press 0")
		}
	}
}
