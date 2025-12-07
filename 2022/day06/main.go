package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const FilePath = "2022/day06/input.txt"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

type byteSet map[byte]int

func (bs *byteSet) add(element byte) {
	(*bs)[element]++
}

func (bs *byteSet) remove(element byte) {
	(*bs)[element]--
	if (*bs)[element] == 0 {
		delete(*bs, element)
	}
}

func firstIndexOfNDifferentCharacters(datastream string, n int) int {
	bs := byteSet{}

	for i := 0; i < n; i++ {
		bs.add(datastream[i])
	}

	if len(bs) == n {
		return n
	}

	for i := n; i < len(datastream); i++ {
		bs.remove(datastream[i-n])
		bs.add(datastream[i])
		if len(bs) == n {
			return i + 1
		}
	}
	return 0

}

func main() {
	data, err := os.ReadFile(FilePath)
	checkError(err)
	startTime := time.Now()
	datastream := strings.TrimRight(string(data), "\n")

	var part1answer, part2answer int

	part1answer = firstIndexOfNDifferentCharacters(datastream, 4)
	part2answer = firstIndexOfNDifferentCharacters(datastream, 14)

	fmt.Printf("%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Index of end of first 4 different characters is:", part1answer)
	fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Index of end of first 4 different characters is:", part2answer)

	fmt.Printf("\nTotal time elapsed: %v\n", time.Since(startTime))
}
