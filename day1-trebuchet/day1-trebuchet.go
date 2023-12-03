package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// func getDataViaHttp() {
// res, err := http.Get("https://adventofcode.com/2023/day/1/input")
// if err != nil {
// 	fmt.Println("Error making the GET request: ", err)
// 	return
// }
// defer res.Body.Close()

// body, err := ioutil.ReadAll(res.Body)
// if err != nil {
// 	fmt.Println("Error reading the response body: ", err)
// 	return
// }
// fmt.Println(string(body))
// }

func process(text string) int {
	var firstDigit, lastDigit rune
	foundFirst := false

	for _, char := range text {
		if unicode.IsDigit(char) {
			if !foundFirst {
				firstDigit = char
				foundFirst = true
			}
			lastDigit = char
		}
	}

	if foundFirst {
		strResult := string([]rune{firstDigit, lastDigit})
		i, err := strconv.Atoi(strResult)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
		}
		return i
	}

	return 0
}

func processPartTwo(text string) int {
	spelledOutNums := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var firstDigit, lastDigit int
	lowestIdx := math.MaxInt32
	highestIdx := 0

	for i := range spelledOutNums {
		idx := strings.Index(text, spelledOutNums[i])
		if idx != -1 && idx < lowestIdx {
			lowestIdx = idx
			firstDigit = i
		}
		if idx != -1 && idx > highestIdx {
			highestIdx = idx
			lastDigit = i
		}
	}

	// TODO: get digits also

	result, err := strconv.Atoi(strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit))
	if err != nil {
		fmt.Println("Error converting to number:", err)
	}

	return result
}

func main() {
	start := time.Now()

	file, err := os.Open("day1-file.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		result += process(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error eading from the file:", err)
	}

	elapsed := time.Since(start)
	fmt.Println("Answer: ", result)
	fmt.Println("Execution time: ", elapsed)
}
