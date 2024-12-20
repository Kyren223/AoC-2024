package main

import (
	"container/heap"
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

	file, err = os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input = string(file)
	fmt.Println("Part 2:", Part2(input))
}

type Item struct {
	fromPosition Pos
	position     Pos
	priority     int
	index        int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, fromPosition, position Pos, priority int) {
	item.fromPosition = fromPosition
	item.position = position
	item.priority = priority
	heap.Fix(pq, item.index)
}

type Pos struct {
	x int
	y int
}

func Part1(input string) int {
	var racetrack [][]byte
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		racetrack = append(racetrack, []byte(line))
	}

	start := Pos{}
	end := Pos{}
	for y, line := range racetrack {
		for x, c := range line {
			switch c {
			case 'S':
				start = Pos{x, y}
			case 'E':
				end = Pos{x, y}
			}
		}
	}

	Print(racetrack)
	fmt.Println(start, end)

	fmt.Println(len(racetrack), len(racetrack[0]))
	baseline := Dikstra(start, end, racetrack)

	sum := 0
	for y, line := range racetrack {
		for x, c := range line {
			if c != '#' {
				continue
			}
			racetrack[y][x] = '.'
			cheated := Dikstra(start, end, racetrack)
			racetrack[y][x] = '#'

			if baseline-cheated >= 100 {
				sum++
			}
		}
	}

	return sum
}

func Print(racetrack [][]byte) {
	for _, line := range racetrack {
		fmt.Println(string(line))
	}
}

func Dikstra(s Pos, e Pos, plane [][]byte) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// heap.Push(&pq, &Item{Pos{}, s, 0, -1})
	visited := make([][]int, len(plane))
	for i := 0; i < len(plane); i++ {
		visited[i] = make([]int, len(plane[0]))
		for j := 0; j < len(plane[0]); j++ {
			visited[i][j] = -1
		}
	}

	current := s
	distance := 0

	for {
		if current.x == e.x && current.y == e.y {
			return distance
		}

		visited[current.y][current.x] = 0

		arr := []Pos{current, current, current, current}
		arr[0].y -= 1
		arr[1].y += 1
		arr[2].x -= 1
		arr[3].x += 1

		for _, pos := range arr {
			if pos.x < 0 || pos.y < 0 || pos.y >= len(plane) || pos.x >= len(plane[0]) {
				continue
			}
			if plane[pos.y][pos.x] == '#' {
				continue
			}

			var i *Item = nil
			for _, item := range pq {
				if item.position.x == pos.x && item.position.y == pos.y {
					i = item
					break
				}
			}
			if i != nil && i.priority > distance {
				pq.update(i, current, pos, distance+1)
			} else if visited[pos.y][pos.x] == -1 {
				heap.Push(&pq, &Item{current, pos, distance + 1, -1})
			}
		}

		// for i, item := range pq {
		// 	fmt.Printf("%v:%v/%v, ", i, item.position, item.priority)
		// }
		// fmt.Println()

		if pq.Len() == 0 {
			break
		}

		item := heap.Pop(&pq).(*Item)
		current = item.position
		distance = item.priority
	}

	return -1
}

func Part2(input string) int {
	var racetrack [][]byte
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		racetrack = append(racetrack, []byte(line))
	}

	start := Pos{}
	end := Pos{}
	for y, line := range racetrack {
		for x, c := range line {
			switch c {
			case 'S':
				start = Pos{x, y}
			case 'E':
				end = Pos{x, y}
			}
		}
	}

	// Print(racetrack)
	// fmt.Println(start, end)

	// fmt.Println(len(racetrack), len(racetrack[0]))
	baseline := Dikstra(start, end, racetrack)

	iters := 0
	sum := 0
	for y, line := range racetrack {
		for x, c := range line {
			if c == '#' {
				continue
			}
			for y2 := y - 21; y2 < y+21; y2++ {
				for x2 := x - 21; x2 < x+21; x2++ {
					if x2 < 0 || y2 < 0 || y2 >= len(racetrack) || x2 >= len(racetrack[0]){
						continue
					}
					if racetrack[y2][x2] == '#' {
						continue
					}
					distance := (x2-x)*(x2-x) + (y2-y)*(y2-y)
					if distance > 20*20 {
						continue
					}

					distance = DikstraNoWalls(Pos{x, y}, Pos{x2, y2}, racetrack)
					if distance == -1 {
						continue
					}

					cheated := Dikstra2(start, end, racetrack, Pos{x, y}, Pos{x2, y2}, distance)

					if baseline-cheated >= 100 {
						sum++
					}
				}
			}
			fmt.Println(iters)
			iters++
		}
	}

	return sum
}

