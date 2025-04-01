package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func part1(input []byte) int {
	output := 0

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(string(input), -1)

	for _, match := range matches {

		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		output += num1 * num2
	}

	return output
}

func part2(input []byte) int {
	output := 0

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(string(input), -1)
	do := true

	for _, match := range matches {
		if match[0] == "do()" {
			do = true
			continue
		} else if match[0] == "don't()" {
			do = false
			continue
		} else if !do {
			continue
		}

		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		output += num1 * num2
	}

	return output
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}

	fmt.Print("Day 3 Part 1 ", part1(input))
	fmt.Print("Day 3 Part 2 ", part2(input))
}
