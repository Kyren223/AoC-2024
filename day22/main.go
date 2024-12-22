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

func Part1(input string) int {
	var secrets []int64
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		secret, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic("Huh")
		}
		secrets = append(secrets, secret)
	}

	sum := 0
	for _, initialSecret := range secrets {
		secret := initialSecret
		for i := 0; i < 2000; i++ {
			secret = Generate(secret)
		}
		fmt.Println(secret)
		sum += int(secret)
	}

	return sum
}

func Generate(secret int64) int64 {
	secret ^= secret<<6
	secret = secret % 16777216 // 2^24
	secret ^= secret>>5
	secret = secret % 16777216 // 2^24
	secret ^= secret<<11
	secret = secret % 16777216 // 2^24

	return secret
}

func Part2(input string) int {
	sum := 0
	return sum
}
