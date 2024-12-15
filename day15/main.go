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

type Vec2 struct {
	x int
	y int
}

func Part1(input string) int {
	m := strings.Split(input, "\n\n")[0]
	movesInput := strings.Split(input, "\n\n")[1]

	warehouse := [][]byte{}
	robot := Vec2{}
	for y, line := range strings.Split(m, "\n") {
		row := []byte{}
		for x, c := range line {
			if c == '@' {
				robot = Vec2{x, y}
				c = '.'
			}
			row = append(row, byte(c))
		}
		warehouse = append(warehouse, row)
	}

	moves := []Vec2{}
	for _, move := range movesInput {
		switch move {
		case 'v':
			moves = append(moves, Vec2{0, 1})
		case '^':
			moves = append(moves, Vec2{0, -1})
		case '>':
			moves = append(moves, Vec2{1, 0})
		case '<':
			moves = append(moves, Vec2{-1, 0})
		}
	}

	// PrintWarehouse(warehouse)
	// fmt.Println(moves)
	// fmt.Println("Robot:", robot)

	for _, move := range moves {
		MoveRobot(warehouse, &robot, move)
	}

	sum := 0
	for y, line := range warehouse {
		for x, c := range line {
			if c == 'O' {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func PrintWarehouse(warehouse [][]byte) {
	for _, line := range warehouse {
		fmt.Println(string(line))
	}
}

func MoveRobot(warehouse [][]byte, robot *Vec2, move Vec2) {
	pos := Vec2{robot.x + move.x, robot.y + move.y}
	if warehouse[pos.y][pos.x] == '.' {
		*robot = pos
		return
	}
	if warehouse[pos.y][pos.x] == '#' {
		return
	}

	if warehouse[pos.y][pos.x] != 'O' {
		panic("invalid")
	}

	oldPos := pos
	MoveRobot(warehouse, &pos, move)
	if oldPos.x == pos.x && oldPos.y == pos.y {
		return // Couldn't move
	}
	warehouse[pos.y][pos.x] = 'O'
	warehouse[oldPos.y][oldPos.x] = '.'
	*robot = oldPos
}

func Part2(input string) int {
	sum := 0
	return sum
}
