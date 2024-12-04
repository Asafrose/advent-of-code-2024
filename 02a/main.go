package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) ([]int, error) {
	var items []int

	for _, item := range strings.Split(line, " ") {
		item, err := strconv.Atoi(item)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func isSafe(items []int) error {
	if len(items) == 1 {
		return nil
	}

	if items[0] == items[1] {
		return errors.New("Change between item0 and item1 is 0")
	}

	if math.Abs(float64(items[0]-items[1])) > 3 {
		return errors.New("Change between item0 and item1 is greater than 3")
	}

	isAcending := items[0] < items[1]

	for i := 2; i < len(items); i++ {
		if isAcending != (items[i-1] < items[i]) {
			if isAcending {
				return errors.New(fmt.Sprintf("Change from acending to decending in indexes %d and %d", i-1, i))
			} else {
				return errors.New(fmt.Sprintf("Change from decending to acending in indexes %d and %d", i-1, i))
			}
		}

		if items[i-1] == items[i] {
			return errors.New(fmt.Sprintf("Change between item%d and item%d is 0", i-1, i))
		}

		if math.Abs(float64(items[i-1]-items[i])) > 3 {
			return errors.New(fmt.Sprintf("Change between item%d and item%d is greater than 3", i-1, i))
		}
	}

	return nil
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	count := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		items, err := parseLine(line)

		if err != nil {
			panic(err)
		}

		x := isSafe(items)
		if x == nil {
			count++
		}

		fmt.Println(line, "|||", x == nil, x)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(count)
}
