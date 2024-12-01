package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main()	{
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
	for scanner.Scan()	{
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
	if err := scanner.Err(); err != nil	{
		fmt.Println("Error reading file:", err)
	}

	// Sorting both slices
	sort.Ints(column1)
	sort.Ints(column2)

	// fmt.Println("Column 1:", column1)
	// fmt.Println("Column 2:", column2)

	sum := 0
	for i := 0; i < len(column1); i++ {
		sum += int(math.Abs(float64(column1[i] - column2[i])))
	}

	fmt.Println(sum)
}