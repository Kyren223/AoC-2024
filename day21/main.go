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

	// for key, val := range numpadPaths {
	// 	fmt.Print(key)
	// 	fmt.Print(" [")
	// 	for i, s := range val {
	// 		if i != 0 {
	// 			fmt.Print(" ")
	// 		}
	// 		fmt.Print(string(s))
	// 	}
	// 	fmt.Print("]\n")
	// }

	// for key, val := range keypadPaths {
	// 	fmt.Print(key)
	// 	fmt.Print(" [")
	// 	for i, s := range val {
	// 		if i != 0 {
	// 			fmt.Print(" ")
	// 		}
	// 		fmt.Print(string(s))
	// 	}
	// 	fmt.Print("]\n")
	// }

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
	// firstPrev := byte('A')
	firstMoves := GetBestMoves(numpadPaths, code, 'A', 0, []byte(""))
	for key, val := range firstMoves {
		fmt.Println(key, string(val))
	}

	secondMoves := [][]byte{}
	for _, firstMove := range firstMoves {
		moves := GetBestMoves(keypadPaths, string(firstMove), 'A', 0, []byte(""))
		secondMoves = append(secondMoves, moves...)
	}
	bestSecondMoves := [][]byte{}
	for _, path := range secondMoves {
		if len(bestSecondMoves) == 0 {
			bestSecondMoves = append(bestSecondMoves, path)
			continue
		}
		if len(bestSecondMoves[0]) == len(path) {
			bestSecondMoves = append(bestSecondMoves, path)
			continue
		}
		if len(bestSecondMoves[0]) > len(path) {
			bestSecondMoves = [][]byte{path}
		}
	}
	for key, val := range bestSecondMoves {
		fmt.Println(key, string(val))
	}

	// thirdMoves := [][]byte{}
	for _, secondMove := range bestSecondMoves {
		moves := GetBestMoves(keypadPaths, string(secondMove), 'A', 0, []byte(""))
		fmt.Println(len(moves))
		// thirdMoves = append(thirdMoves, moves...)
	}
	// bestthirdMoves := [][]byte{}
	// for _, path := range thirdMoves {
	// 	if len(bestthirdMoves) == 0 {
	// 		bestthirdMoves = append(bestthirdMoves, path)
	// 		continue
	// 	}
	// 	if len(bestthirdMoves[0]) == len(path) {
	// 		bestthirdMoves = append(bestthirdMoves, path)
	// 		continue
	// 	}
	// 	if len(bestthirdMoves[0]) > len(path) {
	// 		bestthirdMoves = [][]byte{path}
	// 	}
	// }
	// for key, val := range bestthirdMoves {
	// 	fmt.Println(key, string(val))
	// }

	return nil
}

func GetBestMoves(cache map[string][][]byte, code string, previous byte, index int, path []byte) [][]byte {
	if index == len(code) {
		return [][]byte{path}
	}

	newPaths := [][]byte{}
	paths := Get(cache, previous, code[index])
	for _, p := range paths {
		copyPath := slices.Clone(path)
		copyPath = append(copyPath, p...)
		ps := GetBestMoves(cache, code, code[index], index+1, copyPath)

		newPs := [][]byte{}
		for _, p := range ps {
			if len(newPs) == 0 {
				newPs = append(newPs, p)
				continue
			}
			if len(newPs[0]) == len(p) {
				newPs = append(newPs, p)
				continue
			}
			if len(newPs[0]) > len(p) {
				newPs = [][]byte{p}
			}
		}
		newPaths = append(newPaths, newPs...)
	}

	bestPaths := [][]byte{}
	for _, path := range newPaths {
		if len(bestPaths) == 0 {
			bestPaths = append(bestPaths, path)
			continue
		}
		if len(bestPaths[0]) == len(path) {
			bestPaths = append(bestPaths, path)
			continue
		}
		if len(bestPaths[0]) > len(path) {
			bestPaths = [][]byte{path}
		}
	}

	return bestPaths
}

func GetBestMovesWithMemo(cache map[string][][]byte, code string, previous byte, index int, path []byte) [][]byte {
	if index == len(code) {
		return [][]byte{path}
	}

	newPaths := [][]byte{}
	paths := Get(cache, previous, code[index])
	for _, p := range paths {
		copyPath := slices.Clone(path)
		copyPath = append(copyPath, p...)
		ps := GetBestMovesWithMemo(cache, code, code[index], index+1, copyPath)

		newPs := [][]byte{}
		for _, p := range ps {
			if len(newPs) == 0 {
				newPs = append(newPs, p)
				continue
			}
			if len(newPs[0]) == len(p) {
				newPs = append(newPs, p)
				continue
			}
			if len(newPs[0]) > len(p) {
				newPs = [][]byte{p}
			}
		}
		newPaths = append(newPaths, newPs...)
	}

	bestPaths := [][]byte{}
	for _, path := range newPaths {
		if len(bestPaths) == 0 {
			bestPaths = append(bestPaths, path)
			continue
		}
		if len(bestPaths[0]) == len(path) {
			bestPaths = append(bestPaths, path)
			continue
		}
		if len(bestPaths[0]) > len(path) {
			bestPaths = [][]byte{path}
		}
	}

	return bestPaths
}

func Part2(input string) int {
	sum := 0
	return sum
}
