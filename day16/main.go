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

	file, err = os.ReadFile("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input = string(file)
	fmt.Println("Part 2 Example:", Part2(input))

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
	// fmt.Println(start, end)
	// fmt.Println(visisted)

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
	// fmt.Println(start, end)
	// fmt.Println(visisted)

	// Manually add +1000 cuz it needs to rotate once to be "north"
	sum := visisted[start] + 1000

	marked := map[Vec2]bool{}
	DFS2(maze, visisted, marked, start, sum, 0)

	for y, line := range maze {
		for x, c := range line {
			if mark, ok := marked[Vec2{x, y}]; ok && mark {
				fmt.Print(string('O'))
			} else {
				fmt.Print(string(c))
			}
			// if _, ok := visisted[Vec2{x, y}]; ok {
			// 	fmt.Print(string('O'))
			// } else {
			// 	fmt.Print(string(c))
			// }
		}
		fmt.Println()
	}

	fmt.Println("HMM", visisted[Vec2{1, 10}], visisted[Vec2{2, 11}])

	sum = 0
	for _, mark := range marked {
		if mark {
			sum++
		}
	}

	return sum
}

func DFS2(maze []string, visited map[Vec2]int, marked map[Vec2]bool, pos Vec2, minimum int, depth int) {
	// if depth > 3 {
	// 	return
	// }
	if _, ok := marked[pos]; ok {
		return
	}

	cost, ok := visited[pos]
	if !ok {
		return
	}

	marked[pos] = cost == minimum
	if cost > minimum {
		// panic("wrong!")
		return
	}

	left := pos
	left.x -= 1
	right := pos
	right.x += 1
	top := pos
	top.y -= 1
	bottom := pos
	bottom.y += 1

	newMin := minimum
	if cost, ok := visited[left]; ok && cost < minimum {
		if newMin == minimum {
			newMin = cost
		}
		newMin = min(newMin, cost)
	}
	if cost, ok := visited[right]; ok && cost < minimum {
		if newMin == minimum {
			newMin = cost
		}
		newMin = min(newMin, cost)
		if depth == 3 {
			fmt.Println("RIGHT", newMin, cost, pos)
		}
	}
	if cost, ok := visited[top]; ok && cost < minimum {
		if newMin == minimum {
			newMin = cost
		}
		newMin = min(newMin, cost)
		if depth == 3 {
			fmt.Println("TOP", newMin, cost)
		}
	}
	if cost, ok := visited[bottom]; ok && cost < minimum {
		if newMin == minimum {
			newMin = cost
		}
		newMin = min(newMin, cost)
	}

	if newMin == minimum {
		fmt.Println("Test")
		return // Dead end
	}

	DFS2(maze, visited, marked, left, newMin, depth+1)
	DFS2(maze, visited, marked, right, newMin, depth+1)
	DFS2(maze, visited, marked, top, newMin, depth+1)
	DFS2(maze, visited, marked, bottom, newMin, depth+1)
}

// func DFS2(maze []string, visited map[Vec2]int, visited2 map[Vec2]struct{}, mins map[int]int, pos Vec2, depth int) {
// 	if _, ok := visited2[pos]; ok {
// 		return
// 	}
// 	visited2[pos] = struct{}{}
//
// 	if m, ok := mins[depth]; ok {
// 		mins[depth] = min(m, visited[pos])
// 	} else {
// 		mins[depth] = visited[pos]
// 	}
//
// 	left := pos
// 	left.x -= 1
// 	right := pos
// 	right.x += 1
// 	top := pos
// 	top.y -= 1
// 	bottom := pos
// 	bottom.y += 1
//
// 	DFS2(maze, visited, marked, left, newMin, depth+1)
// 	DFS2(maze, visited, marked, right, newMin, depth+1)
// 	DFS2(maze, visited, marked, top, newMin, depth+1)
// 	DFS2(maze, visited, marked, bottom, newMin, depth+1)
// }
//
// func DFS3(maze []string, visited map[Vec2]int, marked map[Vec2]bool, pos Vec2, depth int) {
// 	// if depth > 3 {
// 	// 	return
// 	// }
// 	if _, ok := marked[pos]; ok {
// 		return
// 	}
//
// 	cost, ok := visited[pos]
// 	if !ok {
// 		return
// 	}
//
// 	marked[pos] = cost == minimum
// 	if cost > minimum {
// 		// panic("wrong!")
// 		return
// 	}
//
// 	left := pos
// 	left.x -= 1
// 	right := pos
// 	right.x += 1
// 	top := pos
// 	top.y -= 1
// 	bottom := pos
// 	bottom.y += 1
//
// 	DFS2(maze, visited, marked, left, newMin, depth+1)
// 	DFS2(maze, visited, marked, right, newMin, depth+1)
// 	DFS2(maze, visited, marked, top, newMin, depth+1)
// 	DFS2(maze, visited, marked, bottom, newMin, depth+1)
// }
