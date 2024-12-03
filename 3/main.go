package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type NumPair struct {
	x      int
	y      int
	isDont bool
}

func main() {
	f, _ := os.Open("./bigboy.txt")
	defer f.Close()

	pattern := `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`
	re, _ := regexp.Compile(pattern)

	scanner := bufio.NewScanner(f)

	var pairs []NumPair
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
				pairs = append(pairs, NumPair{x: x, y: y, isDont: isDont})
			}
		}
	}

	var multotal int
	var dototal int
	for _, pair := range pairs {
		temp := pair.x * pair.y
		multotal += temp
		if !pair.isDont {
			dototal += temp
		}
	}
	fmt.Println(multotal, dototal)
}
