package main

import (
	"strconv"
	"strings"
)

// Parses a bingo file to a list of numbers to be drawn and the list
// of active game boards
// Returns the list of numbers to draw in order, the game boards
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

// PlayBingo Plays bingo until a winner is found. If multiple winners are found on the same
// round all the winners are removed and the list of new player boards is returned.
// Returns the boards remaining in play, the winning board, and the remaining round numbers
func PlayBingo(rounds []string, boards [][][]string) ([][][]string, [][]string, []string) {
	for r, round := range rounds {
		var winner [][]string = nil
		for i := len(boards) - 1; i >= 0; i-- {
			board := boards[i]
			w := PlayRound(round, board)
			if w {
				winner = board
				boards = KickPlayer(i, boards)
			}
		}
		if winner != nil {
			return boards, winner, rounds[r:]
		}
	}
	return make([][][]string, 0), make([][]string, 0), make([]string, 0)
}

// PlayRound plays a round of bingo for the give game board.
// Marks the played number as and empty string if it has been found.
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

// CheckRowWins Checks if all the numbers in the given row have been marked (empty string).
func CheckRowWins(row []string) bool {
	for _, num := range row {
		if num != "" {
			return false
		}
	}
	return true
}

// CheckColumnWins Checks if all the numbers in the given col have been marked (empty string).
// Returns true if they are all marked.
func CheckColumnWins(board [][]string, col int) bool {
	for _, row := range board {
		if row[col] != "" {
			return false
		}
	}
	return true
}

// CalcRemainingNumbers
// Returns the sum of all unmarked numbers on the players board.
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

// KickPlayer Removes a player at the given index from the active boards.
// Returns the new list of active players.
func KickPlayer(board int, boards [][][]string) [][][]string {
	newPlayerCount := len(boards) - 1
	boards[board] = boards[newPlayerCount]
	return boards[:newPlayerCount]
}
