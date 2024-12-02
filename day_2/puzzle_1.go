package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func puzzle_1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safe := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line) // Split the line into fields
		num_slice := string_to_int_slice(numbers)

		if (all_increase(num_slice) || all_decrease(num_slice)) && diff_of_at_least_1_or_3(num_slice) {
			safe++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(safe)
}

func string_to_int_slice(numbers []string) []int {
	ints := make([]int, 0, len(numbers))
	for _, s := range numbers {
		num, _ := strconv.Atoi(s)
		ints = append(ints, num)
	}
	return ints
}

func all_increase(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func all_decrease(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func diff_of_at_least_1_or_3(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if 1 <= math.Abs(float64(nums[i]-nums[i+1])) && math.Abs(float64(nums[i]-nums[i+1])) <= 3 {
			continue
		} else {
			return false
		}
	}
	return true
}
