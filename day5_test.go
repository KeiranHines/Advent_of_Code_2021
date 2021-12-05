package main

import (
	"testing"
)

func TestBasicOverlap(t *testing.T) {
	x := []string{"0,0 -> 0,4", "0,0 -> 4,0"}
	pointMap := ParsePipeMap(x, false)
	if pointMap[point{0, 0}] != 2 {
		t.Fatal("Expected 2 hits at 0,0 found: ", pointMap[point{0, 0}])
	}
	if len(pointMap) != 9 {
		t.Fatal("Expected 9 points to be marked on the map not ", len(pointMap))
	}
}

func TestDay5Example(t *testing.T) {
	input := FileToStringArray("inputs/d5t", "\n")
	pointMap := ParsePipeMap(input, false)
	gt2 := 0
	for _, count := range pointMap {
		if count > 1 {
			gt2++
		}
	}
	if gt2 != 5 {
		t.Fatal("Expected 5 points of overlap but found", gt2)
	}
}

func TestDay5Part1(t *testing.T) {
	ans := 6225
	input := FileToStringArray("inputs/d5a", "\n")
	pointMap := ParsePipeMap(input, false)
	gt2 := 0
	for _, count := range pointMap {
		if count > 1 {
			gt2++
		}
	}
	if gt2 != ans {
		t.Fatalf("Expected %v points of overlap but found %v", ans, gt2)
	}
}

func TestDiagonalOverlap(t *testing.T) {
	x := []string{"0,0 -> 0,4", "0,0 -> 4,0", "0,0 -> 4,4"}
	pointMap := ParsePipeMap(x, true)
	if pointMap[point{0, 0}] != 3 {
		t.Fatal("Expected 3 hits at 0,0 found: ", pointMap[point{0, 0}])
	}
	if len(pointMap) != 13 {
		t.Fatal("Expected 15 points to be marked on the map not ", len(pointMap))
	}
}

func TestDay5ExamplePart2(t *testing.T) {
	input := FileToStringArray("inputs/d5t", "\n")
	pointMap := ParsePipeMap(input, true)
	gt2 := 0
	for _, count := range pointMap {
		if count > 1 {
			gt2++
		}
	}
	if gt2 != 12 {
		t.Fatal("Expected 12 points of overlap but found", gt2)
	}
}

func TestDay5Part2(t *testing.T) {
	ans := 22116
	input := FileToStringArray("inputs/d5a", "\n")
	pointMap := ParsePipeMap(input, true)
	gt2 := 0
	for _, count := range pointMap {
		if count > 1 {
			gt2++
		}
	}
	if gt2 != ans {
		t.Fatalf("Expected %v points of overlap but found %v", ans, gt2)
	}
}
