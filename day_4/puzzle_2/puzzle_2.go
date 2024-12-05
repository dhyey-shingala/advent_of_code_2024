package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	matrix, err := read_file_to_matrix("../input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	rows := len(matrix)
	cols := len(matrix[0])
	count := 0

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if matrix[i][j] == 'A' && check_diagonal(matrix, i, j) {
				// fmt.Printf("Found a match at (%d, %d):\n", i, j)
				// for x := i - 1; x <= i+1; x++ {
				// 	for y := j - 1; y <= j+1; y++ {
				// 		fmt.Printf("%c ", matrix[x][y])
				// 	}
				// 	fmt.Println()
				// }
				// fmt.Println()
				count++
			}
		}
	}
	fmt.Println(count)

}

func read_file_to_matrix(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	var matrix [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := []rune(scanner.Text())
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	return matrix, nil
}

func check_diagonal(matrix [][]rune, i, j int) bool {
	count_m, count_s := 0, 0
	for dx := -1; dx <= 1; dx += 2 {
		for dy := -1; dy <= 1; dy += 2 {
			if matrix[i+dx][j+dy] == 'M' {
				count_m++
			} else if matrix[i+dx][j+dy] == 'S' {
				count_s++
			}
		}
	}
	return count_m == 2 && count_s == 2 && matrix[i-1][j-1] != matrix[i+1][j+1] && matrix[i-1][j+1] != matrix[i+1][j-1]
}
