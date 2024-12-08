package main

import (
	"bufio"
	"fmt"
	"os"
)

type guardPos struct {
	x int
	y int
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var board []string
	var lineNum int
	var pos guardPos
	var stopCheck bool

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if !stopCheck {
			for i, char := range line {
				if char == '^' {
					stopCheck = true
					pos.x = lineNum
					pos.y = i
				}
			}
		}
		board = append(board, line)
		lineNum++
	}

	fmt.Println(silver(pos, board))

	//brute force all gold possibilities
	var gold int
	for x, line := range board {
		for y, char := range line {
			if char != '#' && char != '^' {
				obstructionPos := guardPos{x, y}
				didPass := isGold(pos, board, obstructionPos)
				// fmt.Println(obstructionPos, didPass)
				if didPass {
					gold++
				}
			}
		}
	}
	fmt.Println(gold)

	//the heuristic is simply checking if its touched hashes its already touched before
	//a certain amount of times

}

func silver(pos guardPos, board []string) int {
	var hasLeftBoard bool
	var hasHitHash bool
	hasVisited := make(map[guardPos]bool)
	dir := "up"

	silver := 1
	for !hasLeftBoard {
		for !hasHitHash {
			check := peek(pos, dir)
			if check.x >= len(board) || check.y >= len(board[check.x]) {
				hasLeftBoard = true
				break
			}
			if board[check.x][check.y] == '#' {
				hasHitHash = true
				break
			}
			if !hasVisited[pos] {
				hasVisited[pos] = true
				silver++
			}
			pos = check
		}
		if hasLeftBoard {
			break
		}
		if hasHitHash {
			hasHitHash = false
			dir = changeDir(dir)
		}
	}

	return silver
}

func isGold(pos guardPos, board []string, obstructionPos guardPos) bool {
	hasTouchedHash := make(map[guardPos]bool)
	var hasHitHash bool
	var hasLeftBoard bool
	var isInLoop bool

	var loopHeuristic int
	loopMax := 100

	dir := "up"
	for !hasLeftBoard || !isInLoop {
		for !hasHitHash {
			check := peek(pos, dir)
			if check.x >= len(board) || check.x < 0 ||
				check.y >= len(board[check.x]) || check.y < 0 {
				return false //left board
			}
			if board[check.x][check.y] == '#' || check == obstructionPos {
				hasHitHash = true
				if !hasTouchedHash[check] {
					hasTouchedHash[check] = true
				} else {
					loopHeuristic++
					if loopHeuristic > loopMax {
						return true
					}
				}
				break
			}
			pos = check
		}
		if hasHitHash {
			hasHitHash = false
			dir = changeDir(dir)
		}
	}

	//should never be reached?
	return false
}

func changeDir(cur string) string {
	var dir string
	switch cur {
	case "up":
		dir = "right"
	case "right":
		dir = "down"
	case "down":
		dir = "left"
	case "left":
		dir = "up"
	}
	return dir
}

func peek(pos guardPos, dir string) guardPos {
	switch dir {
	case "up":
		return guardPos{
			x: pos.x - 1,
			y: pos.y,
		}
	case "right":
		return guardPos{
			x: pos.x,
			y: pos.y + 1,
		}
	case "down":
		return guardPos{
			x: pos.x + 1,
			y: pos.y,
		}
	case "left":
		return guardPos{
			x: pos.x,
			y: pos.y - 1,
		}
	default:
		panic("malformed peek dir")
	}
}
