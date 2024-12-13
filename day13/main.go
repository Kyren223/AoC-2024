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

type Machine struct {
	ax int
	ay int
	bx int
	by int
	px int
	py int
}

func Part1(input string) int {
	var machines []Machine
	machinesStr := strings.Split(input, "\n\n")

	for _, machine := range machinesStr {
		lines := strings.Split(machine, "\n")
		if len(lines) < 3 {
			continue
		}

		a := strings.Split(strings.Split(lines[0], ":")[1], ",")
		ax, _ := strconv.ParseInt(strings.TrimSpace(a[0])[1:], 10, 64)
		ay, _ := strconv.ParseInt(strings.TrimSpace(a[1])[1:], 10, 64)

		b := strings.Split(strings.Split(lines[1], ":")[1], ",")
		bx, _ := strconv.ParseInt(strings.TrimSpace(b[0])[1:], 10, 64)
		by, _ := strconv.ParseInt(strings.TrimSpace(b[1])[1:], 10, 64)

		prize := strings.Split(strings.Split(lines[2], ":")[1], ",")
		px, _ := strconv.ParseInt(strings.TrimSpace(prize[0])[2:], 10, 64)
		py, _ := strconv.ParseInt(strings.TrimSpace(prize[1])[2:], 10, 64)

		machines = append(machines, Machine{
			ax: int(ax),
			ay: int(ay),
			bx: int(bx),
			by: int(by),
			px: int(px),
			py: int(py),
		})
	}

	sum := 0
	for _, machine := range machines {
		// fmt.Println(machine)
		minCost := 1000000000000000000 // Infinity
		found := false
		for a := 0; a <= 100; a++ {
			for b := 0; b <= 100; b++ {
				rx := machine.ax*a + machine.bx*b
				ry := machine.ay*a + machine.by*b
				cost := a*3 + b
				if rx == machine.px && ry == machine.py {
					found = true
					minCost = min(minCost, cost)
				}
			}
		}
		if found {
			sum += minCost
		}

	}

	return sum
}

func Part2(input string) int {
	sum := 0
	return sum
}
