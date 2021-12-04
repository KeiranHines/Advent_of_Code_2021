package main

import (
	"strconv"
	"strings"
)

func ParseFileString(input []string) ([]string, [][][]string) {
	numString := strings.Split(input[0], ",")
	line := strings.ReplaceAll(input[2], "  ", " ")
	size := len(strings.Split(line, " ")) // Using first line of first board for dimensions
	total := (len(input) - 1) / (size + 1)
	current := 2
	boards := make([][][]string, total)
	for i := 0; i < total; i++ {
		board := make([][]string, size)
		for j, v := range input[current : current+size] {
			board[j] = strings.Split(strings.ReplaceAll(strings.Trim(v, " "), "  ", " "), " ")
		}
		boards[i] = board
		current += size + 1
	}
	return numString, boards
}

// Return boardNum, round of the winner.
func PlayBingo(rounds []string, boards [][][]string) (int, int) {
	for _, round := range rounds {
		for i, board := range boards {
			winner := PlayRound(round, board)
			if winner {
				winningNum, _ := strconv.Atoi(round)
				return i, winningNum
			}
		}
	}
	return -1, 0
}

func PlayRound(round string, board [][]string) bool {
	for x, row := range board {
		for y, num := range row {
			if num == round {
				board[x][y] = ""
				rowWins := CheckRowWins(board[x])
				colWins := CheckColumnWins(board, y)
				return rowWins || colWins
			}
		}
	}
	return false
}

func CheckRowWins(row []string) bool {
	for _, num := range row {
		if num != "" {
			return false
		}
	}
	return true
}

func CheckColumnWins(board [][]string, col int) bool {
	for _, row := range board {
		if row[col] != "" {
			return false
		}
	}
	return true
}

func CalcRemainingNumbers(board [][]string) int {
	count := 0
	for _, row := range board {
		for _, number := range row {
			num, _ := strconv.Atoi(number)
			count += num
		}
	}
	return count
}
