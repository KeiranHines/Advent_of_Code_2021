package main

import (
	"strconv"
	"strings"
)

// ProcessFish Naive approach to brute force the exponential growth
// of fish. This soultion is fine for a small number of fish and a
// small number of days.
// initial: The fish to start with and their initial days to reproduction.
// days: The number of days to run the simulation for.
// Returns the number of fish after n days.
func ProcessFish(initial string, days int) int {
	str := strings.Split(initial, ",")
	fish := make([]int, len(str))
	for i, v := range str {
		next, _ := strconv.Atoi(v)
		fish[i] = next
	}

	for i := 0; i < days; i++ {
		current := len(fish)
		for j := 0; j < current; j++ {
			if fish[j] == 0 {
				fish[j] = 6
				fish = append(fish, 8)
			} else {
				fish[j] = fish[j] - 1
			}
		}
	}
	return len(fish)
}

// ProcessFishOptimal simulates the exponential growth rate of fish in a scalable solution.
// initial: The fish to start with and their initial days to reproduction.
// days: The number of days to run the simulation for.
// repoRate: The number of days for reproduction after the initial reproduction.
// initRepoRate: The number of days for the initial reproduction.
// Returns the number of fish after n days.
func ProcessFishOptimal(initial string, days int, repoRate int, initRepoRate int) int {
	str := strings.Split(initial, ",")
	fishMap := make([]int, initRepoRate+1)
	for _, v := range str {
		num, _ := strconv.Atoi(v)
		fishMap[num]++
	}
	for i := 0; i < days; i++ {
		newFish := fishMap[0]
		for j := 1; j <= initRepoRate; j++ {
			fishMap[j-1] = fishMap[j]
		}
		fishMap[initRepoRate] = newFish
		fishMap[repoRate] += newFish
	}
	sum := 0
	for _, v := range fishMap {
		sum += v
	}
	return sum
}
