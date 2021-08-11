// Golang program to illustrate the
// concept of inheritance
package main

import (
	"fmt"
)

// declaring a struct
type Comic struct {
	Universe string
}

func (comic *Comic) ComicUniverse() string {
	return comic.Universe
}

type Marvel struct {
	Comic
}

type DC struct {
	Comic
}

func main() {
	c1 := Marvel{

		Comic{
			Universe: "MCU",
		},
	}

	c2 := DC{
		Comic{
			Universe: "DC",
		},
	}

	fmt.Println("Universe is:", c2.ComicUniverse())
	fmt.Println("Universe is:", c1.ComicUniverse())
}
