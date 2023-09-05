package main

import (
	"fmt"
	"math"
	"time"
)

func myFunction(a float64, b float64) (result_1, result_2 float64) {
	result_1 = math.Sqrt(a)
	result_2 = math.Sqrt(b)
	return

}

func main() {
	firstName := "Manh"
	lastName := "Nguyen"
	fullName := firstName + "+" + lastName
	i := 0

	msg := "hello world 123"
	fmt.Println("hello world")
	fmt.Println(msg)
	fmt.Println(firstName, "+", lastName)
	fmt.Println(fullName)
	fullName = firstName + "-" + firstName
	fmt.Println(fullName)
	fmt.Print(time.Now())

	for i := 0; i < 10; i++ {
		fmt.Println(i+1, fullName)
	}

	for {
		fmt.Println(i+1, fullName)
		i = i + 1
		if i >= 5 {
			break

		}

	}
	fmt.Println(myFunction(100, 25))

}
