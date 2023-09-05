// Check number
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your the number you want to check:")
	number, _ := reader.ReadString('\n')
	number = strings.Replace(number, "\n", "", -1)
	num, err := strconv.Atoi(number)
	if err == nil {
		fmt.Println("error converting", number)
		return
	}

	fmt.Println("Number is:", num)

	if num < 10 && num > 5 {
		fmt.Println("correct")
	} else {
		fmt.Println("uncorrect")
	}

}
