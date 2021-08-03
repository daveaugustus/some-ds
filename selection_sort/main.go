// arr[] = 64 25 12 22 11

// // Find the minimum element in arr[0...4]
// // and place it at beginning
// 11 25 12 22 64

// // Find the minimum element in arr[1...4]
// // and place it at beginning of arr[1...4]
// 11 12 25 22 64
package main

import (
	"fmt"
)

func main() {
	list := []int{13, 25, 12, 22, 64}

	temp := 0
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if list[j] > list[i] {
				temp = list[i]
				list[i] = list[j]
				list[j] = temp
			}
		}
	}

	// fmt.Println(list1)
	fmt.Println(list)
}
