package main

import "fmt"

func main() {
	ch := make(chan bool)

	for {

		go Player1(ch)
		go Player2(ch)
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
