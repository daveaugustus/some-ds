package main

import "fmt"

func main() {
	list := []int{4, 6, 7, 3, 2, 8, 1, 5, 9, 2, 3, 4}

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if list[i] < list[j] {
				a := list[i]
				list[i] = list[j]
				list[j] = a
			}
		}
	}

	fmt.Println(list)
}
