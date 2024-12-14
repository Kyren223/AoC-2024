package main

import (
	"fmt"
	"io"
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
	fmt.Println("Part 1 Example:", Part1(input, true))

	file, err = os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input = string(file)
	fmt.Println("Part 1:", Part1(input, false))

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

type Vec2 struct {
	x int
	y int
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.x + other.x,
		v.y + other.y,
	}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		v.x - other.x,
		v.y - other.y,
	}
}

func (v Vec2) Mul(other Vec2) Vec2 {
	return Vec2{
		v.x * other.x,
		v.y * other.y,
	}
}

func (v Vec2) Div(other Vec2) Vec2 {
	return Vec2{
		v.x / other.x,
		v.y / other.y,
	}
}

type Robot struct {
	pos      Vec2
	velocity Vec2
}

func Part1(input string, example bool) int {
	var robots []Robot
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		p := strings.Split(line, " ")[0][2:]
		px, _ := strconv.ParseInt(strings.Split(p, ",")[0], 10, 64)
		py, _ := strconv.ParseInt(strings.Split(p, ",")[1], 10, 64)

		v := strings.Split(line, " ")[1][2:]
		vx, _ := strconv.ParseInt(strings.Split(v, ",")[0], 10, 64)
		vy, _ := strconv.ParseInt(strings.Split(v, ",")[1], 10, 64)

		robots = append(robots, Robot{
			pos:      Vec2{int(px), int(py)},
			velocity: Vec2{int(vx), int(vy)},
		})
	}

	width := 101
	height := 103
	if example {
		width = 11
		height = 7
	}
	horizontal := width / 2
	vertical := height / 2

	// fmt.Println("Before:", robots)

	for i := 0; i < 100; i++ {
		// for y := 0; y < height; y++ {
		// 	for x := 0; x < width; x++ {
		// 		if robots[0].pos.x == x && robots[0].pos.y == y {
		// 			fmt.Print("1")
		// 		} else {
		// 			fmt.Print(".")
		// 		}
		// 	}
		// 	fmt.Print("\n")
		// }
		// fmt.Print("\n")
		Step(robots, width, height)
	}
	// for y := 0; y < height; y++ {
	// 	for x := 0; x < width; x++ {
	// 		count := 0
	// 		for _, robot := range robots {
	// 			if robot.pos.x == x && robot.pos.y == y {
	// 				count++
	// 			}
	// 		}
	// 		if count > 0 {
	// 			fmt.Printf("%c", '0'+count)
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }
	//
	// fmt.Println("After:", robots)

	topLeft := 0
	topRight := 0
	bottomLeft := 0
	bottomRight := 0
	for _, robot := range robots {
		if robot.pos.x < horizontal && robot.pos.y < vertical {
			topLeft++
		} else if robot.pos.x < horizontal && robot.pos.y > vertical {
			bottomLeft++
		} else if robot.pos.x > horizontal && robot.pos.y < vertical {
			topRight++
		} else if robot.pos.x > horizontal && robot.pos.y > vertical {
			bottomRight++
		}
	}
	fmt.Println(topLeft, topRight, bottomLeft, bottomRight)

	sum := topLeft * topRight * bottomLeft * bottomRight

	return sum
}

func Step(robots []Robot, width, height int) {
	for i, robot := range robots {
		newPos := robot.pos.Add(robot.velocity)
		if newPos.x < 0 {
			newPos.x = width + newPos.x
		}
		if newPos.y < 0 {
			newPos.y = height + newPos.y
		}
		if newPos.x >= width {
			newPos.x = newPos.x % width
		}
		if newPos.y >= height {
			newPos.y = newPos.y % height
		}
		robot.pos = newPos
		robots[i] = robot
	}
}

func Part2(input string) int {
	var robots []Robot
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		p := strings.Split(line, " ")[0][2:]
		px, _ := strconv.ParseInt(strings.Split(p, ",")[0], 10, 64)
		py, _ := strconv.ParseInt(strings.Split(p, ",")[1], 10, 64)

		v := strings.Split(line, " ")[1][2:]
		vx, _ := strconv.ParseInt(strings.Split(v, ",")[0], 10, 64)
		vy, _ := strconv.ParseInt(strings.Split(v, ",")[1], 10, 64)

		robots = append(robots, Robot{
			pos:      Vec2{int(px), int(py)},
			velocity: Vec2{int(vx), int(vy)},
		})
	}

	width := 101
	height := 103
	horizontal := width / 2
	vertical := height / 2

	// fmt.Println("Before:", robots)

	// file, err := os.OpenFile("picture.txt", os.O_CREATE|os.O_APPEND, 0o644)
	// file, err := os.Create("picture.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	for i := 0; i < 10000; i++ {
		Step(robots, width, height)
		if IsDense(robots, width, height, 4.0) {
			fmt.Println("i:", i)
			Print(os.Stdout, robots, width, height)
		}
		// Print(os.Stdout, robots)
		// Print(file, robots)
	}

	// fmt.Println("After:", robots)

	topLeft := 0
	topRight := 0
	bottomLeft := 0
	bottomRight := 0
	for _, robot := range robots {
		if robot.pos.x < horizontal && robot.pos.y < vertical {
			topLeft++
		} else if robot.pos.x < horizontal && robot.pos.y > vertical {
			bottomLeft++
		} else if robot.pos.x > horizontal && robot.pos.y < vertical {
			topRight++
		} else if robot.pos.x > horizontal && robot.pos.y > vertical {
			bottomRight++
		}
	}
	fmt.Println(topLeft, topRight, bottomLeft, bottomRight)

	sum := topLeft * topRight * bottomLeft * bottomRight

	return sum
}

func Print(writer io.Writer, robots []Robot, width, height int) {
	var err error
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			count := 0
			for _, robot := range robots {
				if robot.pos.x == x && robot.pos.y == y {
					count++
				}
			}
			if count > 0 {
				_, err = fmt.Fprintf(writer, "%c", '0'+count)
				if err != nil {
					panic(err)
				}
			} else {
				_, err = fmt.Fprint(writer, ".")
				if err != nil {
					panic(err)
				}
			}
		}
		_, err = fmt.Fprint(writer, "\n")
		if err != nil {
			panic(err)
		}
	}
	_, err = fmt.Fprint(writer, "\n\n\n\n\n")
	if err != nil {
		panic(err)
	}
}

func IsDense(robots []Robot, width, height int, threshold float64) bool {
	for y := 0; y < height; y++ {
		line := 0
		for x := 0; x < width; x++ {
			hasLine := false
			for _, robot := range robots {
				if robot.pos.x == x && robot.pos.y == y {
					hasLine = true
				}
			}
			if hasLine {
				line++
			} else {
				line = 0
			}
			if line > 30 {
				return true
			}
		}
	}

	// l := int(threshold * float64(len(robots)) / float64(height))
	// for y := 0; y < height; y++ {
	// 	line := 0
	// 	for x := 0; x < width; x++ {
	// 		for _, robot := range robots {
	// 			if robot.pos.x == x && robot.pos.y == y {
	// 				line++
	// 			}
	// 		}
	// 	}
	// 	if line > l {
	// 		return true
	// 	}
	// }

	return false
	//
	// c := int(threshold * float64(len(robots)) / float64(width))
	// for x := 0; x < width; x++ {
	// 	col := 0
	// 	for y := 0; y < height; y++ {
	// 		for _, robot := range robots {
	// 			if robot.pos.x == x && robot.pos.y == y {
	// 				col++
	// 			}
	// 		}
	// 	}
	// 	if col > c {
	// 		return true
	// 	}
	// }
	//
	// return false
}
