package main

import (
	"os"
	"strconv"
	"strings"
)

func FileToStringArray(filename string, delim string) []string {
	data, _ := os.ReadFile(filename)
	return strings.Split(string(data), delim)

}

func FileToIntArray(filename string, delim string) []int {
	var parsed []int
	lines := FileToStringArray(filename, delim)
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		parsed = append(parsed, num)
	}

	return parsed
}
