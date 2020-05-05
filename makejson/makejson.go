package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var name string
	var address string
	var nameAddress = make(map[string]string)

	fmt.Println("Enter a name")
	fmt.Scanln(&name)

	fmt.Println("Enter a address")
	fmt.Scanln(&address)

	nameAddress[name] = address
	jsonObj, _ := json.Marshal(nameAddress)

	fmt.Println(string(jsonObj))

}
