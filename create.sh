#!/bin/bash

# Check if a parameter is passed
if [ -z "$1" ]; then
    echo "Usage: $0 <task>"
    exit 1
fi

# Create a directory with the name of the first parameter
mkdir "$1"

# Change to the new directory
cd "$1" || exit

# Initialize a Go module
go mod init github.com/asafrose/advent-of-code-2024/"$1"

# Create a main.go file
cat <<EOL > main.go
package main

import "fmt"

func main() {
        fmt.Println("Hello, Advent of Code 2024!")
}
EOL