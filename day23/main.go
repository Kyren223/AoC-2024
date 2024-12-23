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

	// fmt.Println(groups)

	sum := len(groups)
	return sum
}

func Part2(input string) int {
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

	// Loop thru all nodes
	// For each node, loop through all it's connections
	// If the current group is empty add the node to the list
	// Otherwise get the connections of the node
	// If all the group nodes exist in the node's connections
	// Add it and move on

	groups := [][]string{{}}
	i := 0
	for key, conns := range nodes {
		for a := range conns {
			if len(groups[i]) == 0 {
				groups[i] = append(groups[i], key, a)
				slices.Sort(groups[i])
				continue
			}

			add := true
			for _, b := range groups[i] {
				if _, ok := nodes[a][b]; !ok {
					add = false
					break
				}
			}
			if add {
				groups[i] = append(groups[i], a)
			} else {
				groups = append(groups, []string{})
				i++
			}
		}
	}

	fmt.Println(groups)

	largest := 0
	index := -1
	for i, group := range groups {
		if len(group) > largest {
			largest = len(group)
			index = i
		}
	}

	fmt.Println(index, largest, groups[index])

	group := groups[index]
	slices.Sort(group)

	for i, s := range group {
		if i != 0 {
			fmt.Print(",")
		}
		fmt.Print(s)
	}
	fmt.Println()

	sum := 0
	return sum
}
