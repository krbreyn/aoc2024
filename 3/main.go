package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()

	pattern := `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`
	re, _ := regexp.Compile(pattern)

	scanner := bufio.NewScanner(f)

	var multotal int
	var dototal int
	isDont := false
	for scanner.Scan() {
		text := scanner.Text()
		matches := re.FindAllStringSubmatch(text, -1)

		for _, match := range matches {
			if match[0] == "do()" {
				isDont = false
			} else if match[0] == "don't()" {
				isDont = true
			} else if len(match) == 3 {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				val := x * y
				multotal += val
				if !isDont {
					dototal += val
				}
			}
		}
	}

	fmt.Println(multotal, dototal)
}
