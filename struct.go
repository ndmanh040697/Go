// getFullName
package main

import "fmt"

type student struct {
	FirstName string
	LastName  string
	Age       int
	Grade     float32
	ClassName string
}

func main() {
	manh := student{
		FirstName: "Manh",
		LastName:  "Nguyen",
		Age:       26,
		Grade:     9,
		ClassName: "HHM",
	}
	fmt.Println(getFullName(manh.FirstName, manh.LastName, manh))
	fmt.Println(setName("Nhi", "Nguyen", manh))

}

func getFullName(a, b string, c student) string {
	return c.FirstName + c.LastName

}
func setName(a, b string, c student) student {
	c.FirstName = a
	c.LastName = b
	return c
}
