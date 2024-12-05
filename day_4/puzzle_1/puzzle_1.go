package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][]int{
	{0, 1},   // Right
	{0, -1},  // Left
	{1, 0},   // Down
	{-1, 0},  // Up
	{1, 1},   // Down Right
	{1, -1},  // Down Left
	{-1, 1},  // Up Right
	{-1, -1}, // Up Left
}

func main() {
	matrix, err := read_file_to_matrix("../input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	word := "XMAS"
	count := count_word_occurrences(matrix, word)
	fmt.Println(count)
}

func read_file_to_matrix(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var matrix [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := []rune(scanner.Text())
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return matrix, nil
}

func count_word_occurrences(matrix [][]rune, word string) int {
	word_runes := []rune(word)
	rows := len(matrix)
	cols := len(matrix[0])
	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range directions {
				if word_found(matrix, word_runes, i, j, dir[0], dir[1]) {
					count++
				}
			}
		}
	}

	return count
}

func word_found(matrix [][]rune, word []rune, start_row, start_col, dir_row, dir_col int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	// Check if first character matches
	if matrix[start_row][start_col] != word[0] {
		return false
	}

	curr_row, curr_col := start_row, start_col

	// Check subsequent characters
	for i := 1; i < len(word); i++ {
		// Move to next position
		curr_row += dir_row
		curr_col += dir_col

		// Check matrix bounds
		if curr_row < 0 || curr_row >= rows || curr_col < 0 || curr_col >= cols {
			return false
		}

		// Check character match
		if matrix[curr_row][curr_col] != word[i] {
			return false
		}
	}

	return true
}