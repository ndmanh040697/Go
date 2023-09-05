package main

import (
	"fmt"
)

func main() {
	var firstName string
	var lastName string
	fmt.Print("First Name:")
	fmt.Scan(&firstName)
	fmt.Print("Last Name:")
	fmt.Scan(&lastName)
	fmt.Println(getFullName(firstName, lastName))
}

func getFullName(firstName, lastName string) string {
	return "Hello I am" + " " + firstName + " " + lastName

}

