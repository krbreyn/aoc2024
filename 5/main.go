package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"math/rand"
)

type beforesMap map[string][]string

func (bm beforesMap) split_line(line string) {
	split := strings.Split(line, "|")
	bm[split[1]] = append(bm[split[1]], split[0])
}

func (bm beforesMap) process_splits(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else {
			bm.split_line(line)
		}
	}
}

func (bm beforesMap) process_reports(scanner *bufio.Scanner) ([][]string, [][]string) {
	var valid_reports [][]string
	var invalid_reports [][]string

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")

		did_fail := bm.validate_report(fields)

		if !did_fail {
			valid_reports = append(valid_reports, fields)
		} else {
			invalid_reports = append(invalid_reports, fields)
		}
	}

	return valid_reports, invalid_reports
}

func (bm beforesMap) validate_report(fields []string) bool {

	encountered := make(map[string]bool)
	var failed bool

	for _, num := range fields {
		if !encountered[num] {
			encountered[num] = true
		}
		for _, before := range bm[num] {
			if slices.Contains(fields, before) && !encountered[before] {
				failed = true
			}
		}
	}

	return failed
}

func (bm beforesMap) brute_force_updates(report []string) []string {
	for bm.validate_report(report) {
		rand.Shuffle(len(report), func(i, j int) {
			report[i], report[j] = report[j], report[i]
		})
	}
	return report
}

func (bm beforesMap) order(report, new_report []string) ([]string, []string) {
	var left_overs []string

	for _, val := range report {
		canPlace := true
		for _, before := range bm[val] {
			if slices.Contains(report, before) && !slices.Contains(new_report, before) {
				canPlace = false
			}
		}
		if canPlace {
			new_report = append(new_report, val)
		} else {
			left_overs = append(left_overs, val)
		}
	}

	return new_report, left_overs
}

func (bm beforesMap) order_updates(report []string) []string {
	new_report := make([]string, 0, len(report))

	new_report, left_overs := bm.order(report, new_report)
	for len(new_report) != len(report) {
		new_report, left_overs = bm.order(left_overs, new_report)
	}

	return new_report
}

func count_middles(reports [][]string) int {
	var count int

	for _, report := range reports {
		middle := len(report) / 2
		for i, val := range report {
			if i == middle {
				temp, _ := strconv.Atoi(val)
				count += temp
			}
		}
	}
	return count
}

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)
	befores_map := make(beforesMap)

	befores_map.process_splits(scanner)
	valid_reports, invalid_reports := befores_map.process_reports(scanner)

	newly_valid_reports := make([][]string, 0, len(invalid_reports))
	for _, report := range invalid_reports {
		newly_valid_reports = append(newly_valid_reports, befores_map.order_updates(report))
	}

	silver := count_middles(valid_reports)
	gold := count_middles(newly_valid_reports)
	fmt.Println(silver, gold)
}
