package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) (int, int, error) {
	items := strings.Split(line, "   ")

	if len(items) != 2 {
		return 0, 0, fmt.Errorf("Invalid input line: %s", line)
	}

	firstItem, err := strconv.Atoi(items[0])

	if err != nil {
		return 0, 0, err
	}

	secondItem, err := strconv.Atoi(items[1])

	if err != nil {
		return 0, 0, err
	}

	return firstItem, secondItem, nil
}

func parse() ([]int, []int, error) {
	var left, right []int

	file, err := os.Open("input.txt")

	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		leftItem, rightItem, err := parseLine(line)

		if err != nil {
			return nil, nil, err
		}

		left = append(left, leftItem)
		right = append(right, rightItem)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil
}

func calculate(left []int, right []int) int {
	rightItemToCount := make(map[int]int)

	for _, item := range right {
		rightItemToCount[item]++
	}

	var result int

	for _, item := range left {
		result += item * rightItemToCount[item]
	}

	return result
}

func main() {
	first, second, err := parse()

	if err != nil {
		panic(err)
	}

	result := calculate(first, second)

	fmt.Println(result)
}
