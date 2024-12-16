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

	// file, err = os.ReadFile("input.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// input = string(file)
	// fmt.Println("Part 1:", Part1(input))
	//
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

func (v Vec2) Add(o Vec2) Vec2 {
	return Vec2{v.x + o.x, v.y + o.y}
}

func Part1(input string) int {
	maze := strings.Split(input, "\n")
	maze = maze[:len(maze)-1]

	visisted := map[Vec2]struct{}{}
	start := Vec2{}
	dir := Vec2{1, 0}
	for y, line := range maze {
		for x, c := range line {
			if c == 'S' {
				start = Vec2{x, y}
			}
		}
	}

	// for _, line := range maze {
	// 	fmt.Println(line)
	// }

	sum := DFS(maze, visisted, start, dir, 0)
	return sum
}

func RotateLeft(dir Vec2) Vec2 {
	if dir.x == 0 && dir.y == 1 {
		return Vec2{1, 0}
	} else if dir.x == 1 && dir.y == 0 {
		return Vec2{0, -1}
	} else if dir.x == 0 && dir.y == -1 {
		return Vec2{-1, 0}
	} else if dir.x == -1 && dir.y == 0 {
		return Vec2{0, 1}
	} else {
		panic(dir)
	}
}

func RotateRight(dir Vec2) Vec2 {
	if dir.x == 0 && dir.y == 1 {
		return Vec2{-1, 0}
	} else if dir.x == -1 && dir.y == 0 {
		return Vec2{0, -1}
	} else if dir.x == 0 && dir.y == -1 {
		return Vec2{1, 0}
	} else if dir.x == 1 && dir.y == 0 {
		return Vec2{0, 1}
	} else {
		panic(dir)
	}
}

func DFS(maze []string, visited map[Vec2]struct{}, pos, dir Vec2, depth int) int {
	// fmt.Println(depth, pos, dir)
	const infinity = 100000000000

	if pos.x < 0 || pos.x >= len(maze[0]) || pos.y < 0 || pos.y >= len(maze) {
		return infinity
	}

	if _, ok := visited[pos]; ok {
		return infinity
	}

	c := maze[pos.y][pos.x]
	switch c {
	case '#':
		return infinity
	case 'E':
		return 0
	}

	visited[pos] = struct{}{}

	// Normal case
	left := RotateLeft(dir)
	right := RotateRight(dir)
	// back := RotateLeft(RotateLeft(dir))

	scoreForward := DFS(maze, visited, pos.Add(dir), dir, depth+1) + 1
	scoreLeft := DFS(maze, visited, pos.Add(left), left, depth+1) + 1001
	scoreRight := DFS(maze, visited, pos.Add(right), right, depth+1) + 1001
	// scoreBack := DFS(maze, visited, pos.Add(back), back, depth+1) + 2001

	if depth == 30 {
		fmt.Println(depth, pos, scoreForward, scoreLeft, scoreRight)
	}

	return min(scoreForward, scoreLeft, scoreRight)
}

func Part2(input string) int {
	sum := 0
	return sum
}
