package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	// Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
	// The program should be written as a loop.
	// Before entering the loop, the program should create an empty integer slice of size (length) 3.
	// During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
	// The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
	// The slice must grow in size to accommodate any number of integers which the user decides to enter.
	// The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

	inputtedNumbers := make([]int, 0, 3)
	fmt.Println("Enter x to stop")

	for {
		// Get user input
		fmt.Println("Enter an integer")
		var userInput string
		fmt.Scanln(&userInput)

		// Break the loop if input is x
		if userInput == "x" {
			break
		}

		// Convert user input string to int
		intUserInput, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Input Error")
		}

		// Append the int value to the declared slice
		inputtedNumbers = append(inputtedNumbers, intUserInput)

		// Sort the slice
		sort.SliceStable(inputtedNumbers, func(i, j int) bool { return inputtedNumbers[i] < inputtedNumbers[j] })

		// Print items of the slice
		fmt.Println(inputtedNumbers)

	}

	// show exit
	fmt.Println("Function End. Final value is as below:")
	// Print items of the slice
	fmt.Println(inputtedNumbers)

}
