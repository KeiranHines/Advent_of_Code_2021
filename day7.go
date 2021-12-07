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
