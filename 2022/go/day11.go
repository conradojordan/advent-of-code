package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const FilePath = "2022/inputs/day11"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

type Monkey struct {
	id                    int
	items                 []int
	newWorryLevelFunction func(int) int
	divisibleBy           int
	ifTrue, ifFalse       int
	inspectedItems        int
}

func (m Monkey) hasItems() bool {
	return len(m.items) > 0
}

func (m *Monkey) addItem(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) processItem(factor int, reducer int) (int, int) {
	monkeyIndex := m.ifFalse
	currentItem := m.items[0]
	currentItem = m.newWorryLevelFunction(currentItem)
	currentItem = currentItem / factor
	if reducer != 0 {
		currentItem = currentItem % reducer
	}
	if currentItem%m.divisibleBy == 0 {
		monkeyIndex = m.ifTrue
	}
	m.items = m.items[1:]
	m.inspectedItems++

	return monkeyIndex, currentItem
}

func parseId(idMonkeyData string) int {
	var id int
	_, err := fmt.Sscanf(idMonkeyData, "Monkey %d:", &id)
	checkError(err)
	return id
}

func parseItems(itemsMonkeyData string) []int {
	var items []int

	itemsString := strings.TrimSpace(strings.Split(itemsMonkeyData, ":")[1])
	splittedItems := strings.Split(itemsString, ",")
	for _, item := range splittedItems {
		if item != "" {
			cleanedItem, err := strconv.Atoi(strings.TrimSpace(item))
			checkError(err)
			items = append(items, cleanedItem)
		}
	}
	return items
}

func add(a, b int) int { return a + b }

func mul(a, b int) int { return a * b }

var funcMap = map[string]func(int, int) int{
	"+": add,
	"*": mul,
}

func parseNewWorryLevelFunction(worryLevelMonkeyData string) func(int) int {
	funcString := strings.TrimSpace(strings.Split(worryLevelMonkeyData, ": ")[1])

	rightSide := strings.Split(funcString, " = ")[1]
	splittedRightSide := strings.Split(rightSide, " ")

	leftElement := splittedRightSide[0]
	operator := splittedRightSide[1]
	rightElement := splittedRightSide[2]

	if leftElement != "old" {
		a, err := strconv.Atoi(leftElement)
		checkError(err)
		return func(old int) int { return funcMap[operator](a, old) }
	} else if rightElement != "old" {
		b, err := strconv.Atoi(rightElement)
		checkError(err)
		return func(old int) int { return funcMap[operator](old, b) }
	} else {
		return func(old int) int { return funcMap[operator](old, old) }
	}
}

func parseDivisibleBy(divisibleByMonkeyData string) int {
	test := strings.TrimSpace(strings.Split(divisibleByMonkeyData, ":")[1])
	divisibleBy, err := strconv.Atoi(strings.Split(test, "divisible by ")[1])
	checkError(err)
	return divisibleBy
}

func parseIfCondition(ifConditionMonkeyData string) int {
	ifCondition, err := strconv.Atoi(strings.Split(strings.Split(ifConditionMonkeyData, ":")[1], "throw to monkey ")[1])
	checkError(err)
	return ifCondition
}

func parseMonkey(monkeyDataString string) Monkey {
	monkeyData := strings.Split(monkeyDataString, "\n")

	id := parseId(monkeyData[0])
	items := parseItems(monkeyData[1])
	newWorryLevelFunction := parseNewWorryLevelFunction(monkeyData[2])
	divisibleBy := parseDivisibleBy(monkeyData[3])
	ifTrue := parseIfCondition(monkeyData[4])
	ifFalse := parseIfCondition(monkeyData[5])

	return Monkey{id: id, items: items, newWorryLevelFunction: newWorryLevelFunction, divisibleBy: divisibleBy, ifTrue: ifTrue, ifFalse: ifFalse}
}

func processMonkeys(monkeys []Monkey, rounds int, factor int, reducer int) []Monkey {
	var index, currentItem int
	for round := 0; round < rounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			for monkeys[i].hasItems() {
				index, currentItem = monkeys[i].processItem(factor, reducer)

				monkeys[index].addItem(currentItem)
			}
		}
	}
	return monkeys
}

func main() {
	data, err := os.ReadFile(FilePath)
	checkError(err)
	startTime := time.Now()

	stringData := strings.TrimRight(string(data), "\n")
	monkeysData := strings.Split(stringData, "\n\n")
	monkeys := []Monkey{}

	// Monkeys parsing
	for _, monkeyData := range monkeysData {
		monkeys = append(monkeys, parseMonkey(monkeyData))
	}
	monkeysPart2 := make([]Monkey, len(monkeys))
	copy(monkeysPart2, monkeys)

	// Monkeys processing
	monkeys = processMonkeys(monkeys, 20, 3, 0)

	reducer := 1
	for _, monkey := range monkeys {
		reducer *= monkey.divisibleBy
	}
	monkeysPart2 = processMonkeys(monkeysPart2, 10000, 1, reducer)

	// Monkeys sorting
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedItems < monkeys[j].inspectedItems
	})
	sort.Slice(monkeysPart2, func(i, j int) bool {
		return monkeysPart2[i].inspectedItems < monkeysPart2[j].inspectedItems
	})

	part1answer := monkeys[len(monkeys)-1].inspectedItems * monkeys[len(monkeys)-2].inspectedItems
	part2answer := monkeysPart2[len(monkeysPart2)-1].inspectedItems * monkeysPart2[len(monkeysPart2)-2].inspectedItems
	fmt.Printf("%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Answer to part 1 is:", part1answer)
	fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Answer to part 2 is:", part2answer)

	fmt.Printf("\nTotal time elapsed: %v\n", time.Since(startTime))
}
