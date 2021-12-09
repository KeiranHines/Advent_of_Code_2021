package main

import (
	"strconv"
	"strings"
)

const ALL_CHARS = "abcdefg"

// FindCountOfUniqueNumbers finds the count of unique numbers.
// Unique numbers being the keys for 1, 4, 7 and 8.
func FindCountOfUniqueNumbers(input []string) int {
	count := 0
	for _, s := range input {
		split := strings.Split(s, " | ")
		final := strings.Split(split[1], " ")
		for _, num := range final {
			len := len(num)
			if len == 2 || len == 3 || len == 4 || len == 7 {
				count++
			}
		}
	}
	return count
}

// ProcessInputsAndOutputs processes the inputs and outputs to a count of the
// output numbers.
// Returns the count of the output numbers.
func ProcessInputsAndOutputs(input []string) int {
	count := 0
	for _, s := range input {
		split := strings.Split(s, " | ")
		random := strings.Split(split[0], " ")
		mapping := make(map[int]string, 10)
		fives := make([]string, 0)
		sixes := make([]string, 0)
		for _, num := range random {
			len := len(num)
			switch len {
			case 2:
				mapping[1] = num
			case 3:
				mapping[7] = num
			case 4:
				mapping[4] = num
			case 5:
				fives = append(fives, num)
			case 6:
				sixes = append(sixes, num)
			case 7:
				mapping[8] = num
			}
		}

		six, c, f := find6cf(mapping[1], sixes)
		sixes = remove(sixes, six)
		mapping[6] = six
		two, three, five := parseFives(fives, c, f)
		mapping[2] = two
		mapping[3] = three
		mapping[5] = five

		zero, nine := parseZeroNine(sixes, five)
		mapping[0] = zero
		mapping[9] = nine

		final := strings.Split(split[1], " ")
		flipped := flipMap(mapping)
		str := ""
		for _, word := range final {
			str = str + flipped[SortStringByCharacter(word)]
		}
		num, _ := strconv.Atoi(str)
		count += num
	}
	return count
}

// find6cf Finds 6 'c' and 'f' by using the segments with 6 segments lit
// and the fact 6 is the only six segment number that doesn't have one
// part of the segments for 1. That being the 'c' segment.
// Returns six, c, f in that order.
func find6cf(one string, sixes []string) (string, rune, rune) {
	chars := []rune(one)
	for i, c := range chars {
		for _, sixChars := range sixes {
			if !strings.ContainsRune(sixChars, c) {
				if i == 0 {
					return sixChars, c, chars[1]
				}
				return sixChars, c, chars[0]
			}
		}
	}
	return "", 'c', 'f'
}

// parseFives parses all the 5 segment displays to map each one to the correct
// Number.  5 wont contain 'c', 2 wont contain 'f' and three will contain both.
// Returns two, three and five in that order.
func parseFives(fives []string, c rune, f rune) (string, string, string) {
	var two string
	var three string
	var five string
	for _, fiveChars := range fives {
		if !strings.ContainsRune(fiveChars, c) {
			five = fiveChars
		} else if !strings.ContainsRune(fiveChars, f) {
			two = fiveChars
		} else {
			three = fiveChars
		}
	}
	return two, three, five
}

// parseZeroNine finds zero and nine using 5 because all the characters
// that make up 5 make up 9 but not zero.
// Returns Zero and Nine
func parseZeroNine(sixes []string, five string) (string, string) {
	option := sixes[0]
	for _, c := range []rune(five) {
		if !strings.ContainsRune(option, c) {
			return option, sixes[1]
		}
	}
	return sixes[1], option
}

// flipMap flips the numbered keys to be a sorted map of the sorted
// input seqments to the string representation e.g. mapping[1] = fc
// would flip to be flipped[cf] = "1"
// Retruns the flipped map.
func flipMap(mapping map[int]string) map[string]string {
	flipped := make(map[string]string, 10)
	for key, value := range mapping {
		sort := SortStringByCharacter(value)
		flipped[sort] = strconv.Itoa(key)
	}
	return flipped
}

// Removes an element from the list
// Returns the list.
func remove(list []string, elem string) []string {
	for i, s := range list {
		if s == elem {
			list[i] = list[len(list)-1]
			return list[:len(list)-1]
		}
	}
	return list
}
