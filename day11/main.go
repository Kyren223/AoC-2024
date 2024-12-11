package main

import (
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

func Part1(input string) int {
	var stones []int

	for _, stoneStr := range strings.Split(input, " ") {
		stone, err := strconv.ParseInt(stoneStr, 10, 64)
		if err != nil {
			continue
		}
		stones = append(stones, int(stone))
	}

	for i := 0; i < 25; i++ {
		// fmt.Println(stones)
		stones = Blink(stones)
	}
	// fmt.Println(stones)

	return len(stones)
}

func Blink(stones []int) []int {
	var modified []int

	for _, num := range stones {
		s := strconv.FormatInt(int64(num), 10)
		if num == 0 {
			modified = append(modified, 1)
		} else if len(s)%2 == 0 {
			middle := len(s) / 2
			before, _ := strconv.ParseInt(s[:middle], 10, 64)
			after, _ := strconv.ParseInt(s[middle:], 10, 64)
			modified = append(modified, int(before))
			modified = append(modified, int(after))
		} else {
			modified = append(modified, num*2024)
		}
	}

	return modified
}

func Part2(input string) int {
	sum := 0
	return sum
}
