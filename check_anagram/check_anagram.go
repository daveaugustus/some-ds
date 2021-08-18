package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input1 := "listen"
	input2 := "silenta"

	fmt.Println(CheckAnagram(input1, input2))
}

func CheckAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	str1 = SortStr(str1)
	str2 = SortStr(str2)

	if str1 == str2 {
		return true
	}
	return false
}

func SortStr(s string) string {
	strSlice := strings.Split(s, "")
	sort.Strings(strSlice)

	return strings.Join(strSlice, "")
}
