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
	fmt.Println("Part 2 Example:", Part2(input))

	file, err = os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input = string(file)
	fmt.Println("Part 2:", Part2(input))
}

type Test struct {
	result   int
	operands []int
}

func Part2(input string) int {
	var tests []Test
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		split := strings.Split(line, ":")
		result, _ := strconv.ParseInt(split[0], 10, 64)
		strOperands := strings.Split(split[1], " ")
		var operands []int
		for _, strOperand := range strOperands {
			operand, err := strconv.ParseInt(strOperand, 10, 64)
			if err != nil {
				continue
			}
			operands = append(operands, int(operand))
		}

		tests = append(tests, Test{
			result:   int(result),
			operands: operands,
		})
	}

	// fmt.Println(tests)

	sum := 0
	for _, test := range tests {
		if IsValidTest(test) {
			sum += test.result
		}
	}

	return sum
}

func IsValidTest(test Test) bool {
	return SumOperands(test, 0, 0, 0)
}

func SumOperands(test Test, i, op, sum int) bool {
	if len(test.operands) == i {
		return sum == test.result
	}

	switch op {
	case 0:
		sum += test.operands[i]
	case 1:
		sum *= test.operands[i]
	case 2:
		sumStr := strconv.Itoa(sum)
		operandStr := strconv.Itoa(test.operands[i])
		resultStr := sumStr + operandStr
		result, _ := strconv.ParseInt(resultStr, 10, 64)
		sum = int(result)
	}

	return SumOperands(test, i+1, 0, sum) || SumOperands(test, i+1, 1, sum) || SumOperands(test, i+1, 2, sum)
}
