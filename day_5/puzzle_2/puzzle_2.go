package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type ordering_rule struct {
	rules [][2]int
}

func main() {
	file_rules, err := os.Open("../rules.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file_rules.Close()

	file_updates, err := os.Open("../updates.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file_updates.Close()

	scanner_rules := bufio.NewScanner(file_rules)
	rules := new_ordering_rule()

	scanner_updates := bufio.NewScanner(file_updates)
	var updates [][]int
	
	for scanner_rules.Scan() {
		line := strings.TrimSpace(scanner_rules.Text())
		parts := strings.Split(line, "|")

		before, err_before := strconv.Atoi(strings.TrimSpace(parts[0]))
		after, err_after := strconv.Atoi(strings.TrimSpace(parts[1]))
		
		if err_before != nil || err_after != nil {
			fmt.Println("Error parsing rule input", line)
			continue
		}

		rules.add_rule(before, after)

	}

	for scanner_updates.Scan() {
		line := strings.TrimSpace(scanner_updates.Text())
		if line == "" {
			continue
		}

		str_nums := strings.Split(line, ",")
		update := make([]int, len(str_nums))

		for i, str_num := range str_nums {
			num, err := strconv.Atoi(strings.TrimSpace(str_num))
			if err != nil {
				fmt.Println("Error parsing number", str_num, err)
				return
			}
			update[i] = num
		}

		updates = append(updates, update)
	}

	middle_num_sum := 0

	incorrect_slices := incorrect_slices(updates, rules)
	var after_corrected [][]int
	for _, i := range incorrect_slices {
		after_corrected = append(after_corrected, corrected_update(i, rules))
	}

	for _, i := range after_corrected {
		middle_num_sum += i[len(i)/2]
	}
	fmt.Println(middle_num_sum)

}

func new_ordering_rule() *ordering_rule {
	return &ordering_rule{
		rules: [][2]int{},
	}
}

func (or *ordering_rule) add_rule(before, after int) {
	or.rules = append(or.rules, [2]int{before, after})
}

func (or *ordering_rule) check_rules(sequence []int) bool {
	for _, rule := range or.rules {
		before, after := rule[0], rule[1]

		before_index := slices.Index(sequence, before)
		after_index := slices.Index(sequence, after)

		if before_index == -1 || after_index == -1 {
			continue
		}

		if before_index >= after_index {
			return false
		}
	}
	return true
}

func incorrect_slices(slices [][]int, rules *ordering_rule) [][]int {
	var incorrect_slices [][]int

	for _, slice := range slices {
		if !rules.check_rules(slice) {
			incorrect_slices = append(incorrect_slices, slice)
		}
	}

	return incorrect_slices
}

func corrected_update(update []int, rules *ordering_rule) []int {
	corrected_slice := make([]int, len(update))
	copy(corrected_slice, update)

	max_iterations := len(update) * len(update)
	iterations := 0

	for iterations < max_iterations {
		if rules.check_rules(corrected_slice) {
			return corrected_slice
		}

		var violated_rule [2]int
		var before_index, after_index int
		rule_violated := false

		for _, rule := range rules.rules {
			before, after := rule[0],rule[1]

			before_index = -1
			after_index = -1

			for i, num := range corrected_slice {
				if num == before {
					before_index = i
				}
				if num == after {
					after_index = i
				}
			}

			if before_index != -1 && after_index != -1 && before_index >= after_index {
				violated_rule = rule
				rule_violated = true
				break
			}
		}

		if !rule_violated {
			break
		}

		before, after := violated_rule[0], violated_rule[1]

		for i := 0; i < len(corrected_slice); i++ {
			for j := i+1; j < len(corrected_slice); j++ {
				if corrected_slice[i] == after && corrected_slice[j] == before {
					corrected_slice[i], corrected_slice[j] = corrected_slice[j], corrected_slice[i]
					break
				}
			}
		}
		iterations++
	}
	return corrected_slice
}