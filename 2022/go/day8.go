package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const FilePath = "2022/inputs/day8"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

type trees map[string]bool

func (t *trees) addTree(i, j int) {
	(*t)[fmt.Sprintf("(%d,%d)", i, j)] = true
}

func horizontalVisibleTrees(forest []string, visibleTrees trees) trees {
	for i, line := range forest {
		currentMax := line[0]
		for j := 1; j < len(line); j++ {
			if line[j] > currentMax {
				currentMax = line[j]
				visibleTrees.addTree(i, j)
			}
		}
		currentMax = line[len(line)-1]
		for j := len(line) - 2; j >= 0; j-- {
			if line[j] > currentMax {
				currentMax = line[j]
				visibleTrees.addTree(i, j)
			}
		}
		visibleTrees.addTree(i, 0)
		visibleTrees.addTree(i, len(line)-1)
	}
	return visibleTrees
}

func verticalVisibleTrees(forest []string, visibleTrees trees) trees {
	for j := 0; j < len(forest[0]); j++ {
		currentMax := forest[0][j]
		for i := 1; i < len(forest); i++ {
			if forest[i][j] > currentMax {
				currentMax = forest[i][j]
				visibleTrees.addTree(i, j)
			}
		}
		currentMax = forest[len(forest)-1][j]
		for i := len(forest) - 2; i >= 0; i-- {
			if forest[i][j] > currentMax {
				currentMax = forest[i][j]
				visibleTrees.addTree(i, j)
			}
		}
		visibleTrees.addTree(0, j)
		visibleTrees.addTree(len(forest)-1, j)
	}
	return visibleTrees
}



func main() {
	data, err := os.ReadFile(FilePath)
	checkError(err)
	startTime := time.Now()

	stringData := strings.TrimRight(string(data), "\n")
	forest := strings.Split(stringData, "\n")

	visibleTrees := trees{}

	visibleTrees = horizontalVisibleTrees(forest, visibleTrees)
	visibleTrees = verticalVisibleTrees(forest, visibleTrees)

	fmt.Printf("%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Number of visible trees in the forest is:", len(visibleTrees))
	// fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	// fmt.Printf("Tree with highest scenic score has scenic score of: %d\n", part2answer)

	fmt.Printf("\nTotal time elapsed: %v\n", time.Since(startTime))
}
