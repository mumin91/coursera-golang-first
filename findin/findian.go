package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var userInput string

	const prefix = "i"
	const middle = "a"
	const suffix = "n"

	fmt.Printf("Enter a string: \n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() == nil {
		userInput = scanner.Text()
	}

	var finalText = strings.ToLower(userInput)

	if strings.HasPrefix(finalText, prefix) && strings.HasSuffix(finalText, suffix) && strings.Contains(finalText, middle) {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not found!")
	}
}
