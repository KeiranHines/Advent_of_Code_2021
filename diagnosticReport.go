package main

import (
	"math"
)

// CalculateGammaEpsilon Calculates the gamma and epsilon
// values for a byte array of \n seperated binary values
func CalculateGammaEpsilon(input []byte) (int, int) {
	size := 0
	for i, char := range input {
		if char == 10 {
			size = i
			break
		}
		i++
	}
	posCount := make([]int, size)
	lineCount := 0
	pos := 0
	for _, char := range input {
		switch char {
		case 10: // Newline
			lineCount++
			pos = 0
		case 49: //1
			posCount[pos]++
			pos++
		case 48: //0
			pos++
		}
	}

	half := lineCount / 2
	gamma := 0
	eps := 0
	for i, count := range posCount {
		if count > half {
			// Most common is 1
			gamma += int(math.Pow(2, float64(size-i-1)))
		} else {
			eps += int(math.Pow(2, float64(size-i-1)))
		}
	}
	return gamma, eps
}

// Calculates the most common bit and least common bit from
// an array of binary values.
func GetMcbLcb(input []string, pos int) (byte, byte) {
	count := 0
	for _, v := range input {
		if v[pos] == 49 {
			count++
		}
	}
	if count*2 == len(input) { // Tie breaker of equal
		return 49, 48
	}
	if count > len(input)/2 {
		return 49, 48
	}
	return 48, 49
}

// CalculateOxAndCo2 takes a string array of binary values
// and calculates the oxygen and Co2 values
func CalculateOxAndCo2(input []string) (int, int) {
	size := len(input[0])
	oxyList := input
	co2List := input
	for i := 0; i < size; i++ {
		temp := make([]string, 0)
		if len(oxyList) != 1 {
			mcb, _ := GetMcbLcb(oxyList, i)
			for _, v := range oxyList {
				if v[i] == mcb {
					temp = append(temp, v)
				}
			}
			oxyList = temp
		}
		if len(co2List) != 1 {
			_, lcb := GetMcbLcb(co2List, i)
			temp = make([]string, 0)
			for _, v := range co2List {
				if v[i] == lcb {
					temp = append(temp, v)
				}
			}
			co2List = temp
		}
	}
	oxy := 0
	co2 := 0
	for i := 0; i < size; i++ {
		if oxyList[0][i] == 49 {
			oxy += int(math.Pow(2, float64(size-i-1)))
		}
		if co2List[0][i] == 49 {
			co2 += int(math.Pow(2, float64(size-i-1)))
		}
	}
	return oxy, co2
}
