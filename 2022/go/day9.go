package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const FilePath = "2022/inputs/day9"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

type Rope struct {
	x, y        int
	horiz, vert int8

	tailPositions map[string]bool
}

func (r *Rope) R(times int) {
	for i := 0; i < times; i++ {
		if r.horiz == 1 {
			r.x++
			if r.vert != 0 {
				r.y += int(r.vert)
				r.vert = 0
			}
			r.addTailPosition(r.x, r.y)
		} else {
			r.horiz++
		}
	}
}

func (r *Rope) L(times int) {
	for i := 0; i < times; i++ {
		if r.horiz == -1 {
			r.x--
			if r.vert != 0 {
				r.y += int(r.vert)
				r.vert = 0
			}
			r.addTailPosition(r.x, r.y)
		} else {
			r.horiz--
		}
	}
}

func (r *Rope) U(times int) {
	for i := 0; i < times; i++ {
		if r.vert == 1 {
			r.y++
			if r.horiz != 0 {
				r.x += int(r.horiz)
				r.horiz = 0
			}
			r.addTailPosition(r.x, r.y)
		} else {
			r.vert++
		}
	}
}

func (r *Rope) D(times int) {
	for i := 0; i < times; i++ {
		if r.vert == -1 {
			r.y--
			if r.horiz != 0 {
				r.x += int(r.horiz)
				r.horiz = 0
			}
			r.addTailPosition(r.x, r.y)
		} else {
			r.vert--
		}
	}
}

func (r *Rope) addTailPosition(x, y int) {
	r.tailPositions[fmt.Sprintf("(%d,%d)", x, y)] = true
}

func main() {
	data, err := os.ReadFile(FilePath)
	checkError(err)
	startTime := time.Now()

	stringData := strings.TrimRight(string(data), "\n")
	ropeMovements := strings.Split(stringData, "\n")

	var rope Rope = Rope{
		tailPositions: make(map[string]bool),
	}
	rope.addTailPosition(0, 0)

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
	fmt.Println("Answer to part 1 is:", len(rope.tailPositions))
	fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	// fmt.Println("Answer to part 2 is:", part2answer)

	fmt.Printf("\nTotal time elapsed: %v\n", time.Since(startTime))
}
