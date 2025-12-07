package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/conradojordan/advent-of-code/utils"
)

const FilePath = "./2025/day01/input.txt"

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func part1(data []string) int {
	count := 0
	dial := 50
	for _, rot := range data {
		rot_amount, err := strconv.Atoi(rot[1:])
		checkError(err)

		if rot[0] == 'L' {
			dial -= rot_amount
		} else {
			dial += rot_amount
		}

		dial = dial % 100

		if dial == 0 {
			count++
		}
	}
	return count
}

// func part2(data []string) {
// }

func main() {
	data, err := utils.ParseInput(FilePath)
	checkError(err)
	fmt.Println("---------- Part 1 ----------")
	fmt.Println("Number of times dial pointed at 0:", part1(data))
	fmt.Println("---------- Part 2 ----------")
}
