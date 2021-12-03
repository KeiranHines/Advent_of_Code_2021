package main

import (
	"testing"
)

func TestDay3Part1Example(t *testing.T) {
	ans := 198
	bytes := []byte{48, 48, 49, 48, 48, 10, 49, 49, 49, 49, 48, 10, 49, 48, 49, 49, 48, 10, 49, 48, 49, 49, 49, 10, 49, 48, 49, 48, 49, 10, 48, 49, 49, 49, 49, 10, 48, 48, 49, 49, 49, 10, 49, 49, 49, 48, 48, 10, 49, 48, 48, 48, 48, 10, 49, 49, 48, 48, 49, 10, 48, 48, 48, 49, 48, 10, 48, 49, 48, 49, 48}
	gamma, e := CalculateGammaEpsilon(bytes)
	if gamma*e != ans {
		t.Fatalf("Expected gamma * e to be %v, got %v from gamma=%v e=%v", ans, gamma*e, gamma, e)
	}
}

func TestDay3Part1Input(t *testing.T) {
	ans := 3882564
	bytes := ReadFile("inputs/d3a")
	gamma, e := CalculateGammaEpsilon(bytes)
	if gamma*e != ans {
		t.Fatalf("Expected gamma * e to be %v, got %v from gamma=%v e=%v", ans, gamma*e, gamma, e)
	}
}

func TestDay3Part2Example(t *testing.T) {
	ans := 230
	bins := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	oxy, co2 := CalculateOxAndCo2(bins)
	if oxy*co2 != ans {
		t.Fatalf("Expected oxy * co2 to be %v, got %v from oxy=%v co2=%v", ans, oxy*co2, oxy, co2)
	}
}

func TestDay3Part2Input(t *testing.T) {
	ans := 3385170
	bins := FileToStringArray("inputs/d3a", "\n")
	oxy, co2 := CalculateOxAndCo2(bins)
	if oxy*co2 != ans {
		t.Fatalf("Expected oxy * co2 to be %v, got %v from oxy=%v co2=%v", ans, oxy*co2, oxy, co2)
	}
}
