package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	for i := 0; i < 100; i++ {

		go Player1(ch)
		time.Sleep(time.Second)
		go Player2(ch)
		time.Sleep(time.Second)
	}

}

func Player1(ch chan bool) {
	ch <- true
	fmt.Println("Player 1 Played")
}

func Player2(ch chan bool) {
	<-ch
	fmt.Println("Player 2 Played")
}
