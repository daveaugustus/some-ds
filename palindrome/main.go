package main

import (
	"fmt"
	"strings"
)

func main() {
	// Declare the variable
	var a uint32
	var inputs []string
	var ip string
	// Read the variable from STDIN
	fmt.Scanf("%v", &a)

	for i := 0; i < int(a); i++ {
		fmt.Scanf("%s\n", &ip)
		inputs = append(inputs, ip)
	}

	for i := 0; i < int(a); i++ {
		if CheckPalindrome(inputs[i]) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
	// // Output the variable to STDOUT
	// fmt.Println(a)

	// fmt.Println()
}

func CheckPalindrome(str string) bool {
	str = strings.ToLower(str)
	letterCounter := 0
	rCount := len(str) - 1
	for i := range str {
		if string(str[i]) == string(str[rCount]) {
			letterCounter++
		}
		// fmt.Println(string(str[i]), string(str[rCount]))
		rCount--
	}

	if len(str) == letterCounter {
		return true
	}
	return false
}
