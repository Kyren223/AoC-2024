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

func (v Vec2) Add(o Vec2) Vec2 {
	return Vec2{v.x + o.x, v.y + o.y}
}

func Part1(input string) int {
	maze := strings.Split(input, "\n")
	maze = maze[:len(maze)-1]

	visisted := map[Vec2]int{}
	start := Vec2{}
	end := Vec2{}
	for y, line := range maze {
		for x, c := range line {
			switch c {
			case 'E':
				end = Vec2{x, y}
			case 'S':
				start = Vec2{x, y}
			}
		}
	}

	DFS(maze, visisted, end, 0, 1, 1, 0)
	fmt.Println(start, end)
	fmt.Println(visisted)

	// Manually add +1000 cuz it needs to rotate once to be "north"
	sum := visisted[start] + 1000
	return sum
}

func DFS(maze []string, visited map[Vec2]int, pos Vec2, cost, xcost, ycost int, depth int) {
	if pos.x < 0 || pos.x >= len(maze[0]) || pos.y < 0 || pos.y >= len(maze) {
		return
	}
	c := maze[pos.y][pos.x]
	if c == '#' {
		return
	}

	if pos.x == 13 && pos.y == 13 {
		fmt.Println("13YO")
	}

	if visistedCost, ok := visited[pos]; ok && visistedCost <= cost {
		return
	}

	visited[pos] = cost
	if c == 'S' {

	}

	left := pos
	left.x -= 1
	right := pos
	right.x += 1
	top := pos
	top.y -= 1
	bottom := pos
	bottom.y += 1

	DFS(maze, visited, left, cost+xcost, 1, 1001, depth+1)
	DFS(maze, visited, right, cost+xcost, 1, 1001, depth+1)
	DFS(maze, visited, top, cost+ycost, 1001, 1, depth+1)
	DFS(maze, visited, bottom, cost+ycost, 1001, 1, depth+1)
}

func Part2(input string) int {
	sum := 0
	return sum
}
