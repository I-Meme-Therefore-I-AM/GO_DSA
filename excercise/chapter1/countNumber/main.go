package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var input string

	// fmt.Printf("What is the input string:? ")
	// fmt.Scan(&input)

	// fmt.Printf("%s owner has %d characters", input, len(input))

	// if user input nothing
	fmt.Printf("What is the input string:? ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Printf("What is the input string:? ")
		text := scanner.Text()

		if text == "" {
			fmt.Println("Enter a value empty string not rewarded")
			fmt.Printf("What is the input string:? \n")
			continue
		}
		fmt.Printf("%s has %d characters", text, len(text))
		break
	}
}
