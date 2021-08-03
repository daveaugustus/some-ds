package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Declare the variable
	var a uint32

	// Read the variable from STDIN
	fmt.Scanf("%v", &a)

	var numOfLetters uint32
	fmt.Scanf("%v", &numOfLetters)

	var letters string

	fmt.Scanf("%s\n", &letters)
	if len(letters) != int(numOfLetters) {
		return
	}

	re := regexp.MustCompile("[0-9]+")
	// Output the variable to STDOUT
	nums := re.FindAllString(letters, -1)
	fmt.Println(len(nums))
}
