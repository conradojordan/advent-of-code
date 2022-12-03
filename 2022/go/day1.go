package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FilePath = "2022/inputs/day1"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func countCalories(elf string) int {
	var elfFoods []string
	elfFoods = strings.Split(elf, "\n")

	var caloriesCount int
	for _, food := range elfFoods {
		if food == "" {
			continue
		}
		calories, err := strconv.Atoi(food)
		checkError(err)
		caloriesCount += calories
	}
	return caloriesCount
}

func insertIntoTop3(top3 [3]int, newValue int) [3]int {
	if newValue > top3[2] {
		// Insert in 1st position
		top3[0], top3[1], top3[2] = top3[1], top3[2], newValue
	} else if newValue > top3[1] {
		// Insert in 2nd position
		top3[0], top3[1] = top3[1], newValue
	} else {
		// Insert in 3rd position
		top3[0] = newValue
	}
	return top3
}

func main() {
	data, err := os.ReadFile(FilePath)
	checkError(err)

	stringData := string(data)
	splittedData := strings.Split(stringData, "\n\n")

	var top3calories [3]int
	var currentCalories int
	for _, elf := range splittedData {
		currentCalories = countCalories(elf)
		if currentCalories > top3calories[0] {
			top3calories = insertIntoTop3(top3calories, currentCalories)
		}
	}
	fmt.Printf("Top 3 most calories: %v\n", top3calories)
	fmt.Printf("%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Printf("Elf with the highest amount of calories has: %v calories\n", top3calories[2])
	fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Printf("Sum of the top 3 calories: %d\n", top3calories[0]+top3calories[1]+top3calories[2])

}
