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
	do := true

	pattern := regexp.MustCompile(`(?:(mul)\((\d{1,3}),(\d{1,3})\))|(?:(do)\(\))|(?:(don't)\(\))`)

	matches := pattern.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if match[1] == "mul" {
			if !do {
				continue
			}
			first, err := strconv.Atoi(match[2])

			if err != nil {
				return 0, err
			}

			second, err := strconv.Atoi(match[3])

			if err != nil {
				return 0, err
			}

			result += first * second
		} else if match[4] == "do" {
			do = true
		} else if match[5] == "don't" {
			do = false
		}
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
