package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat", "mat", "tma"}

	output := AnagramWrapper(input)

	fmt.Printf("%v\n", output)
}

func AnagramWrapper(input []string) [][]string {
	var output [][]string

	for i := 0; i < len(input); i++ {
		var arr []string

		for j := 0; j < len(input); j++ {
			conString := input[i] + input[i]
			if strings.Contains(conString, input[j]) {
				if !SliceContains(arr, input[j]) {
					arr = append(arr, input[j])
				}
			}

		}
		if !SliceContainsSlice(output, arr) {
			output = append(output, arr)
		}
	}

	return output
}

func SliceContains(s []string, val string) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

func SliceContainsSlice(slice [][]string, subSlice []string) bool {
	for _, v := range slice {
		if reflect.DeepEqual(v, subSlice) {
			return true
		}
	}

	return false
}

func SliceContainsSliceOrderLess() {

}
