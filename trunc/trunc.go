package main

import (
	"fmt"
	"strconv"
)

func main() {
	var userInput float64 //Declare variable for user input
	var intValue int      // integer version fo user input

	fmt.Printf("Please, enter a float number: \n")
	fmt.Scan(&userInput)
	intValue = int(userInput)
	fmt.Printf(strconv.Itoa(intValue))
}
