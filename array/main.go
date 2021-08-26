package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4}
	fmt.Println(ReverseArrayRotatingElements(input))
}

func ReverseArrayRotatingElements(arr []int) []int {
	start := 0
	end := len(arr) - 1

	for start < end {
		temp := arr[end]
		arr[end] = arr[start]
		arr[start] = temp
		start++
		end--
	}

	return arr
}

func foo() {

	for a := 0; a < 10; a++ {

	}
}

// func ReverseArrayRecursiveWay(arr []int) []int {
// 	start := 0
// 	end := len(arr) - 1

// 	if start < end {
// 		return nil
// 	}

// 	temp := arr[end]
// 	arr[end] = arr[start]
// 	arr[start] = temp

// 	return nil
// }
