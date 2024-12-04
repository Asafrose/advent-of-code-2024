package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func count(content string) (int, error) {
	result := 0

	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := pattern.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		first, err := strconv.Atoi(match[1])

		if err != nil {
			return 0, err
		}

		second, err := strconv.Atoi(match[2])

		if err != nil {
			return 0, err
		}

		result += first * second
	}

	return result, nil
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	text, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	result, err := count(string(text))

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
