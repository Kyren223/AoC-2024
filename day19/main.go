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

func Part1(input string) int {
	towels := strings.Split(strings.Split(input, "\n\n")[0], ", ")
	designs := strings.Split(strings.Split(input, "\n\n")[1], "\n")
	designs = designs[:len(designs)-1]

	// fmt.Println(towels)
	// fmt.Println(designs)

	sum := 0
	for _, design := range designs {
		if IsPossible(towels, design) {
			sum += 1
		}
	}

	return sum
}

func IsPossible(towels []string, design string) bool {
	if design == "" {
		return true
	}

	for _, towel := range towels {
		if !strings.HasPrefix(design, towel) {
			continue
		}
		remainingDesign := design[len(towel):]
		if IsPossible(towels, remainingDesign) {
			return true
		}
	}

	return false
}

func Part2(input string) int {
	towels := strings.Split(strings.Split(input, "\n\n")[0], ", ")
	designs := strings.Split(strings.Split(input, "\n\n")[1], "\n")
	designs = designs[:len(designs)-1]

	// fmt.Println(towels)
	// fmt.Println(designs)

	memo := map[string]int{}
	sum := 0
	for _, design := range designs {
		sum += IsPossible2(towels, design, memo)
	}

	return sum
}

func IsPossible2(towels []string, design string, memo map[string]int) int {
	if design == "" {
		return 1
	}

	sum := 0
	for _, towel := range towels {
		if !strings.HasPrefix(design, towel) {
			continue
		}
		remainingDesign := design[len(towel):]
		if value, ok := memo[remainingDesign]; ok {
			sum += value
		} else {
			value := IsPossible2(towels, remainingDesign, memo)
			memo[remainingDesign] = value
			sum += value
		}
	}

	return sum
}
