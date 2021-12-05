package main

import (
	"strconv"
	"testing"
)

func TestDay4BingoParser(t *testing.T) {
	input := FileToStringArray("inputs/d4t", "\n")
	rounds, boards := ParseFileString(input)
	if boards[0][0][0] != "22" {
		t.Fatal("Expected board 0 postition 0,0 to be 22")
	}

	if rounds[0] != "7" {
		t.Fatal("First round should have drawn 0")
	}

	if boards[2][4][4] != "7" {
		t.Fatalf("Expected board 3 postition 4,4 (the last number) to be 7 not %v", boards[2][4][4])
	}
}

func TestPlayingRoundUpdatesTheBoard(t *testing.T) {
	board := [][]string{{"10", "11", "12"}, {"1", "2", "3"}, {"21", "22", "23"}}
	PlayRound("21", board)
}

func TestDay4ExampleWinner(t *testing.T) {
	input := FileToStringArray("inputs/d4t", "\n")
	rounds, boards := ParseFileString(input)
	_, winner, round := PlayBingo(rounds, boards)
	if round[0] != "24" {
		t.Fatal("The winning number for player 3 should be 24 not", round)
	}

	count := CalcRemainingNumbers(winner)
	if count != 188 {
		t.Fatal("Incorrect sum of remaining numbers, expected 188, got ", count)
	}
	winningNum, _ := strconv.Atoi(round[0])
	result := count * winningNum
	if result != 4512 {
		t.Fatal("One of the other tests should have failed.")
	}
}

func TestDay4Part1Input(t *testing.T) {
	ans := 32844
	input := FileToStringArray("inputs/d4a", "\n")
	rounds, boards := ParseFileString(input)
	_, winner, round := PlayBingo(rounds, boards)
	count := CalcRemainingNumbers(winner)
	winningNum, _ := strconv.Atoi(round[0])
	result := count * winningNum
	if result != ans {
		t.Fatal("Expected", ans, ", got", result)
	}
}

func TestDay4ExampleLoser(t *testing.T) {
	input := FileToStringArray("inputs/d4t", "\n")
	rounds, boards := ParseFileString(input)
	for ok := true; ok; ok = len(boards) > 1 {
		newBoards, _, roundLeft := PlayBingo(rounds, boards)
		boards = newBoards
		rounds = roundLeft[1:]
	}
	if len(boards) != 1 {
		t.Fatal("Last board not found there is ", len(boards), "in play")
	}
	_, board, roundLeft := PlayBingo(rounds, boards)

	if roundLeft[0] != "13" {
		t.Fatal("The final number should be 13 not", roundLeft[0])
	}

	count := CalcRemainingNumbers(board)
	if count != 148 {
		t.Fatal("Incorrect sum of remaining numbers, expected 148, got ", count)
	}
	finalNum, _ := strconv.Atoi(roundLeft[0])
	result := count * finalNum
	if result != 1924 {
		t.Fatal("One of the other tests should have failed.")
	}
}

func TestDay4Part2Input(t *testing.T) {
	ans := 4920
	input := FileToStringArray("inputs/d4a", "\n")
	rounds, boards := ParseFileString(input)
	for ok := true; ok; ok = len(boards) > 1 {
		newBoards, _, roundLeft := PlayBingo(rounds, boards)
		boards = newBoards
		rounds = roundLeft[1:]
	}
	if len(boards) != 1 {
		t.Fatal("Last board not found there is ", len(boards), "in play")
	}
	_, board, roundLeft := PlayBingo(rounds, boards)
	count := CalcRemainingNumbers(board)
	finalNum, _ := strconv.Atoi(roundLeft[0])
	result := count * finalNum
	if result != ans {
		t.Fatal("Expected", ans, ", got", result)
	}
}
