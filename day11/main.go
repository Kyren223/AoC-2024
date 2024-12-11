package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// input := string(file)
	// fmt.Println("Part 1 Example:", Part1(input))

	// file, err := os.ReadFile("input.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// input := string(file)
	// fmt.Println("Part 1:", Part1(input))

	// file, err = os.ReadFile("example.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// input = string(file)
	// fmt.Println("Part 2 Example:", Part2(input))

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input := string(file)
	fmt.Println("Part 2:", Part2(input))
}

func Part1(input string) int {
	var stones []int

	for _, stoneStr := range strings.Split(input, " ") {
		stone, err := strconv.ParseInt(stoneStr, 10, 64)
		if err != nil {
			continue
		}
		stones = append(stones, int(stone))
	}
	fmt.Println(stones)

	// sum := 0
	// for j := 0; j < len(stones); j++ {
	// 	tmpStones := stones[j : j+1]
	// 	for i := 0; i < 75; i++ {
	// 		// fmt.Println(stones)
	// 		tmpStones = Blink(tmpStones)
	// 	}
	// 	// fmt.Println(stones)
	// 	sum += len(tmpStones)
	// }

	stones = stones[0:1]
	for i := 0; i < 75; i++ {
		// fmt.Println(stones)
		stones = Blink(stones)
	}
	fmt.Println(stones)

	return len(stones)
}

func Blink(stones []int) []int {
	var modified []int

	for _, num := range stones {
		if num == 0 {
			modified = append(modified, 1)
		} else {
			s := strconv.FormatInt(int64(num), 10)
			if len(s)%2 == 0 {
				middle := len(s) / 2
				before, _ := strconv.ParseInt(s[:middle], 10, 64)
				after, _ := strconv.ParseInt(s[middle:], 10, 64)
				modified = append(modified, int(before))
				modified = append(modified, int(after))
			} else {
				modified = append(modified, num*2024)
			}
		}
	}

	return modified
}

type Point struct {
	stone int64
	depth int
}

var memo = map[Point]int64{}

func Blink2(stone int64, depth int) int64 {
	mem, ok := memo[Point{stone, depth}]
	if ok {
		return mem
	}

	if depth == 75 {
		return 1
	}

	if stone == 0 {
		result := Blink2(1, depth+1)
		memo[Point{stone, depth}] = result
		return result
	}

	s := strconv.FormatInt(int64(stone), 10)
	if len(s)%2 == 0 {
		middle := len(s) / 2
		before, _ := strconv.ParseInt(s[:middle], 10, 64)
		after, _ := strconv.ParseInt(s[middle:], 10, 64)

		result := Blink2(before, depth+1) + Blink2(after, depth+1)
		memo[Point{stone, depth}] = result
		return result
	}

	result := Blink2(stone*2024, depth+1)
	memo[Point{stone, depth}] = result
	return result
}

func Part2(input string) int {
	var stones []int64

	for _, stoneStr := range strings.Split(input, " ") {
		stone, err := strconv.ParseInt(stoneStr, 10, 64)
		if err != nil {
			continue
		}
		stones = append(stones, stone)
	}
	fmt.Println(stones)

	sum := int64(0)
	for i := 0; i < len(stones); i++ {
		sum += Blink2(stones[i], 0)
	}

	return int(sum)
}
