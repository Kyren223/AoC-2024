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

func Part1(input string) int {
	antenasMap := strings.Split(input, "\n")
	antenasMap = antenasMap[:len(antenasMap)-1]

	// fmt.Println("Width:", len(antenasMap[0]), "Height:", len(antenasMap))
	width := len(antenasMap[0])
	height := len(antenasMap)

	marks := make([][]bool, height)
	for y := range marks {
		marks[y] = make([]bool, width)
	}

	for y, line := range antenasMap {
		for x, c := range line {
			if ('0' > c || c > '9') && ('a' > c || c > 'z') && ('A' > c || c > 'Z') {
				continue
			}
			fmt.Println(x, y, c)

			for y2, line2 := range antenasMap {
				for x2, c2 := range line2 {
					if c != c2 || (x == x2 && y == y2) {
						continue
					}

					xDiff := x2 - x
					yDiff := y2 - y

					xPoint1 := x2 + xDiff
					yPoint1 := y2 + yDiff
					if 0 <= xPoint1 && xPoint1 < width && 0 <= yPoint1 && yPoint1 < height {
						marks[yPoint1][xPoint1] = true
					}

					xPoint2 := x - xDiff
					yPoint2 := y - yDiff
					if 0 <= xPoint2 && xPoint2 < width && 0 <= yPoint2 && yPoint2 < height {
						marks[yPoint2][xPoint2] = true
					}

				}
			}

		}
	}

	sum := 0
	for _, mark := range marks {
		for _, marked := range mark {
			if marked {
				sum++
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

	return sum
}

func Part2(input string) int {
	sum := 0
	return sum
}