func Dikstra2(s Pos, e Pos, plane [][]byte, cheatS Pos, cheatE Pos, cheatD int) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	visited := make([][]int, len(plane))
	for i := 0; i < len(plane); i++ {
		visited[i] = make([]int, len(plane[0]))
		for j := 0; j < len(plane[0]); j++ {
			visited[i][j] = -1
		}
	}

	current := s
	distance := 0

	for {
		if current.x == e.x && current.y == e.y {
			return distance
		}

		visited[current.y][current.x] = 0

		if current.x == cheatS.x && current.y == cheatS.y {
			var i *Item = nil
			for _, item := range pq {
				if item.position.x == cheatE.x && item.position.y == cheatE.y {
					i = item
					break
				}
			}
			if i != nil && i.priority > distance+cheatD {
				pq.update(i, current, cheatE, distance+cheatD)
			} else if visited[cheatE.y][cheatE.x] == -1 {
				heap.Push(&pq, &Item{current, cheatE, distance + cheatD, -1})
			}
		}

		arr := []Pos{current, current, current, current}
		arr[0].y -= 1
		arr[1].y += 1
		arr[2].x -= 1
		arr[3].x += 1

		for _, pos := range arr {
			if pos.x < 0 || pos.y < 0 || pos.y >= len(plane) || pos.x >= len(plane[0]) {
				continue
			}
			if plane[pos.y][pos.x] == '#' {
				continue
			}

			var i *Item = nil
			for _, item := range pq {
				if item.position.x == pos.x && item.position.y == pos.y {
					i = item
					break
				}
			}
			if i != nil && i.priority > distance {
				pq.update(i, current, pos, distance+1)
			} else if visited[pos.y][pos.x] == -1 {
				heap.Push(&pq, &Item{current, pos, distance + 1, -1})
			}
		}

		// for i, item := range pq {
		// 	fmt.Printf("%v:%v/%v, ", i, item.position, item.priority)
		// }
		// fmt.Println()

		if pq.Len() == 0 {
			break
		}

		item := heap.Pop(&pq).(*Item)
		current = item.position
		distance = item.priority
	}

	return -1
}

func DikstraNoWalls(s Pos, e Pos, plane [][]byte) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	visited := make([][]int, len(plane))
	for i := 0; i < len(plane); i++ {
		visited[i] = make([]int, len(plane[0]))
		for j := 0; j < len(plane[0]); j++ {
			visited[i][j] = -1
		}
	}

	current := s
	distance := 0

	for {
		if current.x == e.x && current.y == e.y {
			return distance
		}

		visited[current.y][current.x] = 0

		arr := []Pos{current, current, current, current}
		arr[0].y -= 1
		arr[1].y += 1
		arr[2].x -= 1
		arr[3].x += 1

		for _, pos := range arr {
			if pos.x < 0 || pos.y < 0 || pos.y >= len(plane) || pos.x >= len(plane[0]) {
				continue
			}

			var i *Item = nil
			for _, item := range pq {
				if item.position.x == pos.x && item.position.y == pos.y {
					i = item
					break
				}
			}
			if i != nil && i.priority > distance {
				pq.update(i, current, pos, distance+1)
			} else if visited[pos.y][pos.x] == -1 {
				heap.Push(&pq, &Item{current, pos, distance + 1, -1})
			}
		}

		// for i, item := range pq {
		// 	fmt.Printf("%v:%v/%v, ", i, item.position, item.priority)
		// }
		// fmt.Println()

		if pq.Len() == 0 {
			break
		}

		item := heap.Pop(&pq).(*Item)
		current = item.position
		distance = item.priority

		if distance > 20 {
			return -1
		}
	}

	return -1
}
