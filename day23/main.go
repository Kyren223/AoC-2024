package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

type Group [3]string

func Part1(input string) int {
	nodes := map[string]map[string]struct{}{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		a := strings.Split(line, "-")[0]
		b := strings.Split(line, "-")[1]

		if m, ok := nodes[a]; ok {
			m[b] = struct{}{}
		} else {
			m := map[string]struct{}{}
			m[b] = struct{}{}
			nodes[a] = m
		}

		if m, ok := nodes[b]; ok {
			m[a] = struct{}{}
		} else {
			m := map[string]struct{}{}
			m[a] = struct{}{}
			nodes[b] = m
		}
	}

	// fmt.Println(nodes)

	groups := map[Group]struct{}{}
	for key, conns := range nodes {
		if key[0] != 't' {
			continue
		}

		for a := range conns {
			aConns := nodes[a]
			for b := range aConns {
				if _, ok := conns[b]; ok {
					group := Group{key, a, b}
					slices.Sort(group[:])
					groups[group] = struct{}{}
				}
			}
		}
	}

	fmt.Println(groups)

	sum := len(groups)
	return sum
}

func Part2(input string) int {
	sum := 0
	return sum
}
