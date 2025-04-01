package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func fetchData(pathNumber int) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("error loading .env file: %v", err)
	}

	sessionToken := os.Getenv("SESSION_TOKEN")
	if sessionToken == "" {
		return "", fmt.Errorf("SESSION_TOKEN not found in .env file")
	}

	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", pathNumber)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: sessionToken,
	}
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	return string(bodyBytes), nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a path number as an argument (e.g., 'go run . 1')")
	}

	pathArg := os.Args[len(os.Args)-1]

	pathNumber, err := strconv.Atoi(pathArg)
	if err != nil {
		log.Fatalf("Invalid path number: %v", err)
	}

	response, err := fetchData(pathNumber)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	fmt.Print(response)
}
