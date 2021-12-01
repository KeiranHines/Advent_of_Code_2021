package main

import (
	"os"
	"strconv"
	"strings"
)

func fileToIntArray(filename string, delim string) []int {
	var parsed []int
	data, _ := os.ReadFile("inputs/d1a")
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		parsed = append(parsed, num)
	}

	return parsed
}
