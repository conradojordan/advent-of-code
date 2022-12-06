package main

import (
	"fmt"
	"os"
	"strings"
)

const FilePath = "2022/inputs/day2"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

var scores map[string]int = map[string]int{
	"A X": 4,
	"A Y": 8,
	"A Z": 3,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 7,
	"C Y": 2,
	"C Z": 6,
}

var scoresPart2 map[string]int = map[string]int{
	// A - Rock 1, B - Paper 2, C - Scissors 3
	// X - Lose 0, Y - Draw 3, Z - Win 6
	"A X": 3,
	"A Y": 4,
	"A Z": 8,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 2,
	"C Y": 6,
	"C Z": 7,
}

func main() {

	data, err := os.ReadFile(FilePath)
	checkError(err)

	stringData := strings.TrimSpace(string(data))
	var plays []string = strings.Split(stringData, "\n")

	var totalScorePart1, totalScorePart2 int

	for _, play := range plays {
		totalScorePart1 += scores[play]
		totalScorePart2 += scoresPart2[play]
	}

	fmt.Printf("%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Total score:", totalScorePart1)
	fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Total score:", totalScorePart2)

}
