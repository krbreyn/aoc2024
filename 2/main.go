package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func convert_levels(in []string) (out []int) {
	for _, val := range in {
		i, _ := strconv.Atoi(val)
		out = append(out, i)
	}
	return out
}

func is_safe1(levels []int) bool {
	var up bool
	var down bool
	var prev int
	for i, val := range levels {
		if i > 0 {
			prev = levels[i-1]
			diff := val - prev

			if diff == 0 || diff > 3 || diff < -3 {
				return false
			}
			if !up && val > prev {
				up = true
			}
			if !down && prev > val {
				down = true
			}
			if down && up {
				return false
			}
		}
	}
	return true
}

func is_safe2(levels []int) bool {
	var new_levels []int
	for j := range len(levels) {
		new_levels = new_levels[:0]
		for i, val := range levels {
			if j != i {
				new_levels = append(new_levels, val)
			}
		}
		if is_safe1(new_levels) {
			return true
		}
	}
	return false
}

type Target struct {
	first  bool
	second bool
}

func worker(jobs <-chan []int, results chan<- Target, updates chan<- int) {
	for levels := range jobs {
		if is_safe1(levels) {
			results <- Target{
				first:  true,
				second: true,
			}
		} else if is_safe2(levels) {
			results <- Target{
				first:  false,
				second: true,
			}
		}
		updates <- 1
	}
}

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()

	numsafe1 := 0
	numsafe2 := 0

	numWorkers := 8
	results := make(chan Target)
	jobs := make(chan []int)

	counter := 0
	updates := make(chan int)
	done := make(chan bool)
	var wg sync.WaitGroup

	go func() {
		for {
			select {
			case value := <-updates:
				counter += value
				fmt.Printf("\rfinished %d jobs", counter)
			case <-done:
				return
			}
		}
	}()

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, results, updates)
		}()
	}

	go func() {
		for result := range results {
			if result.first && result.second {
				numsafe1 += 1
				numsafe2 += 1
			} else if result.second {
				numsafe2 += 1
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			levels := convert_levels(strings.Fields(scanner.Text()))
			jobs <- levels
		}
		close(jobs)
	}()

	wg.Wait()
	close(results)
	close(done)
	fmt.Println("\ndone!")

	fmt.Println(numsafe1)
	fmt.Println(numsafe2)
}
