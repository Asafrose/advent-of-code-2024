package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	var first, second []int

	file, err := os.Open("input.txt")

	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		firstItem, secondItem, err := parseLine(line)

		if err != nil {
			return nil, nil, err
		}

		first = append(first, firstItem)
		second = append(second, secondItem)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return first, second, nil
}

func calculate(first []int, second []int) int {
	var result int

	for i := 0; i < len(first); i++ {
		result += int(math.Abs(float64(first[i] - second[i])))
	}

	return result
}

func main() {
	first, second, err := parse()

	if err != nil {
		panic(err)
	}

	sort.Ints(first)
	sort.Ints(second)

	result := calculate(first, second)

	fmt.Println(result)
}
