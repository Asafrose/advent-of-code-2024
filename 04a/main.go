package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput() ([]string, error) {
	file, err := os.Open("input.txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func getColumns(lines []string) []string {
	var columns []string

	for i := 0; i < len(lines[0]); i++ {
		var column string

		for j := 0; j < len(lines); j++ {
			column += string(lines[j][i])
		}

		columns = append(columns, column)
	}

	return columns
}

func getFirstDiagonals(lines []string) []string {
	var diagonals []string

	for i := 0; i < len(lines); i++ {
		var diagonal string

		for j := 0; j < len(lines); j++ {
			if i+j < len(lines) {
				diagonal += string(lines[i+j][j])
			}
		}

		diagonals = append(diagonals, diagonal)
	}

	for i := 1; i < len(lines); i++ {
		var diagonal string

		for j := 0; j < len(lines); j++ {
			if i+j < len(lines) {
				diagonal += string(lines[j][i+j])
			}
		}

		diagonals = append(diagonals, diagonal)
	}

	return diagonals
}

func getSecondDiagonals(lines []string) []string {
	var diagonals []string

	for i := 0; i < len(lines); i++ {
		var diagonal string

		for j := 0; j < len(lines); j++ {
			if i-j >= 0 {
				diagonal += string(lines[i-j][j])
			}
		}

		diagonals = append(diagonals, diagonal)
	}

	for i := 1; i < len(lines); i++ {
		var diagonal string

		for j := 0; j < len(lines); j++ {
			if i+j < len(lines) {
				diagonal += string(lines[len(lines)-j-1][i+j])
			}
		}

		diagonals = append(diagonals, diagonal)
	}

	return diagonals
}

func main() {
	lines, err := readInput()

	if err != nil {
		panic(err)
	}

	var items []string

	items = append(items, lines...)
	items = append(items, getColumns(lines)...)
	items = append(items, getFirstDiagonals(lines)...)
	items = append(items, getSecondDiagonals(lines)...)

	count := 0

	for _, item := range items {

		count += strings.Count(item, "XMAS")
		count += strings.Count(item, "SAMX")
	}

	fmt.Println(count)
}
