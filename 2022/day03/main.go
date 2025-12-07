package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const FilePath = "2022/day03/input.txt"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func getCommonElement(string1, string2 string) string {
	for _, c := range string1 {
		if strings.Contains(string2, string(c)) {
			return string(c)
		}
	}
	return ""
}

func getCommonElementPart2(string1, string2, string3 string) string {
	firstCommonElements := make([]string, 0, 20)
	for _, c := range string1 {
		if strings.Contains(string2, string(c)) {
			firstCommonElements = append(firstCommonElements, string(c))
		}
	}
	for _, c := range firstCommonElements {
		if strings.Contains(string3, string(c)) {
			return string(c)
		}
	}
	return ""
}

func splitInHalf(inputString string) (string, string) {
	return inputString[:len(inputString)/2], inputString[len(inputString)/2:]
}

func generatePointsMap() map[string]int {
	pointsMap := make(map[string]int, 52)
	for i := 97; i < 123; i++ {
		pointsMap[string(i)] = i - 96
	}
	for i := 65; i < 91; i++ {
		pointsMap[string(i)] = i - 38
	}

	return pointsMap
}

func main() {
	data, err := os.ReadFile(FilePath)
	checkError(err)
	startTime := time.Now()

	stringData := strings.TrimRight(string(data), "\n")
	rucksacks := strings.Split(stringData, "\n")
	pointsMap := generatePointsMap()

	var totalPoints int
	var commonElement string

	for _, rucksack := range rucksacks {
		commonElement = getCommonElement(splitInHalf(rucksack))
		totalPoints += pointsMap[commonElement]
	}

	var totalPointsPart2 int
	for i := 0; i < (len(rucksacks) / 3); i++ {
		commonElement = getCommonElementPart2(rucksacks[i*3], rucksacks[i*3+1], rucksacks[i*3+2])
		totalPointsPart2 += pointsMap[commonElement]
	}

	fmt.Printf("%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Total score:", totalPoints)
	fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Total score:", totalPointsPart2)
	fmt.Printf("\nTotal time elapsed: %v\n", time.Since(startTime))
}
