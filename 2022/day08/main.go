package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const FilePath = "2022/day08/input.txt"

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

func getHighestScenicScore(forest []string) (int, int, int) {
	var maxScenicScore, scenicScore int
	var maxI, maxJ int

	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[i])-1; j++ {
			scenicScore = getScenicScore(i, j, forest)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
				maxI, maxJ = i, j
			}
		}
	}
	return maxScenicScore, maxI, maxJ
}

func getScenicScore(i, j int, forest []string) int {
	var up, down, left, right int

	for k := i - 1; k >= 0; k-- {
		up++
		if forest[k][j] >= forest[i][j] {
			break
		}
	}
	for k := i + 1; k <= len(forest)-1; k++ {
		down++
		if forest[k][j] >= forest[i][j] {
			break
		}
	}
	for k := j - 1; k >= 0; k-- {
		left++
		if forest[i][k] >= forest[i][j] {
			break
		}
	}
	for k := j + 1; k <= len(forest[i])-1; k++ {
		right++
		if forest[i][k] >= forest[i][j] {
			break
		}
	}
	return up * down * left * right
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

	highestScenicScore, i, j := getHighestScenicScore(forest)

	fmt.Printf("%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Number of visible trees in the forest is:", len(visibleTrees))
	fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Printf("Tree with highest scenic score has scenic score of: %d (found at coordinates {%d, %d})\n", highestScenicScore, i, j)

	fmt.Printf("\nTotal time elapsed: %v\n", time.Since(startTime))
}
