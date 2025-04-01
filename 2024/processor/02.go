package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1(input []byte) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	output := 0

	for _, line := range lines {
		parts := strings.Fields(line)
		levels := make([]int, len(parts))

		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			levels[i] = num
		}

		isInc := true
		isDec := true
		isSafe := true

		for i := 0; i < len(levels)-1; i++ {
			diff := levels[i+1] - levels[i]
			abs_diff := math.Abs(float64(diff))
			if diff < 0 {
				isInc = false
			}
			if diff > 0 {
				isDec = false
			}
			if !(1 <= abs_diff && abs_diff <= 3) {
				isSafe = false
			}
		}
		if isSafe && (isInc || isDec) {
			output++
		}
	}

	return output
}

func part2(input []byte) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	output := 0

	for _, line := range lines {
		parts := strings.Fields(line)
		levels := make([]int, len(parts))

		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			levels[i] = num
		}

		isSafe := false
		isSafe = isLevelsSafe(levels)

		if !isSafe {
			for i := range levels {
				modifiedLevels := []int{}
				modifiedLevels = append(modifiedLevels, levels[:i]...)
				modifiedLevels = append(modifiedLevels, levels[i+1:]...)
				isSafe = isLevelsSafe(modifiedLevels)

				if isSafe {
					break
				}
			}
		}

		if isSafe {
			output++
		}
	}

	return output
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}

	part1Ans := part1(input)
	fmt.Println("Output Day 2 Part 1:", part1Ans)
	part2Ans := part2(input)
	fmt.Println("Output Day 2 Part 2:", part2Ans)
}

func isLevelsSafe(levels []int) bool {
	isInc := true
	isDec := true
	isSafe := true

	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		abs_diff := math.Abs(float64(diff))
		if diff < 0 {
			isInc = false
		}
		if diff > 0 {
			isDec = false
		}
		if !(1 <= abs_diff && abs_diff <= 3) {
			isSafe = false
		}
		if !(isSafe && (isInc || isDec)) {
			isSafe = false
		}
	}

	return isSafe
}
