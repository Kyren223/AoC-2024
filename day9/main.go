package main

import (
	"fmt"
	"log"
	"os"
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
	input = input[:len(input)-1]

	var fs []int
	id := 0
	isEmpty := false
	for _, c := range input {
		digit := int(c - '0')
		for i := 0; i < digit; i++ {
			if isEmpty {
				fs = append(fs, -1)
			} else {
				fs = append(fs, id)
			}
		}
		if !isEmpty {
			id++
		}
		isEmpty = !isEmpty
	}

	// fmt.Println(fs)

	// Compress
	for {
		last := fs[len(fs)-1]

		if last == -1 {
			fs = fs[:len(fs)-1]
			continue
		}
		shouldBreak := true
		for i, c := range fs {
			if c == -1 {
				fs[i] = last
				fs = fs[:len(fs)-1]
				shouldBreak = false
				break
			}
		}
		if shouldBreak {
			break
		}
	}

	// fmt.Println(fs)

	sum := 0
	for i, block := range fs {
		// fmt.Println(i, block)
		if block != -1 {
			sum += block * i
		}
	}

	return sum
}

func Part2(input string) int {
	input = input[:len(input)-1]

	var fs []int
	id := 0
	isEmpty := false
	for _, c := range input {
		digit := int(c - '0')
		for i := 0; i < digit; i++ {
			if isEmpty {
				fs = append(fs, -1)
			} else {
				fs = append(fs, id)
			}
		}
		if !isEmpty {
			id++
		}
		isEmpty = !isEmpty
	}

	fmt.Println(fs)

	// Compress
	lastId := fs[len(fs)-1]

	for lastId > 0 {
		length := 0
		start := -1
		for i, block := range fs {
			if block == lastId {
				length++
				if start == -1 {
					start = i
				}
			}
		}

		freeSpace := 0
		spaceStart := -1
		for i, block := range fs {
			if block == -1 {
				freeSpace++
				if spaceStart == -1 {
					spaceStart = i
				}
			} else {
				freeSpace = 0
				spaceStart = -1
			}
			if freeSpace >= length {
				break
			}
		}

		// fmt.Println("ID:", lastId, start, length, spaceStart, freeSpace)

		if freeSpace >= length && start > spaceStart {
			copy(fs[spaceStart:spaceStart+length], fs[start:start+length])
			for i := 0; i < length; i++ {
				fs[start+i] = -1
			}
		}

		// fmt.Println(fs)

		lastId--
	}

	// fmt.Println(fs)

	sum := 0
	for i, block := range fs {
		// fmt.Println(i, block)
		if block != -1 {
			sum += block * i
		}
	}

	return sum
}
