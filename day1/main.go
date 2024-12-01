package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day1")

	file, _ := os.ReadFile("example_input1.txt")
	input := string(file)
	fmt.Println("Part 1 Example:", Part1(input))

	file, _ = os.ReadFile("input1.txt")
	input = string(file)
	fmt.Println("Part 1:", Part1(input))
}

func Part1(input string) int {
	var left []int
	var right []int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		nums := strings.Split(line, "   ")
		leftNum, _ := strconv.ParseInt(nums[0], 10, 64)
		rightNum, _ := strconv.ParseInt(nums[1], 10, 64)

		left = append(left, int(leftNum))
		right = append(right, int(rightNum))
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := range left {
		val := left[i] - right[i]
		if val < 0 {
			val = -val
		}
		sum += val
	}

	return sum
}
