package main

import (
	"fmt"
)

func main() {
	// slice := []int{}
	// for {
	// 	var input int
	// 	fmt.Print("Input Number=")
	// 	fmt.Scan(&input)
	// 	slice = append(slice, input)

	// 	if input == 100 {

	// 		var confirm string
	// 		fmt.Println("Would you want to stop ? y/n:")
	// 		fmt.Scan(&confirm)

	// 		if confirm == "y" {
	// 			break
	// 		} else if confirm == "n" {
	// 			continue
	// 		}

	// 	}
	// }
	// fmt.Println(bubbleSort(slice))

	a := []int{1, 4, 7, 2, 3, 8, 5, 9, 6}
	fmt.Println("Bubble Sort:", bubbleSort(a))
	// fmt.Println("Merge Sort:", mergeSort(a))
}

// BTVN 2
func bubbleSort(a []int) []int {
	for i := 0; i < (len(a) - 2); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]

			}
		}

	}
	return a
}

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	mid := len(a) / 2
	left := a[:mid]
	right := a[mid:]
	return merge(left, right)
}
func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j := 0, 0
	for k := 0; k < len(result); k++ {
		if i >= len(left) {
			result[k] = right[j]
			j++
		} else if j >= len(right) {
			result[k] = left[i]
			i++
		} else if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}
	return result
}
