package main

import (
	"fmt"
)

func main() {
	slice := []int{}
	for {
		var input int
		fmt.Print("Input Number=")
		fmt.Scan(&input)
		slice = append(slice, input)

		if input == 100 {
			break
		}
		fmt.Println(bubbleSort(slice))
	}

	// a := []int{1, 4, 7, 2, 3}
	// fmt.Println(bubbleSort(a))
}

func bubbleSort(a []int) []int {
	for i := 0; i < (len(a) - 1); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]

			}
		}

	}
	return a
}
