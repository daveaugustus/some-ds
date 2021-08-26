package main

import "fmt"

func main() {
	counter := make(chan int, 5)

	for i := 0; i < 10; i++ {
		go func(i int) {
			counter <- i
		}(i)
	}

	for i := 0; i < 10; i++ {
		val := <-counter
		fmt.Println(val)
	}

}
