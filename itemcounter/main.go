package main

import "fmt"

func main() {
	arr := []int{1, 2, 1, 3, 5, 2, 1, 1, 9, 9, 9}

	fmt.Println(ItemCount(arr))
}

func ItemCount(arr []int) map[int]int {
	countTracker := make(map[int]int)

	for i := range arr {
		counter := 0
		for j := range arr {
			if arr[j] == arr[i] {
				counter++
			}
		}
		countTracker[arr[i]] = counter
	}

	return countTracker
}
