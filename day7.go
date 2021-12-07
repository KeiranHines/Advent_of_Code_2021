package main

import (
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	key   int
	value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[j].value < p[i].value } // Less is actually more likely

// FindOptimalPosition finds the optimal position to move all crabs to assuming a cost of 1 fuel cell per move
// This method assumes, based on the fuel cost the optimal solution will be at a position where one set of crabs do not move
// input: a comma seperated list of crab positions.
// Returns the optimal position and the cost.
func FindOptimalPosition(input string) (int, int) {
	data := strings.Split(input, ",")
	mapping := make(map[int]int, len(data))
	for _, v := range data {
		x, _ := strconv.Atoi(v)
		mapping[x]++
	}

	values := make(PairList, len(mapping))
	i := 0
	for k, v := range mapping {
		values[i] = Pair{k, v}
		i++
	}

	sort.Sort(values)

	bestPos := -1
	bestCount := -1

	for _, pair := range values {
		sum := 0
		for _, pair2 := range values {
			sum += (Max(pair.key, pair2.key) - Min(pair.key, pair2.key)) * pair2.value
		}
		if bestCount == -1 || sum < bestCount {
			bestCount = sum
			bestPos = pair.key
		}
	}
	return bestPos, bestCount
}

// FindOptimalPositionCostly Finds the optimal position to move all crabs to assuming a cost of n+1 per move where n is the number of
// positions already moved. E.g. the first move costs 1, the second 3, the third 6,
// This method assumes the optimal position will be somewhere between the min and max of all starting positions.
// input: a comma seperated list of crab positions.
// Returns the optimal position and the cost.
func FindOptimalPositionCostly(input string) (int, int) {
	data := strings.Split(input, ",")
	mapping := make(map[int]int, len(data))
	minX := -1
	maxX := -1
	for _, v := range data {
		x, _ := strconv.Atoi(v)
		if minX == -1 || x < minX {
			minX = x
		}
		if maxX < x {
			maxX = x
		}
		mapping[x]++
	}

	bestPos := -1
	bestCount := -1

	for i := minX; i <= maxX; i++ {
		sum := 0
		for key, value := range mapping {
			diff := (Max(i, key) - Min(i, key))
			cost := (float32(diff+1) / float32(2)) * float32(diff)
			sum += int(cost * float32(value))
		}
		// println("Cost of ", pair.key, "is", sum)
		if bestCount == -1 || sum < bestCount {
			bestCount = sum
			bestPos = i
		}
	}
	return bestPos, bestCount
}
