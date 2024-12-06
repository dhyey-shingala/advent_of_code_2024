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
	for _, update := range updates {
		fmt.Printf("Sequence %v: %v\n", update, rules.check_rules(update))
		if rules.check_rules(update) {
			if len(update) % 2 == 1 {
				middle_num_sum += update[len(update)/2] // For Odd, should always be odd? idk question is vague
			} else {
				fmt.Println("What are the elves smoking on??")
			}
		}
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
