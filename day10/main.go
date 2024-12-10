package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input := string(file)
	fmt.Println("Part 1 Example:", Part1(input))

	file, err = os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input = string(file)
	fmt.Println("Part 1:", Part1(input))

	// file, err = os.ReadFile("example.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// input = string(file)
	// fmt.Println("Part 2 Example:", Part2(input))
	//
	// file, err = os.ReadFile("input.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// input = string(file)
	// fmt.Println("Part 2:", Part2(input))
}

func Part1(input string) int {
	topMap := strings.Split(input, "\n")
	topMap = topMap[:len(topMap)-1]

	// fmt.Println("Width:", len(antenasMap[0]), "Height:", len(antenasMap))
	width := len(topMap[0])
	height := len(topMap)

	sum := 0
	for y, line := range topMap {
		for x, c := range line {
			if c != '0' {
				continue
			}

			marks := make([][]bool, height)
			for y := range marks {
				marks[y] = make([]bool, width)
			}
			Traverse(topMap, marks, x, y, byte(c))

			for _, mark := range marks {
				for _, marked := range mark {
					if marked {
						sum++
					}
				}
			}
		}
	}

	return sum
}

func Traverse(topMap []string, marks [][]bool, x, y int, height byte) {
	if x < 0 || x >= len(topMap[0]) || y < 0 || y >= len(topMap) {
		return
	}
	if topMap[y][x] != height {
		return
	}
	if topMap[y][x] == '9' {
		marks[y][x] = true
		return
	}

	Traverse(topMap, marks, x+1, y, height+1)
	Traverse(topMap, marks, x-1, y, height+1)
	Traverse(topMap, marks, x, y+1, height+1)
	Traverse(topMap, marks, x, y-1, height+1)
}

func Part2(input string) int {
	sum := 0
	return sum
}
