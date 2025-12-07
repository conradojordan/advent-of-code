package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/conradojordan/advent-of-code/utils"
)

const FilePath = "./2025/day02/example.txt"

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
		rot_amount, err := strconv.Atoi(rot[1:])
		checkError(err)

		// Count and reduce full revolutions
		count += (rot_amount / 100)
		rot_amount = rot_amount % 100

		if rot[0] == 'L' {
			if rot_amount >= dial && dial != 0 {
				count++
			}
			dial -= rot_amount
			if dial < 0 {
				dial += 100
			}
		} else {
			dial += rot_amount
			if dial >= 100 {
				dial -= 100
				count++
			}
		}
	}
	return count
}

func main() {
	data, err := utils.ParseInput(FilePath, ",")
	checkError(err)

	var rangesLens = make([][2]int, 0, 100)

	for _, val := range data {
		range_str := strings.Split(val, "-")
		range_low := len(range_str[0])
		range_high := len(range_str[1])
		curr_range := [2]int{range_low, range_high}
		rangesLens = append(rangesLens, curr_range)
	}

	fmt.Printf("%#v\n", rangesLens)
	// fmt.Println("---------- Part 1 ----------")
	// fmt.Println("Number of times dial pointed at 0:", part1(data))
	// fmt.Println("---------- Part 2 ----------")
	// fmt.Println("Number of times dial passed through 0:", part2(data))

}
