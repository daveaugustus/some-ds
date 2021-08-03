package main

import (
	"fmt"
	"reflect"
)

func main() {
	// sentence_1 := []string{"amazing", "acting", "abilities"}
	// sentence_2 := []string{"fine", "theatrics", "talent"}
	// similaWords := [][]string{{"amazing", "fine"}, {"fine", "good"}, {"acting", "theatrics"}, {"abilities", "talent"}}

	sentence_1 := []string{"amazing", "acting"}
	sentence_2 := []string{"fine", "theatrics"}
	similaWords := [][]string{{"amazing", "fine"}, {"fine", "good"}, {"acting", "theatrics"}, {"abilities", "talent"}}

	fmt.Println(isSimilar(sentence_1, sentence_2, similaWords))
}

func isSimilar(sentence_1 []string, sentence_2 []string, similarity_matrix [][]string) bool {
	// Write your code here
	if len(sentence_1) == len(sentence_2) {
		innerArr := make([][]string, len(sentence_1))
		for i := range sentence_1 {
			innerArr[i] = []string{sentence_1[i], sentence_2[i]}
		}
		count := 0
		for _, v := range similarity_matrix {
			for _, w := range innerArr {
				if reflect.DeepEqual(w, v) {
					fmt.Println(w, v)
					count++
				}
			}
		}
		if count == len(sentence_1) {
			return true
		}
		return false
	}
	return false

}
