package main

import "fmt"

func main() {
	arr := []int{4, 6, 7, 3, 2, 8, 1, 5, 9, 2, 3, 4}

	for i := 0; i < len(arr); i++ {
		if i == len(arr) {
			fmt.Println(i)
			break
		}

		// temp := arr[i]
		arr[i] = arr[i+1]

	}

	fmt.Println(arr)
}
