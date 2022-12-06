package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const FilePath = "2022/inputs/day4"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func splitElfs(elfPair string) ([2]int, [2]int) {
	var elf1, elf2 [2]int
	var err error

	splitted := strings.Split(elfPair, "-")
	elf1[0], err = strconv.Atoi(splitted[0])
	checkError(err)
	elf2[1], err = strconv.Atoi(splitted[2])
	checkError(err)

	splitted = strings.Split(splitted[1], ",")
	elf1[1], err = strconv.Atoi(splitted[0])
	checkError(err)
	elf2[0], err = strconv.Atoi(splitted[1])
	checkError(err)

	return elf1, elf2
}

func hasSubset(elf1, elf2 [2]int) bool {
	return !((elf1[0] < elf2[0] && elf1[1] < elf2[1]) || (elf1[0] > elf2[0] && elf1[1] > elf2[1]))
}

func hasOverlap(elf1, elf2 [2]int) bool {
	return !(elf1[1] < elf2[0] || elf1[0] > elf2[1])
}

func main() {
	data, err := os.ReadFile(FilePath)
	checkError(err)
	startTime := time.Now()

	stringData := strings.TrimSpace(string(data))
	elfPairs := strings.Split(stringData, "\n")

	var subsetCount int
	var overlapCount int
	for _, elfPair := range elfPairs {
		elf1, elf2 := splitElfs(elfPair)
		if hasSubset(elf1, elf2) {
			subsetCount++
		}
		if hasOverlap(elf1, elf2) {
			overlapCount++
		}
	}

	endTime := time.Now()
	fmt.Printf("%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Total number of subsets:", subsetCount)
	fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Total number of overlaps:", overlapCount)
	fmt.Printf("\nTotal time elapsed: %v\n", endTime.Sub(startTime))
}
