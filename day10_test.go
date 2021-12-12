package main

import (
	"testing"
)

func TestDay1Example1(t *testing.T) {
	input := FileToStringArray("inputs/d10t", "\n")
	_, cor := FindCorruptedAndIncompleteLines(input)
	score := CalculateScore(cor)
	if score != 26397 {
		t.Fatalf("Expect %v got %v", 26397, score)
	}
}

func TestDay10Part1(t *testing.T) {
	ans := 167379
	input := FileToStringArray("inputs/d10a", "\n")
	_, cor := FindCorruptedAndIncompleteLines(input)
	score := CalculateScore(cor)
	if score != ans {
		t.Fatalf("Expect %v got %v", ans, score)
	}
}

func TestDay10Example2(t *testing.T) {
	input := FileToStringArray("inputs/d10t", "\n")
	inc, _ := FindCorruptedAndIncompleteLines(input)
	score := FindCompletingWinner(inc)
	if score != 288957 {
		t.Fatalf("Expect %v got %v", 288957, score)
	}
}

func TestDay10Part2(t *testing.T) {
	ans := 2776842859
	input := FileToStringArray("inputs/d10a", "\n")
	inc, _ := FindCorruptedAndIncompleteLines(input)
	score := FindCompletingWinner(inc)
	if score != ans {
		t.Fatalf("Expect %v got %v", ans, score)
	}
}
