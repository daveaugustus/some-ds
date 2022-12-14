// Coding implementation:splicing the numbers in the
// slice below, and ensure that the number converted to is
// the smallest number

// var arr = []int{3, 9, 87, 18, 69, 65, 74}
// EX:
// exArr := []int{16, 15, 27, 9}
// expectStrResult := "15" + "16" + "27" + "9"
// smallest number
// expectNumResult := 1516279

package main

import "fmt"

func main() {
	arr := []int{3, 9, 87, 18, 69, 65, 74}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if AddAllDigits(arr[j]) > AddAllDigits(arr[i]) {
				a := arr[i]
				arr[i] = arr[j]
				arr[j] = a
			}
		}
	}
	fmt.Println(arr)
}

func AddAllDigits(dgts int) int {
	sum := 0

	for dgts != 0 {
		sum = sum + dgts%10
		dgts = dgts / 10
	}

	return sum
}
