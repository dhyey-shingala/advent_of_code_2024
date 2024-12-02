package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func main() {
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
		num_slice := String_to_int_slice_2(numbers)

		if Error_dampener(num_slice) {
			safe++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(safe)
}

func String_to_int_slice_2(numbers []string) []int {
	ints := make([]int, 0, len(numbers))
	for _, s := range numbers {
		num, _ := strconv.Atoi(s)
		ints = append(ints, num)
	}
	return ints
}

func All_increase_2(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func All_decrease_2(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func Diff_of_at_least_1_or_3_2(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if 1 <= math.Abs(float64(nums[i]-nums[i+1])) && math.Abs(float64(nums[i]-nums[i+1])) <= 3 {
			continue
		} else {
			return false
		}
	}
	return true
}

func Error_dampener(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		// Create a new slice without the i-th element
		removed := append([]int(nil), nums[:i]...)
		removed = append(removed, nums[i+1:]...)

		if Diff_of_at_least_1_or_3_2(removed) && (All_increase_2(removed) || All_decrease_2(removed)) {
			return true
		}
	}
	return false
}
