package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadFile(filename string) []byte {
	data, _ := os.ReadFile(filename)
	return data
}

func FileToStringArray(filename string, delim string) []string {
	data, _ := os.ReadFile(filename)
	return strings.Split(string(data), delim)

}

func FileToIntArray(filename string, delim string) []int {
	lines := FileToStringArray(filename, delim)
	parsed := make([]int, len(lines))

	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		parsed[i] = num
	}

	return parsed
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func SortStringByCharacter(s string) string {
	var r ByRune = []rune(s)
	sort.Sort(r)
	return string(r)
}
