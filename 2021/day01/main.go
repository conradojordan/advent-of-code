package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const FilePath = "2021/day01/input.txt"

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	data, err := os.ReadFile(FilePath)
	checkError(err)

	splitted_data := strings.Split(strings.TrimSpace(string(data)), "\n")
	last_measurement := math.Inf(1)
	var increased_count int

	for _, v := range splitted_data {
		measurement, err := strconv.ParseFloat(v, 64)
		checkError(err)

		if measurement > last_measurement {
			increased_count++
		}
		last_measurement = measurement
	}
	fmt.Printf("Number of times the measurement increased is %v\n", increased_count)
}
