package main

import (
	"testing"
)

func TestDay11SimpleExample(t *testing.T) {
	input := []string{"11111", "19991", "19191", "19991", "11111"}
	_, flashes := SimulateOctopi(input, 2)
	if flashes != 9 {
		t.Fatal("Expected 9 flashes, got ", flashes)
	}
}

func TestDay11Example1(t *testing.T) {
	ans := 1656
	input := FileToStringArray("inputs/d11t", "\n")

	_, flashes := SimulateOctopi(input, 100)
	if flashes != ans {
		t.Fatalf("Expected %v got %v", ans, flashes)
	}
}

func TestDay11Part1(t *testing.T) {
	ans := 1732
	input := FileToStringArray("inputs/d11a", "\n")

	_, flashes := SimulateOctopi(input, 100)
	if flashes != ans {
		t.Fatalf("Expected %v got %v", ans, flashes)
	}
}

func TestDay11Example2(t *testing.T) {
	ans := 195
	input := FileToStringArray("inputs/d11t", "\n")

	round := FindFirstSyncPoint(input)
	if round != ans {
		t.Fatalf("Expected %v got %v", ans, round)
	}
}

func TestDay11Part2(t *testing.T) {
	ans := 290
	input := FileToStringArray("inputs/d11a", "\n")

	round := FindFirstSyncPoint(input)
	if round != ans {
		t.Fatalf("Expected %v got %v", ans, round)
	}
}
