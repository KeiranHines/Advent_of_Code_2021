package main

import (
	"fmt"
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

	if boards[2][4][4] == "7" {
		t.Fatal("Expected board 3 postition 4,4 (the last number) to be 7")
	}
}

func TestPlayingRoundUpdatesTheBoard(t *testing.T) {
	board := [][]string{{"10", "11", "12"}, {"1", "2", "3"}, {"21", "22", "23"}}
	PlayRound("21", board)
	fmt.Printf("%v\n", board)
}

func TestDay4ExampleWinner(t *testing.T) {
	input := FileToStringArray("inputs/d4t", "\n")
	rounds, boards := ParseFileString(input)
	winner, round := PlayBingo(rounds, boards)
	if winner != 2 {
		t.Fatal("Player 3 (board 2) should win the example not", winner)
	}
	if round != 24 {
		t.Fatal("The winning nuber for player 3 should be 24 not", round)
	}

	count := CalcRemainingNumbers(boards[winner])
	if count != 188 {
		t.Fatal("Incorrect sum of remaining numbers, expected 188, got ", count)
	}
	result := count * round
	if result != 4512 {
		t.Fatal("One of the other tests should have failed.")
	}
}

func TestDay4Part1Input(t *testing.T) {
	ans := 32844
	input := FileToStringArray("inputs/d4a", "\n")
	rounds, boards := ParseFileString(input)
	winner, round := PlayBingo(rounds, boards)
	count := CalcRemainingNumbers(boards[winner])
	result := count * round
	if result != ans {
		t.Fatal("Expected", ans, ", got", result)
	}
}
