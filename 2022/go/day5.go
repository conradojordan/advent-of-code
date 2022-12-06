package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const FilePath = "2022/inputs/day5"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// type craneStack struct {
// 	[]
// }

func parseCranes(cranes string) [][9]string {
	var cranesColumns [][9]string
	for _, line := range strings.Split(cranes, "\n") {

	}
}

func main() {
	data, err := os.ReadFile(FilePath)
	checkError(err)
	startTime := time.Now()

	stringData := strings.TrimSpace(string(data))
	cranesAndMoves := strings.Split(stringData, "\n\n")
	fmt.Printf("Cranes:\n%v\n\n", cranesAndMoves[0])
	fmt.Printf("Moves:\n%v\n", cranesAndMoves[1])

	endTime := time.Now()
	fmt.Printf("%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	// fmt.Println("Total number of subsets:", subsetCount)
	fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	// fmt.Println("Total number of overlaps:", overlapCount)
	fmt.Printf("\nTotal time elapsed: %v\n", endTime.Sub(startTime))
}
