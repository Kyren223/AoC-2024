package main

import (
	"fmt"
	"log"
	"math"
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

	file, err = os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input = string(file)
	fmt.Println("Part 2:", Part2(input))
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	a64, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(lines[0], ":")[1]), 10, 64)
	b64, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(lines[1], ":")[1]), 10, 64)
	c64, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(lines[2], ":")[1]), 10, 64)

	a := int(a64)
	b := int(b64)
	c := int(c64)

	var code []byte

	programStr := strings.Split(lines[4], " ")[1]
	bytes := strings.Split(programStr, ",")
	for _, bitStr := range bytes {
		bit64, _ := strconv.ParseInt(bitStr, 10, 64)
		bit := byte(bit64)
		code = append(code, bit)
	}
	fmt.Println(a, b, c, code)

	out := VM(code, a, b, c)

	for i, num := range out {
		if i != 0 {
			fmt.Print(",")
		}
		fmt.Print(num)
	}
	fmt.Println()

	sum := 0
	return sum
}

const (
	adv = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func VM(code []byte, a, b, c int) []byte {
	ip := 0
	output := []byte{}
	for 0 <= ip && ip < len(code)-1 {
		opcode := code[ip]
		operand := code[ip+1]
		switch opcode {
		case adv:
			a = a / int(math.Pow(2, float64(Combo(operand, a, b, c))))
		case bxl:
			b ^= int(operand)
		case bst:
			b = Combo(operand, a, b, c) % 8
		case jnz:
			if a != 0 {
				ip = int(operand) - 2
			}
		case bxc:
			b ^= c
		case out:
			combo := byte(Combo(operand, a, b, c) % 8)
			output = append(output, combo)
		case bdv:
			b = a / int(math.Pow(2, float64(Combo(operand, a, b, c))))
		case cdv:
			c = a / int(math.Pow(2, float64(Combo(operand, a, b, c))))
		}
		ip += 2
	}
	return output
}

func Combo(operand byte, a, b, c int) int {
	if operand <= 3 {
		return int(operand)
	}
	if operand == 4 {
		return a
	}
	if operand == 5 {
		return b
	}
	if operand == 6 {
		return c
	}
	panic("invalid")
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	a64, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(lines[0], ":")[1]), 10, 64)
	b64, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(lines[1], ":")[1]), 10, 64)
	c64, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(lines[2], ":")[1]), 10, 64)

	a := int(a64)
	b := int(b64)
	c := int(c64)

	var code []byte

	programStr := strings.Split(lines[4], " ")[1]
	codeStr := strings.Split(programStr, ",")
	for _, bitStr := range codeStr {
		bit64, _ := strconv.ParseInt(bitStr, 10, 64)
		bit := byte(bit64)
		code = append(code, bit)
	}
	fmt.Println(a, b, c, code)

	a = 105817653857465
	out := VM(code, a, b, c)
	fmt.Println(out)

	return 0
}
