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
		// fmt.Println(secret)
		sum += int(secret)
	}

	return sum
}

func Generate(secret int64) int64 {
	secret ^= secret << 6
	secret = secret % 16777216 // 2^24
	secret ^= secret >> 5
	secret = secret % 16777216 // 2^24
	secret ^= secret << 11
	secret = secret % 16777216 // 2^24

	return secret
}

type Price struct {
	price  int8
	change int8
}

func Part2(input string) int {
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

	monkeys := make([][]Price, len(secrets))
	for j, initialSecret := range secrets {
		prices := make([]Price, 2000)
		secret := initialSecret
		previous := int8(initialSecret % 10)
		for i := 0; i < 2000; i++ {
			secret = Generate(secret)
			secretMod10 := int8(secret % 10)
			prices[i] = Price{secretMod10, secretMod10 - previous}
			previous = secretMod10
		}
		monkeys[j] = prices
	}

	bestBananas := 0
	bestN1 := int8(-10)
	bestN2 := int8(-10)
	bestN3 := int8(-10)
	bestN4 := int8(-10)

	for n1 := int8(-9); n1 < 10; n1++ {
		for n2 := int8(-9); n2 < 10; n2++ {
			for n3 := int8(-9); n3 < 10; n3++ {
				for n4 := int8(-9); n4 < 10; n4++ {
					bananas := 0
					for _, monk := range monkeys {
						for j := 0; j < 2000-4; j++ {
							cmp := n1 ^ monk[j].change
							cmp |= n2 ^ monk[j+1].change
							cmp |= n3 ^ monk[j+2].change
							cmp |= n4 ^ monk[j+3].change
							if cmp == 0 {
								bananas += int(monk[j+3].price)
								break
							}
						}
					}
					if bananas > bestBananas {
						bestBananas = bananas
						bestN1 = n1
						bestN2 = n2
						bestN3 = n3
						bestN4 = n4
					}
				}
			}
			fmt.Println(n1, n2)
		}
	}

	fmt.Println(bestBananas, bestN1, bestN2, bestN3, bestN4)
	// Got it

	sum := 0
	return sum
}
