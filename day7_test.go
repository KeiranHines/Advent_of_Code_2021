package main

import (
	"testing"
)

func TestDay7Example1(t *testing.T) {
	b := ReadFile("inputs/d7t")
	input := string(b)
	optimal, moveCount := FindOptimalPosition(input)
	if optimal != 2 {
		t.Fatal("Expected optimal position to be 2, got ", optimal)
	}
	if moveCount != 37 {
		t.Fatal("Expected optimal move count to be 37, got ", moveCount)
	}
}

func TestDay7Part1(t *testing.T) {
	ansPos := 316
	ansCount := 336721
	b := ReadFile("inputs/d7a")
	input := string(b)
	optimal, moveCount := FindOptimalPosition(input)
	if optimal != ansPos {
		t.Fatalf("Expected optimal position to be %v, got %v", ansPos, optimal)
	}
	if moveCount != ansCount {
		t.Fatalf("Expected optimal move count to be %v, got %v", ansCount, moveCount)
	}
}
