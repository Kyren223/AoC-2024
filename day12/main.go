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
	plantsMap := strings.Split(input, "\n")
	plantsMap = plantsMap[:len(plantsMap)-1]

	width := len(plantsMap[0])
	height := len(plantsMap)
	marks := make([][]bool, height)
	for y := range marks {
		marks[y] = make([]bool, width)
	}

	sum := 0
	for y, line := range plantsMap {
		for x, c := range line {
			area, permiter := Count(plantsMap, marks, int(c), x, y)
			sum += area * permiter
		}
	}

	return sum
}

func Count(plantsMap []string, marks [][]bool, c, x, y int) (int, int) {
	if marks[y][x] {
		return 0, 0
	}
	marks[y][x] = true

	top := 0
	if y-1 >= 0 {
		top = int(plantsMap[y-1][x])
	}
	bottom := 0
	if y+1 < len(plantsMap) {
		bottom = int(plantsMap[y+1][x])
	}
	left := 0
	if x-1 >= 0 {
		left = int(plantsMap[y][x-1])
	}
	right := 0
	if x+1 < len(plantsMap[0]) {
		right = int(plantsMap[y][x+1])
	}

	area := 1
	perimeter := 0
	if top == c {
		a, p := Count(plantsMap, marks, c, x, y-1)
		area += a
		perimeter += p
	} else {
		perimeter++
	}

	if bottom == c {
		a, p := Count(plantsMap, marks, c, x, y+1)
		area += a
		perimeter += p
	} else {
		perimeter++
	}

	if left == c {
		a, p := Count(plantsMap, marks, c, x-1, y)
		area += a
		perimeter += p
	} else {
		perimeter++
	}

	if right == c {
		a, p := Count(plantsMap, marks, c, x+1, y)
		area += a
		perimeter += p
	} else {
		perimeter++
	}

	return area, perimeter
}

func Part2(input string) int {
	m = map[Pos]Border{}
	plantsMap := strings.Split(input, "\n")
	plantsMap = plantsMap[:len(plantsMap)-1]

	width := len(plantsMap[0])
	height := len(plantsMap)
	marks := make([][]bool, height)
	for y := range marks {
		marks[y] = make([]bool, width)
	}

	sum := 0
	for y, line := range plantsMap {
		for x, c := range line {
			if marks[y][x] {
				continue
			}
			area, perimeter := Count2(plantsMap, marks, int(c), x, y)
			sum += area * perimeter
			fmt.Println(x, y, string(c), area, perimeter, area*perimeter)
		}
	}

	return sum
}

type Pos struct {
	x int
	y int
	c int
}

type Border struct {
	t bool
	b bool
	l bool
	r bool
}

var m = map[Pos]Border{}

