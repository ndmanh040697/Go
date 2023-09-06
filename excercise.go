// sort slice alphabet
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var name = []string{}
/* cach 1: input lenght of Slice */
	// fmt.Print("Enter len: ")
	// var n int
	// fmt.Scanf("%d", &n)
	// var input string
	// for i := 0; i < n; {
	// 	fmt.Printf("name[%d]", i)
	// 	fmt.Scan(&input)
	// 	name = append(name, input)
	// 	i += 1
	// }
/* Cach 2: input into Slice */
	for {
		var input string
		fmt.Print("inputName=")
		fmt.Scan(&input)
		name = append(name, input)

		if input == " " {
			break
		}
	}
	fmt.Println("UNSORTED")
	printStrings(name)
	sort.Sort(Alphabetic(name))
	fmt.Println()
	fmt.Println("SORTED ALPHABETICALLY")
	printStrings(name)
}

type Alphabetic []string

func (list Alphabetic) Len() int { return len(list) }

func (list Alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func (list Alphabetic) Less(i, j int) bool {
	var si string = list[i]
	var sj string = list[j]
	var si_lower = strings.ToLower(si)
	var sj_lower = strings.ToLower(sj)
	if si_lower == sj_lower {
		return si < sj
	}
	return si_lower < sj_lower
}

func printStrings(name []string) {
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}
}
