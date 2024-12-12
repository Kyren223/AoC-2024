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
	plantsMap := strings.Split(input, "\n")
	plantsMap = plantsMap[:len(plantsMap)-1]

	width := len(plantsMap[0])
	height := len(plantsMap)
	marks := make([][]bool, height)
	for y := range marks {
		marks[y] = make([]bool, width)
	}

	sum := 0
	for y, line := range plantsMap {
		for x, c := range line {
			area, permiter := Count(plantsMap, marks, int(c), x, y)
			sum += area * permiter
		}
	}

	return sum
}

func Count(plantsMap []string, marks [][]bool, c, x, y int) (int, int) {
	if marks[y][x] {
		return 0, 0
	}
	marks[y][x] = true

	top := 0
	if y-1 >= 0 {
		top = int(plantsMap[y-1][x])
	}
	bottom := 0
	if y+1 < len(plantsMap) {
		bottom = int(plantsMap[y+1][x])
	}
	left := 0
	if x-1 >= 0 {
		left = int(plantsMap[y][x-1])
	}
	right := 0
	if x+1 < len(plantsMap[0]) {
		right = int(plantsMap[y][x+1])
	}

	area := 1
	perimeter := 0
	if top == c {
		a, p := Count(plantsMap, marks, c, x, y-1)
		area += a
		perimeter += p
	} else {
		perimeter++
	}

	if bottom == c {
		a, p := Count(plantsMap, marks, c, x, y+1)
		area += a
		perimeter += p
	} else {
		perimeter++
	}

	if left == c {
		a, p := Count(plantsMap, marks, c, x-1, y)
		area += a
		perimeter += p
	} else {
		perimeter++
	}

	if right == c {
		a, p := Count(plantsMap, marks, c, x+1, y)
		area += a
		perimeter += p
	} else {
		perimeter++
	}

	return area, perimeter
}

func Part2(input string) int {
	sum := 0
	return sum
}
