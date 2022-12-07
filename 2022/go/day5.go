package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const FilePath = "2022/inputs/day5"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

type crateStack []string

func (cs *crateStack) push(element string) {
	*cs = append(*cs, element)
}

func (cs *crateStack) pop() string {
	element := (*cs)[len(*cs)-1]
	*cs = (*cs)[:len(*cs)-1]
	return element
}

type elfCrates [9]crateStack

func (ec *elfCrates) makeMove(move string) {
	numCrates, source, destiny := parseMove(move)
	var crate string

	for i := 0; i < numCrates; i++ {
		crate = ec[source-1].pop()
		ec[destiny-1].push(crate)
	}
}

func (ec elfCrates) printCrates() {
	var maxSize int
	var currentSize int

	for _, v := range ec {
		currentSize = len(v)
		if currentSize > maxSize {
			maxSize = currentSize
		}
	}

	for i := maxSize; i >= 0; i-- {
		for j := 0; j < len(ec); j++ {
			if i < len(ec[j]) {
				fmt.Printf("%v ", ec[j][i])
			} else {
				fmt.Print("    ")
			}
		}
		fmt.Println("")
	}
	fmt.Println(" 1   2   3   4   5   6   7   8   9 ")
}

func (ec elfCrates) answer() string {
	var answer string
	for _, crateStack := range ec {
		answer = answer + string(crateStack[len(crateStack)-1][1])
	}
	return answer
}

func parseCrates(crates string) elfCrates {
	var crateStacks elfCrates
	lines := strings.Split(crates, "\n")

	var index uint8
	var element string
	for i := len(lines) - 2; i >= 0; i-- {
		for j := 0; j < len(crateStacks); j++ {
			index = uint8(j) * 4
			element = lines[i][index : index+3]
			if element != "   " {
				crateStacks[j].push(lines[i][index : index+3])

			}
		}
	}
	return crateStacks
}

func parseMove(move string) (int, int, int) {
	splittedMove := strings.Split(move, " ")
	numCrates, err := strconv.Atoi(splittedMove[1])
	checkError(err)

	source, err := strconv.Atoi(splittedMove[3])
	checkError(err)

	destiny, err := strconv.Atoi(splittedMove[5])
	checkError(err)

	return numCrates, source, destiny
}

func main() {
	data, err := os.ReadFile(FilePath)
	checkError(err)
	startTime := time.Now()
	stringData := strings.TrimRight(string(data), "\n")
	cratesAndMoves := strings.Split(stringData, "\n\n")

	crates := parseCrates(cratesAndMoves[0])
	moves := strings.Split(cratesAndMoves[1], "\n")

	fmt.Println("Initial crates arrangement:")
	crates.printCrates()

	for _, move := range moves {
		crates.makeMove(move)
	}

	fmt.Println("\nFinal crates arrangement:")
	crates.printCrates()

	fmt.Printf("%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Arrangement of crates after moves:", crates.answer())
	fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	// fmt.Println("Total number of overlaps:", overlapCount)
	fmt.Printf("\nTotal time elapsed: %v\n", time.Since(startTime))
}
