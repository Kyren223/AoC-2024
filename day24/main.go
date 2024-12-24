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

type Gate struct {
	a   string
	op  string
	b   string
	out string
}

func Part1(input string) int {
	consts := strings.Split(input, "\n\n")[0]
	var gates []Gate
	gatesStr := strings.Split(input, "\n\n")[1]

	variables := map[string]byte{}

	for _, constant := range strings.Split(consts, "\n") {
		if constant == "" {
			continue
		}
		s := strings.Split(constant, ": ")[0]
		n := strings.Split(constant, ": ")[1]
		bit := byte(0)
		if n == "1" {
			bit = 1
		}
		variables[s] = bit
	}

	for _, g := range strings.Split(gatesStr, "\n") {
		if g == "" {
			continue
		}
		args := strings.Split(g, " -> ")[0]
		output := strings.Split(g, " -> ")[1]

		a := strings.Split(args, " ")[0]
		op := strings.Split(args, " ")[1]
		b := strings.Split(args, " ")[2]

		gates = append(gates, Gate{
			a:   a,
			op:  op,
			b:   b,
			out: output,
		})
	}

	z := int64(0)
	for i, gate := range gates {
		if gate.out[0] != 'z' {
			continue
		}
		bit := Compute(gates, variables, i)
		n, _ := strconv.ParseInt(gate.out[1:], 10, 64)
		z |= int64(bit) << n
		fmt.Println("Z", bit, n)
	}

	// fmt.Println(variables)
	// fmt.Println(gates)
	fmt.Println(z)

	sum := int(z)
	return sum
}

func Compute(gates []Gate, variables map[string]byte, index int) byte {
	gate := gates[index]
	if bit, ok := variables[gate.out]; ok {
		return bit
	}

	a := byte(2) // Invalid
	b := byte(2) // invalid

	if bit, ok := variables[gate.a]; ok {
		a = bit
	}
	if bit, ok := variables[gate.b]; ok {
		b = bit
	}

	for i := range gates {
		if a == 2 && gates[i].out == gate.a {
			bit := Compute(gates, variables, i)
			a = bit
			variables[gate.a] = bit
		}
		if b == 2 && gates[i].out == gate.b {
			bit := Compute(gates, variables, i)
			b = bit
			variables[gate.b] = bit
		}
	}

	// fmt.Println(gate, a, b)

	switch gate.op {
	case "OR":
		return a | b
	case "AND":
		return a & b
	case "XOR":
		return a ^ b
	default:
		panic("unknown")
	}
}

func Part2(input string) int {
	sum := 0
	return sum
}
