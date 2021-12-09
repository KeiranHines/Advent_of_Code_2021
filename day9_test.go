package main

import (
	"sort"
	"testing"
)

func TestDay9Example1(t *testing.T) {
	input := FileToStringArray("inputs/d9t", "\n")
	lowPoints := FindLowPoints(input)
	sum := CalculateRiskLevel(lowPoints)
	if sum != 15 {
		t.Fatal("Exepected 15 got ", sum)
	}
}

func TestDay9Part1(t *testing.T) {
	ans := 444
	input := FileToStringArray("inputs/d9a", "\n")
	lowPoints := FindLowPoints(input)
	sum := CalculateRiskLevel(lowPoints)
	if sum != ans {
		t.Fatalf("Exepected %v got %v", ans, sum)
	}
}

func TestDay9Example2(t *testing.T) {
	ans := 1134
	input := FileToStringArray("inputs/d9t", "\n")
	lowPoints := FindLowPoints(input)
	basinSizes := make([]int, 0)
	for _, point := range lowPoints {
		basin := CalculateBasinSize(input, point)
		basinSizes = append(basinSizes, len(basin))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	result := basinSizes[0] * basinSizes[1] * basinSizes[2]
	if result != ans {
		t.Fatalf("Expected %v but got %v", ans, result)
	}
}

func TestDay9Part2(t *testing.T) {
	ans := 1168440
	input := FileToStringArray("inputs/d9a", "\n")
	lowPoints := FindLowPoints(input)
	basinSizes := make([]int, 0)
	for _, point := range lowPoints {
		basin := CalculateBasinSize(input, point)
		basinSizes = append(basinSizes, len(basin))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	result := basinSizes[0] * basinSizes[1] * basinSizes[2]
	if result != ans {
		t.Fatalf("Expected %v but got %v", ans, result)
	}
}
