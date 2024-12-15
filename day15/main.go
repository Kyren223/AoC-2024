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

	file, err = os.ReadFile("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input = string(file)
	fmt.Println("Part 2 Example:", Part2(input))

	file, err = os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input = string(file)
	fmt.Println("Part 2:", Part2(input))
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

func PrintWarehouse(warehouse [][]byte, robot Vec2) {
	for y, line := range warehouse {
		if robot.y == y {
			for x, c := range line {
				if robot.x == x {
					fmt.Print(string('@'))
					continue
				}
				fmt.Print(string(c))
			}
			fmt.Println()
			continue
		}
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
	m := strings.Split(input, "\n\n")[0]
	movesInput := strings.Split(input, "\n\n")[1]

	warehouse := [][]byte{}
	for _, line := range strings.Split(m, "\n") {
		row := []byte{}
		for _, c := range line {
			switch c {
			case '#':
				row = append(row, '#', '#')
			case 'O':
				row = append(row, '[', ']')
			case '@':
				row = append(row, '@', '.')
			case '.':
				row = append(row, '.', '.')
			}
		}
		warehouse = append(warehouse, row)
	}

	robot := Vec2{}
	for y, line := range warehouse {
		for x, c := range line {
			if c == '@' {
				warehouse[y][x] = '.'
				robot = Vec2{x, y}
			}
		}
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

	// PrintWarehouse(warehouse, robot)
	// fmt.Println(moves)
	fmt.Println("Robot:", robot)

	for _, move := range moves {
		MoveRobot2(warehouse, &robot, move, true)
		// fmt.Println(move)
		// PrintWarehouse(warehouse, robot)
		// fmt.Println()
	}

	sum := 0
	for y, line := range warehouse {
		for x, c := range line {
			if c == '[' {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func MoveRobot2(warehouse [][]byte, robot *Vec2, move Vec2, shouldMove bool) {
	pos := Vec2{robot.x + move.x, robot.y + move.y}
	if warehouse[pos.y][pos.x] == '.' {
		*robot = pos
		return
	}
	if warehouse[pos.y][pos.x] == '#' {
		return
	}

	if warehouse[pos.y][pos.x] != '[' && warehouse[pos.y][pos.x] != ']' {
		panic(string(warehouse[pos.y][pos.x]))
	}

	if move.y == 0 {
		end := pos
		end.x += move.x
		oldEnd := end
		MoveRobot2(warehouse, &end, move, true)
		if oldEnd.x == end.x && oldEnd.y == end.y {
			return // Couldn't move
		}
		warehouse[end.y][end.x] = warehouse[oldEnd.y][oldEnd.x]
		warehouse[oldEnd.y][oldEnd.x] = warehouse[pos.y][pos.x]
		warehouse[pos.y][pos.x] = '.'
		*robot = pos
		return
	}

	left := Vec2{}
	right := Vec2{}
	if warehouse[pos.y][pos.x] == '[' {
		left = pos
		right = Vec2{pos.x + 1, pos.y}
	} else {
		left = Vec2{pos.x - 1, pos.y}
		right = pos
	}

	oldLeft := left
	oldRight := right
	MoveRobot2(warehouse, &left, move, false)
	MoveRobot2(warehouse, &right, move, false)
	canLeftMove := !(oldLeft.x == left.x && oldLeft.y == left.y)
	canRightMove := !(oldRight.x == right.x && oldRight.y == right.y)
	if !canLeftMove || !canRightMove {
		return // Couldn't move
	}

	if shouldMove {
		left = oldLeft
		right = oldRight
		MoveRobot2(warehouse, &left, move, true)
		MoveRobot2(warehouse, &right, move, true)
		warehouse[left.y][left.x] = '['
		warehouse[right.y][right.x] = ']'
		warehouse[oldLeft.y][oldLeft.x] = '.'
		warehouse[oldRight.y][oldRight.x] = '.'
	}
	*robot = pos
}
