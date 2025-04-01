package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func part1(input []byte) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	grid := make([][]rune, len(lines))
	ans := 0

	dirs := []int{-1, 0, 1}

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	for r := range len(grid) {
		for c := range len(grid[0]) {
			if grid[r][c] != 'X' {
				continue
			}
			for _, dr := range dirs {
				for _, dc := range dirs {
					if dr == 0 && dc == 0 {
						continue
					}
					if !(0 <= r+3*dr && r+3*dr < len(grid) && 0 <= c+3*dc && c+3*dc < len(grid[r])) {
						continue
					}
					if grid[r+dr][c+dc] == 'M' && grid[r+2*dr][c+2*dc] == 'A' && grid[r+3*dr][c+3*dc] == 'S' {
						ans += 1
					}
				}
			}
		}
	}

	return ans
}

func part2(input []byte) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	grid := make([][]rune, len(lines))
	ans := 0

	valid := []string{"MMSS", "MSSM", "SMMS", "SSMM"}
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			if grid[r][c] != 'A' {
				continue
			}
			corners := []rune{
				grid[r-1][c-1],
				grid[r-1][c+1],
				grid[r+1][c+1],
				grid[r+1][c-1],
			}
			for _, strs := range valid {
				if strs == string(corners) {
					ans++
				}
			}
		}
	}
	return ans
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}

	fmt.Println("Day 4 Part 1 ", part1(input))
	fmt.Println("Day 4 Part 2 ", part2(input))
}
