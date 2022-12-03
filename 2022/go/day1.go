package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FilePath = "2022/inputs/day1"

func check(e error) {
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
		check(err)
		caloriesCount += calories
	}
	return caloriesCount
}

func main() {
	data, err := os.ReadFile(FilePath)
	check(err)

	stringData := string(data)
	splittedData := strings.Split(stringData, "\n\n")

	var maxCalories, currentCalories int
	for _, elf := range splittedData {
		currentCalories = countCalories(elf)
		if currentCalories > maxCalories {
			maxCalories = currentCalories
		}
	}
	fmt.Printf("Elf with the highest amount of calories has: %v calories\n", maxCalories)

}
