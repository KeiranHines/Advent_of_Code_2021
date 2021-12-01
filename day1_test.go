package main

import (
	"testing"
)

func TestPart1Example(t *testing.T) {
	count := CountIncreasing([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	if count != 7 {
		t.Fatal("expected 7, got", count)
	}
}

//TestInputPart1 runs the part 1 challenge
func TestInputPart1Method(t *testing.T) {
	data := fileToIntArray("inputs/d1a", "\n")
	count := CountIncreasing(data)
	if count != 1832 {
		t.Fatal("Expected 1832, got", count)
	}
}

func TestInputPart1UsingPart2Method(t *testing.T) {
	data := fileToIntArray("inputs/d1a", "\n")
	count := CountSlidingWindowIncrease(data, 1)
	if count != 1832 {
		t.Fatal("Expected x, got", count)
	}
}

func TestInputPart2Example(t *testing.T) {
	count := CountSlidingWindowIncrease([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 3)
	if count != 5 {
		t.Fatal("Expected 5 for part 2 example, got", count)
	}
}

func TestInputPart2(t *testing.T) {
	data := fileToIntArray("inputs/d1a", "\n")
	count := CountSlidingWindowIncrease(data, 3)
	if count != 1858 {
		t.Fatal("Expected x, got", count)
	}
}
