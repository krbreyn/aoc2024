package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()

	var board [][]rune

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		row := scanner.Text()
		board = append(board, []rune(row))
	}

	var silver int
	var gold int
	for row, line := range board {
		for col, char := range line {
			if char == 'X' {
				//forward
				if col+3 < len(board) {
					thisrow := board[row]
					if thisrow[col+1] == 'M' &&
						thisrow[col+2] == 'A' &&
						thisrow[col+3] == 'S' {
						silver++
					}
				}
				//backward
				if col >= 3 {
					if line[col-1] == 'M' &&
						line[col-2] == 'A' &&
						line[col-3] == 'S' {
						silver++
					}
				}
				//up
				if row >= 3 {
					if board[row-1][col] == 'M' &&
						board[row-2][col] == 'A' &&
						board[row-3][col] == 'S' {
						silver++
					}
				}
				//down
				if row+3 < len(board) {
					if board[row+1][col] == 'M' &&
						board[row+2][col] == 'A' &&
						board[row+3][col] == 'S' {
						silver++
					}
				}
				//up-back
				if row >= 3 && col >= 3 {
					if board[row-1][col-1] == 'M' &&
						board[row-2][col-2] == 'A' &&
						board[row-3][col-3] == 'S' {
						silver++
					}
				}
				//up-forward
				if row >= 3 && col+3 < len(board[row]) {
					if board[row-1][col+1] == 'M' &&
						board[row-2][col+2] == 'A' &&
						board[row-3][col+3] == 'S' {
						silver++
					}
				}
				//down-back
				if row+3 < len(board) && col >= 3 {
					if board[row+1][col-1] == 'M' &&
						board[row+2][col-2] == 'A' &&
						board[row+3][col-3] == 'S' {
						silver++
					}
				}
				//down-forward
				if row+3 < len(board) && col+3 <= len(board)-1 {
					if board[row+1][col+1] == 'M' &&
						board[row+2][col+2] == 'A' &&
						board[row+3][col+3] == 'S' {
						silver++
					}
				}
			}
			if char == 'A' {
				if row > 0 && col > 0 &&
					row < len(board)-1 && col < len(board[row])-1 {
					// var mcount int
					// var scount int
					corners := []rune{
						board[row-1][col-1],
						board[row+1][col+1],
						board[row+1][col-1],
						board[row-1][col+1],
					}

					if slices.Equal(corners, []rune{'M', 'S', 'M', 'S'}) ||
						slices.Equal(corners, []rune{'M', 'S', 'S', 'M'}) ||
						slices.Equal(corners, []rune{'S', 'M', 'M', 'S'}) ||
						slices.Equal(corners, []rune{'S', 'M', 'S', 'M'}) {
						gold++
					}
				}
			}
		}
	}

	fmt.Println(silver)
	fmt.Println(gold)
}
