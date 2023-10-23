// getFullName
package main

import (
	"app/Go/payload"
	"fmt"
)

func main() {
	// slice := []int{1, 2, 3, 5, 2, 23, 56}
	// slice1 := []int{12, 32, 4, 5, 3, 2, 13, 3}
	// manh := model.Student{
	// 	FirstName: "Manh",
	// 	LastName:  "Nguyen",
	// 	Age:       26,
	// 	Grade:     9,
	// 	ClassName: "HHM",
	// }
	eric := payload.Student{}
	JSONString := `{"FirstName":"Eric","ho":"Nguyen","tuoi":0,"Grade":0,"ClassName":""}`
	eric.FromJson(JSONString)

	v := eric.ToModel()

	s := v.ToJson()
	fmt.Println(s)

}

// // BTVN 3
// func getFullName(a, b string, c Student) string {
// 	return c.FirstName + " " + c.LastName

// }

// // BTVN 4

// // BTVN 5
// func addEndSlice(a []int, b int) []int {
// 	a = append(a, b)
// 	return a

// }

// // BTVN 6
// func addStartSlice(a []int, b int) []int {
// 	a = append([]int{b}, a...)
// 	return a

// }

// // BTVN 7
// func removeElement(a []int, b int) []int {
// 	c := a[:b]
// 	d := a[b+1:]
// 	a = append(c, d...)
// 	return a

// }

// // BTVN 8
// func insertElement(a []int, b int, c int) []int {
// 	d := a
// 	d = append(d[:b], c)
// 	fmt.Println("Slice d:", d)
// 	fmt.Println("Slice a:", a)
// 	// a = append(d, a[b:]...)
// 	return a
// }

// // BTVN 9
// func addSlice(a []int, b []int) []int {
// 	return append(a, b...)
// }
