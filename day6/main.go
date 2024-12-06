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

type Pos struct {
	x int
	y int
}

func (p *Pos) RotateRight() {
	if p.y == -1 {
		p.x = 1
		p.y = 0
	} else if p.x == 1 {
		p.y = 1
		p.x = 0
	} else if p.y == 1 {
		p.x = -1
		p.y = 0
	} else if p.x == -1 {
		p.y = -1
		p.x = 0
	}
}

func Part1(input string) int {
	guardMap := strings.Split(input, "\n")
	guardMap = guardMap[:len(guardMap)-1]

	// fmt.Println("Width:", len(guardMap[0]), "Height:", len(guardMap))

	marks := make([][]bool, len(guardMap))
	for y := range marks {
		marks[y] = make([]bool, len(guardMap[0]))
	}

	guard := Pos{}
	dir := Pos{x: 0, y: -1}

outer:
	for y, line := range guardMap {
		for x, c := range line {
			if c == '^' {
				guard.x = x
				guard.y = y
				break outer
			}
		}
	}

	marks[guard.y][guard.x] = true
	for {
		x := guard.x + dir.x
		y := guard.y + dir.y

		// Bounds check
		if x < 0 || x > len(guardMap[0])-1 || y < 0 || y > len(guardMap)-1 {
			break
		}

		c := guardMap[y][x]
		if c == '#' {
			dir.RotateRight()
		} else {
			guard.x = x
			guard.y = y
			marks[y][x] = true
		}
	}

	sum := 0
	for _, mark := range marks {
		for _, marked := range mark {
			if marked {
				sum++
			}
		}
	}

	return sum
}

func Part2(input string) int {
	return 0
}
