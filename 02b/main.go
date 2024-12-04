package main

import (
	"bufio"
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

func isPairSafe(first, second int, isAcending bool) bool {
	return first != second && math.Abs(float64(first-second)) <= 3 && isAcending == (first < second)
}

func isSafe(items []int) bool {
	if len(items) == 1 {
		return true
	}

	isAcending := items[0] < items[1]

	if !isPairSafe(items[0], items[1], isAcending) {
		return false
	}

	for i := 2; i < len(items); i++ {
		if !isPairSafe(items[i-1], items[i], isAcending) {
			return false
		}
	}

	return true
}

func isSafeCombinations(items []int) bool {
	fmt.Println("Running for", items)
	if isSafe(items) {
		return true
	}

	for i := 0; i < len(items); i++ {

		var x []int

		x = append(x, items[:i]...)
		x = append(x, items[i+1:]...)

		fmt.Println("Reiterating! Running for", items[:i], items[i+1:], x)
		if isSafe(x) {
			return true
		}
	}

	return false
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

		if isSafeCombinations(items) {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(count)
}
