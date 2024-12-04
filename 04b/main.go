package main

import (
	"bufio"
	"fmt"
	"os"
)

type Window [3][3]rune

func (w Window) String() string {
	return fmt.Sprintf("%c%c%c\n%c%c%c\n%c%c%c\n", w[0][0], w[0][1], w[0][2], w[1][0], w[1][1], w[1][2], w[2][0], w[2][1], w[2][2])
}

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

// get lines of text and returns all posibble 3x3 rune windows
func getRuneWindows(lines []string) []Window {
	var windows []Window

	for i := 0; i < len(lines)-2; i++ {
		for j := 0; j < len(lines[i])-2; j++ {
			var window Window

			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					window[k][l] = rune(lines[i+k][j+l])
				}
			}

			windows = append(windows, window)
		}
	}

	return windows
}

func isValid(window Window) bool {
	return window[1][1] == 'A' &&
		((window[0][0] == 'S' && window[2][2] == 'M') || (window[0][0] == 'M' && window[2][2] == 'S')) &&
		((window[0][2] == 'S' && window[2][0] == 'M') || (window[0][2] == 'M' && window[2][0] == 'S'))
}

func main() {

	lines, err := readInput()

	if err != nil {
		panic(err)
	}

	windows := getRuneWindows(lines)

	count := 0

	for _, window := range windows {

		if isValid(window) {
			count++
		}
	}

	fmt.Println(count)
}
