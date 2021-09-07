package main

import (
	"fmt"
)

func main() {
	var input []int
	var len int
	var val int
	var targetSum int

	fmt.Println("Enter length of array")
	fmt.Scanf("%d", &len)
	for i := 0; i < len; i++ {
		fmt.Printf("Enter at index %d: \t", i+1)
		fmt.Scanf("%d\t", &val)
		input = append(input, val)
	}

	fmt.Println()
	fmt.Println()
	fmt.Printf("Enter Target Sum Value: \t")
	fmt.Scanf("%d", &targetSum)

	// fmt.Printf("The following set of integers sums %d\n", n)
	ElementsForTargetSum(input, targetSum)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	// Sending Zero at the end
	SendAllZerosAtLast(input)

}

func ElementsForTargetSum(input []int, targetSum int) {
	subsets := AllSubSets(input)

	for i := 0; i < len(subsets); i++ {
		if targetSum == AddArrElements(subsets[i]) {
			fmt.Println(subsets[i])
		}
	}
}

func AllSubSets(input []int) (subsets [][]int) {
	length := len(input)

	for i := 1; i < (1 << length); i++ {
		var subset []int

		for j := 0; j < length; j++ {
			if (i>>j)&1 == 1 {
				// add j to subset
				subset = append(subset, input[j])
			}
		}

		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

func SendAllZerosAtLast(input []int) {
	// Check is there is any 0 in the array
	if !ZeroFinder(input) {
		fmt.Println("The array didn't contain any 0!")
		return
	}

	// Send all 0's at the end
	increment := 0
	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			input[increment] = input[i]
			increment++
		}
	}
	for ; increment != len(input); increment++ {
		input[increment] = 0

	}
	fmt.Println("All 0's pushed at the end:\t", input)
}

func ZeroFinder(input []int) bool {
	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			return true
		}
	}
	return false
}

func AddArrElements(input []int) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	return sum
}

// 0,4,5,3,2,1, -3, 0
