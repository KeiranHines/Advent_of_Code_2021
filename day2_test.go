package main

import (
	"testing"
)

func TestDay2Part1Example(t *testing.T) {
	input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	x, y := ProcessCommandsBasic(input, 0, 0)
	if x*y != 150 {
		t.Fatalf("Expected x * y to be 150, got %v from x=%v y=%v", x*y, x, y)
	}
}

func TestDay2Part1Input(t *testing.T) {
	ans := 1660158
	input := FileToStringArray("inputs/d2a", "\n")
	x, y := ProcessCommandsBasic(input, 0, 0)
	if x*y != ans {
		t.Fatalf("Expected x * y to be %v, got %v from x=%v y=%v", ans, x*y, x, y)
	}
}

func TestDay2Part2Example(t *testing.T) {
	ans := 900
	input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	x, y := ProcessCommandsCorrectly(input, 0, 0, 0)
	if x*y != ans {
		t.Fatalf("Expected x * y to be %v, got %v from x=%v y=%v", ans, x*y, x, y)
	}
}

func TestDay2Part2Input(t *testing.T) {
	ans := 1604592846
	input := FileToStringArray("inputs/d2a", "\n")
	x, y := ProcessCommandsCorrectly(input, 0, 0, 0)
	if x*y != ans {
		t.Fatalf("Expected x * y to be %v, got %v from x=%v y=%v", ans, x*y, x, y)
	}
}
