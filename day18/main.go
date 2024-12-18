package main

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input := string(file)
	fmt.Println("Part 1 Example:", Part1(input, true))

	file, err = os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input = string(file)
	fmt.Println("Part 1:", Part1(input, false))

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

func Part1(input string, example bool) int {
	var positions []Pos
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		x, _ := strconv.ParseInt(strings.Split(line, ",")[0], 10, 64)
		y, _ := strconv.ParseInt(strings.Split(line, ",")[1], 10, 64)
		positions = append(positions, Pos{x: int(x), y: int(y)})
	}

	length := 71
	if example {
		length = 7
	}

	var plane [][]bool
	for y := 0; y < length; y++ {
		plane = append(plane, []bool{})
		for x := 0; x < length; x++ {
			wall := false
			for i, pos := range positions {
				if (i >= 1024 && !example) || (i >= 12 && example) {
					break
				}
				if pos.x == x && pos.y == y {
					wall = true
				}
			}
			plane[y] = append(plane[y], wall)
		}
	}

	PrintPlane(plane)

	e := Pos{70, 70}
	if example {
		e = Pos{6, 6}
	}

	return Dikstra(Pos{0, 0}, e, length, plane)
}

func PrintPlane(plane [][]bool) {
	for _, line := range plane {
		for _, b := range line {
			if b {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func Dikstra(s Pos, e Pos, length int, plane [][]bool) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// heap.Push(&pq, &Item{Pos{}, s, 0, -1})
	visited := make([][]int, length)
	for i := 0; i < length; i++ {
		visited[i] = make([]int, length)
		for j := 0; j < length; j++ {
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
			if pos.x < 0 || pos.y < 0 || pos.y >= length || pos.x >= length {
				continue
			}
			if plane[pos.y][pos.x] {
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
	sum := 0
	return sum
}
