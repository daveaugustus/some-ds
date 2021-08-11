// wqssstttssstraaabbaaaymzuyt
// Find all the sub

package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "wqssstttssstraaabbaaaymzuyt"
	// input[0:]
	SubsTringPal(input)
	// println(CheckPelindrom("aaaba"))
}

func SubsTringPal(s string) {
	// str := ""
	for i := 0; i < len(s); i++ {
		for j := i + 2; j < len(s); j++ {
			// str = string(s[i]) + string(s[j])
			if CheckPelindrom(s[i:j]) {
				fmt.Println(s[i:j])
			}
			// fmt.Println(s[i:j])
		}
	}
}

func CheckPelindrom(s string) bool {
	s = strings.ToLower(s)
	letterCount := 0
	rCount := len(s) - 1
	for i := range s {
		if string(s[i]) == string(s[rCount]) {
			letterCount++
		}
		rCount--
	}

	if len(s) == letterCount {
		return true
	}
	return false
}
