package main

import (
	"testing"
)

func TestDay6Part1Example(t *testing.T) {
	file := ReadFile("inputs/d6t")
	data := string(file)
	fish := ProcessFish(data, 18)
	if fish != 26 {
		t.Fatal("Expected 26 fish, got ", fish)
	}

	fish = ProcessFish(data, 80)
	if fish != 5934 {
		t.Fatal("Exepected 5934 fish got", fish)
	}
}

func TestDay6Part1(t *testing.T) {
	ans := 396210
	file := ReadFile("inputs/d6a")
	data := string(file)
	fish := ProcessFish(data, 80)
	if fish != ans {
		t.Fatalf("Expected %v fish, got %v", ans, fish)
	}
}

func TestDay6Part2Example(t *testing.T) {
	file := ReadFile("inputs/d6t")
	data := string(file)
	fish := ProcessFishOptimal(data, 256, 6, 8)
	if fish != 26984457539 {
		t.Fatal("Expected 26984457539 fish, got ", fish)
	}
}

func TestDay6Part2(t *testing.T) {
	ans := 1770823541496
	file := ReadFile("inputs/d6a")
	data := string(file)
	fish := ProcessFishOptimal(data, 256, 6, 8)
	if fish != ans {
		t.Fatalf("Expected %v fish, got %v", ans, fish)
	}
}
