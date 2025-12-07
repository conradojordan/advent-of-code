package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/conradojordan/advent-of-code/utils"
)

const FilePath = "./2025/day02/input.txt"
const Debug = false

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func part1(data []string) int {
	invalidIdSum := 0

	for _, val := range data {
		range_str := strings.Split(val, "-")
		range_low_digits := len(range_str[0])
		range_high_digits := len(range_str[1])

		if (range_low_digits%2 == 1) && (range_high_digits%2 == 1) && (range_low_digits == range_high_digits) {
			if Debug {
				fmt.Println("Skipping impossible range:", range_str[0], range_str[1])
				fmt.Println()
			}
			continue
		}

		original_range_low, err := strconv.Atoi(range_str[0])
		checkError(err)
		original_range_high, err := strconv.Atoi(range_str[1])
		checkError(err)

		var range_low int
		if range_low_digits%2 == 1 {
			range_low = int(math.Pow10(range_low_digits))
			range_low_digits++
		} else {
			range_low = original_range_low
		}

		var range_high int
		if range_high_digits%2 == 1 {
			range_high = int(math.Pow10(range_high_digits-1)) - 1
		} else {
			range_high = original_range_high
		}

		candidate := range_low / int(math.Pow10(range_low_digits/2))
		if Debug {
			fmt.Printf("Range (%v - %v) -> (%v - %v) -> first_candidate: %v, range_low_digits: %v\n", original_range_low, original_range_high, range_low, range_high, candidate, range_low_digits)
		}
		for {
			candidate_digits := int(math.Log10(float64(candidate))) + 1
			invalidId := candidate * (int(math.Pow10(candidate_digits)) + 1)
			candidate++
			if Debug {
				fmt.Println("Found invalidId:", invalidId)
			}
			if invalidId > range_high {
				if Debug {
					fmt.Println("invalidId too high, breaking...")
				}
				break
			}
			if invalidId < range_low {
				if Debug {
					fmt.Println("invalidId too high, continuing...")
				}
				continue
			}
			invalidIdSum += invalidId
		}
		if Debug {
			fmt.Println()
		}
	}
	return invalidIdSum
}

// func part2(data []string) int {
// 	count := 0
// 	dial := 50
//
// 	for _, rot := range data {
// 		rot_amount, err := strconv.Atoi(rot[1:])
// 		checkError(err)
//
// 		// Count and reduce full revolutions
// 		count += (rot_amount / 100)
// 		rot_amount = rot_amount % 100
//
// 		if rot[0] == 'L' {
// 			if rot_amount >= dial && dial != 0 {
// 				count++
// 			}
// 			dial -= rot_amount
// 			if dial < 0 {
// 				dial += 100
// 			}
// 		} else {
// 			dial += rot_amount
// 			if dial >= 100 {
// 				dial -= 100
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

func main() {
	data, err := utils.ParseInput(FilePath, ",")
	checkError(err)

	fmt.Println("---------- Part 1 ----------")
	fmt.Println("Sum of invalid IDs:", part1(data))
	// fmt.Println("---------- Part 2 ----------")
	// fmt.Println("Number of times dial passed through 0:", part2(data))

}
