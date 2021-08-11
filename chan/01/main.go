package main

import (
	"fmt"
	"sync"
	"time"
)

var i int
var wg sync.WaitGroup

func main() {
	// ch := make(chan bool)
	count := 0

	for i := 0; i < 20; i++ {
		fmt.Println(count)
		wg.Add(1)
		go func() {
			count++
		}()
		time.Sleep(time.Second)
		// wg.Add(2)
		// go Player1(ch)
		// time.Sleep(time.Millisecond)
		// go Player2(ch)
		// time.Sleep(time.Millisecond)
	}

	wg.Wait()
}

func Player1(ch chan bool) {
	fmt.Println("Player 1 Played it's turn")
	ch <- true

	wg.Done()
	// fmt.Println("Player 1 Played")
}

func Player2(ch chan bool) {
	fmt.Println("Player 2 Played it's turn")
	<-ch

	wg.Done()
}
