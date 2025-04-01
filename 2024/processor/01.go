package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(input []byte) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var left, right []int

	for _, line := range lines {
		parts := strings.Fields(line)

		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0

	for i := range left {
		distance := left[i] - right[i]
		totalDistance += int(math.Abs(float64(distance)))
	}

	return totalDistance
}

func part2(input []byte) int {
	freq := map[int]int{}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var left, right []int

	for _, line := range lines {
		parts := strings.Fields(line)

		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	for _, n := range right {
		freq[n]++
	}

	sum := 0
	for _, n := range left {
		sum += freq[n] * n
	}

	return sum
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}

	part1_ans := part1(input)
	fmt.Println("Output Day 1 Part 1: ", part1_ans)

	part2_ans := part2(input)
	fmt.Println("Output Day 2 Part 2: ", part2_ans)
}
