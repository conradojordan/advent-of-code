package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const FilePath = "2022/day07/input.txt"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

type file struct {
	name string
	size int
}

type dirInfo file

type directory struct {
	name     string
	parent   *directory
	files    []file
	children []*directory
}

func (d *directory) addFile(fileInfo []string) {
	fileSize, err := strconv.Atoi(fileInfo[0])
	checkError(err)
	newFile := file{name: fileInfo[1], size: fileSize}
	d.files = append(d.files, newFile)
}

func (d *directory) addChild(childInfo []string) {
	newChild := directory{name: childInfo[1], parent: d}
	d.children = append(d.children, &newChild)
}

func (d *directory) getChildByName(name string) *directory {
	for _, child := range d.children {
		if child.name == name {
			return child
		}
	}
	return nil
}

func (d *directory) generateDirInfos(dirInfos *[]dirInfo) int {
	dirInfo := dirInfo{name: d.name}
	for _, file := range d.files {
		dirInfo.size += file.size
	}
	for _, dir := range d.children {
		dirInfo.size += dir.generateDirInfos(dirInfos)
	}
	(*dirInfos) = append((*dirInfos), dirInfo)
	return dirInfo.size
}

func main() {
	data, err := os.ReadFile(FilePath)
	checkError(err)
	startTime := time.Now()

	stringData := strings.TrimRight(string(data), "\n")
	commands := strings.Split(stringData, "\n")

	root := directory{name: "/"}
	var currentDirectory *directory = &root

	for i := 1; i < len(commands); i++ {
		args := strings.Split(commands[i], " ")
		if args[0] == "$" {
			if args[1] == "cd" {
				if args[2] == ".." {
					currentDirectory = currentDirectory.parent
				} else {
					currentDirectory = currentDirectory.getChildByName(args[2])
				}
			}
		} else if args[0] == "dir" {
			currentDirectory.addChild(args)
		} else {
			currentDirectory.addFile(args)
		}
	}

	dirInfos := make([]dirInfo, 0, 200)
	root.generateDirInfos(&dirInfos)
	sort.Slice(dirInfos, func(i, j int) bool { return dirInfos[i].size < dirInfos[j].size })

	totalSpace := 70000000
	usedSpace := dirInfos[len(dirInfos)-1].size
	unusedSpace := totalSpace - usedSpace

	requiredSpace := 30000000
	mustFreeSpace := requiredSpace - unusedSpace

	var part1answer, part2answer int
	for _, di := range dirInfos {
		if di.size <= 100000 {
			part1answer += di.size
		} else {
			break
		}
	}
	for _, di := range dirInfos {
		if di.size >= mustFreeSpace {
			part2answer = di.size
			break
		}
	}

	fmt.Printf("%sPart 1%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Println("Sum of directories sizes with at most 100000 of size:", part1answer)
	fmt.Printf("%sPart 2%s\n", strings.Repeat("-", 10), strings.Repeat("-", 10))
	fmt.Printf("Smallest directory size larger than %d: %d\n", mustFreeSpace, part2answer)

	fmt.Printf("\nTotal time elapsed: %v\n", time.Since(startTime))
}
