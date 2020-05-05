package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Name with first and last name
type Name struct {
	fname string
	lname string
}

func main() {
	//Initilize vars
	nameSlice := make([]Name, 0, 55)
	var fileName string

	//Get filename from user
	fmt.Println("Enter file name (e.g. names.txt)")
	fmt.Scanln(&fileName)

	//Read the file and store it in data
	data, error := ioutil.ReadFile(fileName)

	if error != nil {
		fmt.Printf("Error")
	}

	//make a string of full names
	namesArray := strings.Split(string(data), "\n")

	//loop to append data to nameSlice
	for _, fullName := range namesArray {

		//break the full name into two parts array
		singleNameArray := strings.Fields(fullName)

		//Trim if string is longer than 20 chars
		if len(singleNameArray[0]) > 20 {
			singleNameArray[0] = singleNameArray[0][0:20]
		}
		if len(singleNameArray[1]) > 20 {
			singleNameArray[1] = singleNameArray[1][0:20]
		}

		nameSlice = append(nameSlice, Name{fname: singleNameArray[0], lname: singleNameArray[1]})
	}

	//print the slice
	for _, name := range nameSlice {
		fmt.Printf("First Name: %s Last Name: %s\n", name.fname, name.lname)
	}
}
