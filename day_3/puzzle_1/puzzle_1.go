package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	var matches [][]string

	for scanner.Scan() {
		line := scanner.Text()
		line_matches := re.FindAllStringSubmatch(line, -1)
		if line_matches != nil {
			matches = append(matches, line_matches...)
		}
	}

	sum := 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += mul(num1, num2)
	}
	fmt.Println(sum)
}

func mul(x int, y int) int {
	return x * y
}
