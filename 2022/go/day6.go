package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const FilePath = "2022/inputs/day6"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile(FilePath)
	checkError(err)
	startTime := time.Now()
	datastream := strings.TrimRight(string(data), "\n")

	var answer int
	var c01, c02, c03, c12, c13, c23 bool
	c12 = datastream[0] != datastream[1]
	c13 = datastream[0] != datastream[2]
	c23 = datastream[1] != datastream[2]

	for i := 3; i < len(datastream); i++ {
		c01, c02, c12 = c12, c13, c23
		c03 = datastream[i-3] != datastream[i]
		c13 = datastream[i-2] != datastream[i]
		c23 = datastream[i-1] != datastream[i]
		if c01 && c02 && c03 && c12 && c13 && c23 {
			answer = i + 1
			break
		}
	}

	fmt.Printf("%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Index of end of first 4 different characters is:", answer)
	// fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	// fmt.Println("Arrangement of crates after moves:", cratesPart2.answer())
	fmt.Printf("\nTotal time elapsed: %v\n", time.Since(startTime))
}
