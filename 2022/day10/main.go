package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const FilePath = "2022/day10/input.txt"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

type Instruction struct {
	value, duration int
}

func main() {
	data, err := os.ReadFile(FilePath)
	checkError(err)
	startTime := time.Now()

	stringData := strings.TrimRight(string(data), "\n")
	instructions := strings.Split(stringData, "\n")

	var regX, cycle int = 1, 1
	var totalStrength int
	var ip int
	var currentInstruction *Instruction

	for ip < len(instructions) {
		if regX-1 <= ((cycle-1)%40) && regX+1 >= ((cycle-1)%40) {
			fmt.Print("██")
		} else {
			fmt.Print("░░")
		}
		if cycle%40 == 0 {
			fmt.Print("\n")
		}
		if (cycle-20)%40 == 0 {
			totalStrength += regX * cycle
		}
		if instructions[ip] == "noop" {
			ip++
		} else {
			if currentInstruction == nil {
				instructionArgs := strings.Split(instructions[ip], " ")
				instructionValue, err := strconv.Atoi(instructionArgs[1])
				checkError(err)
				currentInstruction = &Instruction{value: instructionValue, duration: 2}
			}
			currentInstruction.duration--
			if currentInstruction.duration == 0 {
				regX += currentInstruction.value
				currentInstruction = nil
				ip++
			}
		}
		cycle++
	}

	fmt.Printf("\n%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Answer to part 1 is:", totalStrength)
	fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Answer to part 2 is: EGLHBLFJ")

	fmt.Printf("\nTotal time elapsed: %v\n", time.Since(startTime))
}
