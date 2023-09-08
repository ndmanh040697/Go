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
	slice := []int{1, 2, 3, 5, 2, 23, 56}
	fmt.Println()
	manh := student{
		FirstName: "Manh",
		LastName:  "Nguyen",
		Age:       26,
		Grade:     9,
		ClassName: "HHM",
	}
	fmt.Println(getFullName(manh.FirstName, manh.LastName, manh))
	fmt.Println(setName("Nhi", "Nguyen", manh))
	// fmt.Println("Add element to the end of slice", addEndSlice(slice, 8))
	// fmt.Println("Add element to the start of slice", addStartSlice(slice, 8))
	// fmt.Println("Remove element in slice", removeElement(slice, 2))
	fmt.Println("Insert element in slice", insertElement(slice, 3, 84))

}

func getFullName(a, b string, c student) string {
	return c.FirstName + " " + c.LastName

}
func setName(a, b string, c student) student {
	c.FirstName = a
	c.LastName = b
	return c
}

func addEndSlice(a []int, b int) []int {
	a = append(a, b)
	return a

}

func addStartSlice(a []int, b int) []int {
	a = append([]int{b}, a...)
	return a

}

func removeElement(a []int, b int) []int {
	c := a[:b]
	d := a[b+1:]
	a = append(c, d...)
	return a

}

func insertElement(a []int, b int, c int) []int {
	d := a
	d = append(d[:b], c)
	fmt.Println("Slice d:",d)
	fmt.Println("Slice a:",a)
	// a = append(d, a[b:]...)
	return a
}
