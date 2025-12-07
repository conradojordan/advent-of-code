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

func part2(data []string) int {
	count := 0
	dial := 50

	for _, rot := range data {
		old_dial := dial

		rot_amount, err := strconv.Atoi(rot[1:])
		checkError(err)

		count += (rot_amount / 100)
		rot_amount = rot_amount % 100

		if rot[0] == 'L' {
			dial -= rot_amount
			if dial < 0 && old_dial != 0 {
				count++
			}
		} else {
			dial += rot_amount
			if dial > 100 && old_dial != 0 {
				count++
			}
		}

		if dial < 0 {
			dial += 100
		}

		if dial >= 100 {
			dial -= 100
		}

		if dial == 0 {
			count++
		}
	}
	return count
}

func main() {
	data, err := utils.ParseInput(FilePath)
	checkError(err)

	fmt.Println("---------- Part 1 ----------")
	fmt.Println("Number of times dial pointed at 0:", part1(data))
	fmt.Println("---------- Part 2 ----------")
	fmt.Println("Number of times dial passed through 0:", part2(data))

}
