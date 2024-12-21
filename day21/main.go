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

func Part1(input string) int {
	codes := strings.Split(input, "\n")
	codes = codes[:len(codes)-1]

	sum := 0
	return sum
}

var numpadPaths = map[string][][]byte{}

func Init() {
	numpad := strings.Split("789\n456\n123\n 0A", "\n")
	for y, line := range numpad {
		for x, c := range line {
			if c == ' ' {
				continue
			}
			for y2, line2 := range numpad {
				for x2, c2 := range line2 {
					if c2 == ' ' {
						continue
					}
					Precompute(numpad, x, y, x2, y2)
				}
			}
		}
	}
}

func Precompute(numpad []string, x1, y1, x2, y2 int, depth int, path []byte) [][]byte {
	if depth > 6 {
		return [][]byte{}
	}
	if x1 < 0 || y1 < 0 || y1 >= len(numpad) || x1 >= len(numpad[0]) {
		return [][]byte{}
	}
	if numpad[y1][x1] == ' ' {
		return [][]byte{}
	}

	if x1 == x2 && y1 == y2 {
		return [][]byte{append(path, 'A')}
	}

	paths := [4][]byte{}
	for _, move := range "<^>v" {
		switch move {
		case '<':
			Precompute(numpad, x1-1, y1, x2, y2, depth+1, path)
		case '>':
			cx++
		case '^':
			y--
		case 'v':
			y++
		}
	}
}

func ShortestPath(code string) []byte {
	return nil
}

func Part2(input string) int {
	sum := 0
	return sum
}
