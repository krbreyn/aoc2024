package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var leftvals []int
	var rightvals []int

	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	seen_map := make(map[int]bool)
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		l, lerr := strconv.Atoi(split[0])
		r, rerr := strconv.Atoi(split[1])
		if lerr != nil || rerr != nil {
			panic("split failed")
		}
		leftvals = append(leftvals, l)
		rightvals = append(rightvals, r)
		if !seen_map[l] {
			seen_map[l] = true
		}
	}

	slices.Sort(leftvals)
	slices.Sort(rightvals)

	var total_dist int
	var sim_score int
	for i, lval := range leftvals {
		rval := rightvals[i]
		total_dist += max(lval, rval) - min(lval, rval)

		if seen_map[rval] {
			sim_score += rval
		}
	}

	fmt.Println(total_dist)
	fmt.Println(sim_score)
}
