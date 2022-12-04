package main

import (
	"fmt"
	"os"
	"strings"
)

const FilePath = "2022/inputs/day3"

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

	stringData := string(data)
	splittedData := strings.Split(stringData, "\n")
	pointsMap := generatePointsMap()

	var totalPoints int
	var commonElement string

	for _, rucksack := range splittedData {
		commonElement = getCommonElement(splitInHalf(rucksack))
		totalPoints += pointsMap[commonElement]
	}

	fmt.Printf("%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Total score:", totalPoints)
}
