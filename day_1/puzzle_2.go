package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func count_occurrences(right []int, num int) int {
	count := 0
	for _, n := range right {
		if n == num {
			count++
		}
	}
	return count
}

func puzzle_2() {
	// Opening a file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure file is closed after we are done

	var column1 []int
	var column2 []int

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by spaces
		columns := strings.Fields(line)

		// Convert the string columns to integers and store them
		if len(columns) >= 2 {
			col1, col2 := columns[0], columns[1]

			// Convert string to integer
			var val1, val2 int
			fmt.Sscanf(col1, "%d", &val1)
			fmt.Sscanf(col2, "%d", &val2)

			column1 = append(column1, val1)
			column2 = append(column2, val2)
		}
	}

	// Checking for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	count_of_numbers := make(map[int]int)

	for _, num := range column1 {
		count_of_numbers[num] = count_occurrences(column2, num)
	}

	sum := 0
	for key, value := range count_of_numbers {
		if value != 0 {
			sum += key * value
		}
	}
	fmt.Println(sum)
}