func Count2(plantsMap []string, marks [][]bool, c, x, y int) (int, int) {
	if marks[y][x] {
		return 0, 0
	}
	marks[y][x] = true

	top := 0
	if y-1 >= 0 {
		top = int(plantsMap[y-1][x])
	}
	bottom := 0
	if y+1 < len(plantsMap) {
		bottom = int(plantsMap[y+1][x])
	}
	left := 0
	if x-1 >= 0 {
		left = int(plantsMap[y][x-1])
	}
	right := 0
	if x+1 < len(plantsMap[0]) {
		right = int(plantsMap[y][x+1])
	}

	t := m[Pos{x - 1, y, c}].t || m[Pos{x + 1, y, c}].t
	b := m[Pos{x - 1, y, c}].b || m[Pos{x + 1, y, c}].b
	l := m[Pos{x, y - 1, c}].l || m[Pos{x, y + 1, c}].l
	r := m[Pos{x, y - 1, c}].r || m[Pos{x, y + 1, c}].r

	if c == 'E' {
		// fmt.Println("A", x, y, border, t, b, l, r, "|", m[Pos{x - 1, y, c}])
		fmt.Println("A", x, y, t, b, l, r, "|", m[Pos{x - 1, y, c}])
	}

	area := 1
	perimeter := 0

	if top == c {
		a, p := Count2(plantsMap, marks, c, x, y-1)
		area += a
		perimeter += p
	} else if !t {
		perimeter++

		pos := Pos{x: x, y: y, c: c}
		border := m[pos]
		border.t = true
		m[pos] = border

		lx := x - 1
		for lx >= 0 && int(plantsMap[y][lx]) == c {
			if y-1 >= 0 && int(plantsMap[y-1][lx]) == c {
				break
			}
			pos := Pos{lx, y, c}
			border := m[pos]
			border.t = true
			m[pos] = border
			lx--
		}
		rx := x + 1
		for rx < len(plantsMap[0]) && int(plantsMap[y][rx]) == c {
			if y-1 >= 0 && int(plantsMap[y-1][rx]) == c {
				break
			}
			pos := Pos{rx, y, c}
			border := m[pos]
			border.t = true
			m[pos] = border
			rx++
		}

		if c == 'E' {
			fmt.Println(x, y, "p:", perimeter, "top")
		}
	}

	t = m[Pos{x - 1, y, c}].t || m[Pos{x + 1, y, c}].t
	b = m[Pos{x - 1, y, c}].b || m[Pos{x + 1, y, c}].b
	l = m[Pos{x, y - 1, c}].l || m[Pos{x, y + 1, c}].l
	r = m[Pos{x, y - 1, c}].r || m[Pos{x, y + 1, c}].r

	if c == 'E' {
		// fmt.Println("A", x, y, border, t, b, l, r, "|", m[Pos{x - 1, y, c}])
		fmt.Println("A", x, y, t, b, l, r, "|", m[Pos{x - 1, y, c}], "at")
	}

	if bottom == c {
		a, p := Count2(plantsMap, marks, c, x, y+1)
		area += a
		perimeter += p
	} else if !b {
		perimeter++

		pos := Pos{x: x, y: y, c: c}
		border := m[pos]
		border.b = true
		m[pos] = border

		lx := x - 1
		for lx >= 0 && int(plantsMap[y][lx]) == c {
			if y+1 < len(plantsMap) && int(plantsMap[y+1][lx]) == c {
				break
			}
			pos := Pos{lx, y, c}
			border := m[pos]
			border.b = true
			m[pos] = border
			lx--
		}
		rx := x + 1
		for rx < len(plantsMap[0]) && int(plantsMap[y][rx]) == c {
			if y+1 < len(plantsMap) && int(plantsMap[y+1][rx]) == c {
				break
			}
			pos := Pos{rx, y, c}
			border := m[pos]
			border.b = true
			m[pos] = border
			rx++
		}

		if c == 'E' {
			fmt.Println(x, y, "p:", perimeter, "bottom")
		}
	}

	t = m[Pos{x - 1, y, c}].t || m[Pos{x + 1, y, c}].t
	b = m[Pos{x - 1, y, c}].b || m[Pos{x + 1, y, c}].b
	l = m[Pos{x, y - 1, c}].l || m[Pos{x, y + 1, c}].l
	r = m[Pos{x, y - 1, c}].r || m[Pos{x, y + 1, c}].r

	if c == 'E' {
		// fmt.Println("A", x, y, border, t, b, l, r, "|", m[Pos{x - 1, y, c}])
		fmt.Println("A", x, y, t, b, l, r, "|", m[Pos{x - 1, y, c}], "ab")
	}

	if left == c {
		a, p := Count2(plantsMap, marks, c, x-1, y)
		area += a
		perimeter += p
	} else if !l {
		perimeter++

		pos := Pos{x: x, y: y, c: c}
		border := m[pos]
		border.l = true
		m[pos] = border

		ty := y - 1
		for ty >= 0 && int(plantsMap[ty][x]) == c {
			if x-1 >= 0 && int(plantsMap[ty][x-1]) == c {
				break
			}
			pos := Pos{x, ty, c}
			border := m[pos]
			border.l = true
			m[pos] = border
			ty--
		}
		by := y + 1
		for by < len(plantsMap) && int(plantsMap[by][x]) == c {
			if x-1 >= 0 && int(plantsMap[by][x-1]) == c {
				break
			}
			pos := Pos{x, by, c}
			border := m[pos]
			border.l = true
			m[pos] = border
			by++
		}

		if c == 'E' {
			fmt.Println(x, y, "p:", perimeter, "left")
		}
	}

	t = m[Pos{x - 1, y, c}].t || m[Pos{x + 1, y, c}].t
	b = m[Pos{x - 1, y, c}].b || m[Pos{x + 1, y, c}].b
	l = m[Pos{x, y - 1, c}].l || m[Pos{x, y + 1, c}].l
	r = m[Pos{x, y - 1, c}].r || m[Pos{x, y + 1, c}].r

	if c == 'E' {
		// fmt.Println("A", x, y, border, t, b, l, r, "|", m[Pos{x - 1, y, c}])
		fmt.Println("A", x, y, t, b, l, r, "|", m[Pos{x - 1, y, c}], "al")
	}

	if right == c {
		a, p := Count2(plantsMap, marks, c, x+1, y)
		area += a
		perimeter += p
	} else if !r {
		perimeter++

		pos := Pos{x: x, y: y, c: c}
		border := m[pos]
		border.r = true
		m[pos] = border

		ty := y - 1
		for ty >= 0 && int(plantsMap[ty][x]) == c {
			if x+1 < len(plantsMap[0]) && int(plantsMap[ty][x+1]) == c {
				break
			}
			pos := Pos{x, ty, c}
			border := m[pos]
			border.r = true
			m[pos] = border
			ty--
		}
		by := y + 1
		for by < len(plantsMap) && int(plantsMap[by][x]) == c {
			if x+1 < len(plantsMap[0]) && int(plantsMap[by][x+1]) == c {
				break
			}
			pos := Pos{x, by, c}
			border := m[pos]
			border.r = true
			m[pos] = border
			by++
		}

		if c == 'E' {
			fmt.Println(x, y, "p:", perimeter, "right")
		}
	}

	return area, perimeter
}
