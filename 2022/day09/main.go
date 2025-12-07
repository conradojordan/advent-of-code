package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const FilePath = "2022/day09/input.txt"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

type Coordinates struct {
	x, y int
}

func calculateNewCoord(currentCoord, pullingKnotCord Coordinates) Coordinates {
	var newCoords Coordinates = currentCoord
	dx, dy := pullingKnotCord.x-currentCoord.x, pullingKnotCord.y-currentCoord.y

	if math.Abs(float64(dx)) == 2 {
		newCoords.x += dx / 2
		if math.Abs(float64(dy)) == 1 {
			newCoords.y += dy
		}
	}
	if math.Abs(float64(dy)) == 2 {
		newCoords.y += dy / 2
		if math.Abs(float64(dx)) == 1 {
			newCoords.x += dx
		}
	}
	return newCoords
}

type Rope struct {
	coords [10]Coordinates

	secondKnotPositions map[string]bool
	tailPositions       map[string]bool
}

func (r *Rope) updateCoords() {
	for i := 1; i < len(r.coords); i++ {
		r.coords[i] = calculateNewCoord(r.coords[i], r.coords[i-1])
	}
}

func (r *Rope) R(times int) {
	for i := 0; i < times; i++ {
		r.coords[0].x++
		r.updateCoords()
		r.savePositions()
	}
}

func (r *Rope) L(times int) {
	for i := 0; i < times; i++ {
		r.coords[0].x--
		r.updateCoords()
		r.savePositions()
	}
}

func (r *Rope) U(times int) {
	for i := 0; i < times; i++ {
		r.coords[0].y++
		r.updateCoords()
		r.savePositions()
	}
}

func (r *Rope) D(times int) {
	for i := 0; i < times; i++ {
		r.coords[0].y--
		r.updateCoords()
		r.savePositions()
	}
}

func (r *Rope) savePositions() {
	r.tailPositions[fmt.Sprintf("(%d,%d)", r.coords[9].x, r.coords[9].y)] = true
	r.secondKnotPositions[fmt.Sprintf("(%d,%d)", r.coords[1].x, r.coords[1].y)] = true
}

func main() {
	data, err := os.ReadFile(FilePath)
	checkError(err)
	startTime := time.Now()

	stringData := strings.TrimRight(string(data), "\n")
	ropeMovements := strings.Split(stringData, "\n")

	var rope Rope = Rope{
		tailPositions:       make(map[string]bool),
		secondKnotPositions: make(map[string]bool),
	}
	rope.savePositions() // Save coordinates (0,0)

	funcMappings := map[string]func(int){
		"R": rope.R,
		"L": rope.L,
		"U": rope.U,
		"D": rope.D,
	}

	var splittedMovement []string
	var direction string
	var times int
	for _, movement := range ropeMovements {
		splittedMovement = strings.Split(movement, " ")
		direction = splittedMovement[0]
		times, err = strconv.Atoi(splittedMovement[1])
		checkError(err)

		funcMappings[direction](times)
	}

	fmt.Printf("%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Answer to part 1 is:", len(rope.secondKnotPositions))
	fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Answer to part 2 is:", len(rope.tailPositions))

	fmt.Printf("\nTotal time elapsed: %v\n", time.Since(startTime))
}
