package main

import (
	"fmt"
	"log"
	"math"
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

	// file, err = os.ReadFile("input.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// input = string(file)
	// fmt.Println("Part 2:", Part2(input))
}

type Machine struct {
	ax int64
	ay int64
	bx int64
	by int64
	px int64
	py int64
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
			ax: ax,
			ay: ay,
			bx: bx,
			by: by,
			px: px,
			py: py,
		})
	}

	sum := int64(0)
	for _, machine := range machines {
		// fmt.Println(machine)
		minCost := int64(1000000000000000000) // Infinity
		found := false
		for a := int64(0); a <= 100; a++ {
			for b := int64(0); b <= 100; b++ {
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

	return int(sum)
}

func Part2(input string) int {
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
			ax: ax,
			ay: ay,
			bx: bx,
			by: by,
			px: px + 10000000000000,
			py: py + 10000000000000,
		})
	}

	sum := int64(0)
	for _, machine := range machines {
		fmt.Println(machine)

		set := map[int64]int64{}

		for a := int64(0); a <= max(machine.bx, machine.by); a++ {
			for b := int64(0); b <= max(machine.ax, machine.ay); b++ {
				set[(a*machine.ax)*(b*machine.bx)] = (a * machine.ay) * (b * machine.by)
			}
		}

		orderedSet := []Entry{}
		for xOffset, yOffset := range set {
			if xOffset == 0 && yOffset == 0 {
				continue
			}
			orderedSet = append(orderedSet, Entry{x: xOffset, y: yOffset})
		}
		slices.SortFunc(orderedSet, func(a, b Entry) int {
			if a.x == b.x {
				return 0
			} else if a.x < b.x {
				return 1
			} else {
				return -1
			}
		})

		// fmt.Println(orderedSet)

		valid := false
		rx := machine.px
		ry := machine.py
		fmt.Println("Starting")
		for {
			stop := true
			for _, offset := range orderedSet {
				if rx-offset.x < 0 || ry-offset.y < 0 {
					continue
				}
				if rx-offset.x == 0 && ry-offset.y == 0 {
					valid = true
					stop = true
					break
				}
				rx -= offset.x
				ry -= offset.y
				stop = false
				// fmt.Println("Hit stop")
			}
			// if rx%machine.ax == 0 && (rx/machine.ax)*machine.ay == ry {
			// 	valid = true
			// 	break
			// }
			// if rx%machine.bx == 0 && (rx/machine.bx)*machine.by == ry {
			// 	valid = true
			// 	break
			// }
			// fmt.Println(stop, rx, ry)
			if stop {
				break
			}
		}
		fmt.Println("Stopped")
		if valid {
			fmt.Println("Valid")
			continue
		} else {
			fmt.Println("Invalid")
			continue
		}

		// AX + BY = C
		// [c factors]

		// 5 * 3 = 15

		// 5 5 5 = 15
		// 5 5 3 3 = 16
		// 5 3 3 3 3 = 17
		// 3 3 3 3 3 = 15

		// 23 and 24
		// 24 - 15 = 9 / 3 | 8 N
		// 24 - 16 = 8 N | 7 N
		// 24 - 17 = 7 N | 6 / 3

		asquared := machine.ax*machine.ax + machine.ay*machine.ay
		alen := math.Sqrt(float64(asquared)) / 3

		bsquared := machine.bx*machine.bx + machine.by*machine.by
		blen := math.Sqrt(float64(bsquared))

		// Try to get as much of the one with the highest distance
		aBigger := alen > blen

		xBig := machine.bx
		yBig := machine.by
		xSmall := machine.ax
		ySmall := machine.ay
		if aBigger {
			xBig = machine.ax
			yBig = machine.ay
			xSmall = machine.bx
			ySmall = machine.by
		}

		iters := 0
		maxBig := machine.px / xBig
		fmt.Println("MaxBig:", maxBig)
		found := false
		for maxBig >= 0 {
			xRemain := machine.px - (maxBig * xBig)
			yRemain := machine.py - (maxBig * yBig)
			maxSmall := xRemain / xSmall

			if xSmall*maxSmall == xRemain && ySmall*maxSmall == yRemain {
				found = true
				break
			}

			maxBig--
			iters++

			if iters > 1000000000 {
				break
			}
		}

		if found {
			fmt.Println("found!")
			// sum += minCost
		} else {
			fmt.Println("Not found")
		}
	}

	return int(sum)
}

type Entry struct {
	x int64
	y int64
}
