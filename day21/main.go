package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
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

	for key, val := range numpadPaths {
		fmt.Print(key)
		fmt.Print(" [")
		for i, s := range val {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Print(string(s))
		}
		fmt.Print("]\n")
	}

	for key, val := range keypadPaths {
		fmt.Print(key)
		fmt.Print(" [")
		for i, s := range val {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Print(string(s))
		}
		fmt.Print("]\n")
	}

	sum := 0
	for _, code := range codes {
		path := ShortestPath(code)
		fmt.Println(code, string(path))

		num, _ := strconv.ParseInt(code[:3], 10, 64)
		sum += len(path) * int(num)
		break
	}

	return sum
}

var (
	numpadPaths = map[string][][]byte{}
	keypadPaths = map[string][][]byte{}
)

func init() {
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
					paths := Precompute(numpad, x, y, x2, y2, 0, []byte(""))
					numpadPaths[string([]byte{byte(c), byte(c2)})] = paths
				}
			}
		}
	}

	keypad := strings.Split(" ^A\n<v>", "\n")
	for y, line := range keypad {
		for x, c := range line {
			if c == ' ' {
				continue
			}
			for y2, line2 := range keypad {
				for x2, c2 := range line2 {
					if c2 == ' ' {
						continue
					}
					paths := Precompute(keypad, x, y, x2, y2, 0, []byte(""))
					keypadPaths[string([]byte{byte(c), byte(c2)})] = paths
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

	paths := [][]byte{}
	for _, move := range "<^>v" {
		cx := x1
		cy := y1
		switch move {
		case '<':
			cx--
		case '>':
			cx++
		case '^':
			cy--
		case 'v':
			cy++
		}
		copyPath := slices.Clone(path)
		copyPath = append(copyPath, byte(move))
		ps := Precompute(numpad, cx, cy, x2, y2, depth+1, copyPath)
		newPaths := [][]byte{}
		for _, p := range ps {
			if len(newPaths) == 0 {
				newPaths = append(newPaths, p)
				continue
			}
			if len(newPaths[0]) == len(p) {
				newPaths = append(newPaths, p)
				continue
			}
			if len(newPaths[0]) > len(p) {
				newPaths = [][]byte{p}
			}
		}
		paths = append(paths, newPaths...)
	}

	return paths
}

func Get(m map[string][][]byte, a, b byte) [][]byte {
	return m[string([]byte{a, b})]
}

func ShortestPath(code string) []byte {
	firstPrev := byte('A')
	firstMoves := []byte{}
	for i := range code {
		firstPaths := Get(numpadPaths, firstPrev, code[i])
		var bestFirstPaths []byte = nil

		for _, firstPath := range firstPaths {
			secondPrev := byte('A')
			secondMoves := []byte{}
			for j := range firstPath {
				secondPaths := Get(keypadPaths, secondPrev, firstPath[j])
				var bestSecondPath []byte = nil

				for _, secondPath := range secondPaths {
					thirdPrev := byte('A')
					thirdMoves := []byte{}
					for k := range secondPath {
						thirdPaths := Get(keypadPaths, thirdPrev, secondPath[k])
						// All lengths are equal so it doesn't matter
						// Just take the first
						bestThirdPath := thirdPaths[0]
						thirdMoves = append(thirdMoves, bestThirdPath...)

						thirdPrev = secondPath[k]
					}

					if bestSecondPath == nil || len(thirdMoves) < len(bestSecondPath) {
						bestSecondPath = thirdMoves
					}
				}

				secondMoves = append(secondMoves, bestSecondPath...)

				secondPrev = firstPath[j]

				if bestFirstPaths == nil || len(secondMoves) < len(bestFirstPaths) {
					bestFirstPaths = secondMoves
				}
			}
		}

		firstMoves = append(firstMoves, bestFirstPaths...)

		firstPrev = code[i]
	}

	return firstMoves
}

func Part2(input string) int {
	sum := 0
	return sum
}
