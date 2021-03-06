package main

import (
	"strconv"
	"strings"
)

//point Typedef for an x,y pair
type point struct {
	x uint16
	y uint16
}

// ParsePipeMap parses a string representation of a map of pipes
// Accepts diagonal lines on a 45 degree angle if allowDiagonal is set
// Returns a map of points and the times they overlap.
func ParsePipeMap(input []string, allowDiagonal bool) map[point]int16 {
	pointMap := make(map[point]int16)
	var validityCheck func(int16, int16, int16, int16) bool
	if allowDiagonal {
		validityCheck = diagonalCheck
	} else {
		validityCheck = nonDiagonalCheck
	}
	for _, line := range input {
		split := strings.Split(line, " -> ")
		start := strings.Split(split[0], ",")
		end := strings.Split(split[1], ",")
		t, _ := strconv.Atoi(start[0])
		x1 := int16(t)
		t, _ = strconv.Atoi(start[1])
		y1 := int16(t)
		t, _ = strconv.Atoi(end[0])
		x2 := int16(t)
		t, _ = strconv.Atoi(end[1])
		y2 := int16(t)

		if validityCheck(x1, y1, x2, y2) {
			incrementPoint(pointMap, x1, y1)
			incrementPoint(pointMap, x2, y2)

			if !nonDiagonalCheck(x1, y1, x2, y2) {
				processDiagonalLine(pointMap, x1, y1, x2, y2)
			} else {
				for i := x1 + 1; i < x2; i++ {
					incrementPoint(pointMap, i, y1)
				}
				for i := y1 + 1; i < y2; i++ {
					incrementPoint(pointMap, x1, i)
				}
				for i := x2 + 1; i < x1; i++ {
					incrementPoint(pointMap, i, y1)
				}
				for i := y2 + 1; i < y1; i++ {
					incrementPoint(pointMap, x1, i)
				}
			}
		}
	}

	return pointMap
}

// incrementPoint Increments the point type in the pointmap.
func incrementPoint(pointMap map[point]int16, x int16, y int16) {
	p := point{uint16(x), uint16(y)}
	pointMap[p]++
}

// nonDiagonalCheck only allows lines that are horizontal or vertical.
func nonDiagonalCheck(x1 int16, y1 int16, x2 int16, y2 int16) bool {
	return x1 == x2 || y1 == y2
}

// diagonalCheck only allows line that are horizontal, vertical or
// on a 45 degree angle.
func diagonalCheck(x1 int16, y1 int16, x2 int16, y2 int16) bool {
	x := x1 - x2
	if x < 0 {
		x = -x
	}

	y := y1 - y2
	if y < 0 {
		y = -y
	}

	return x1 == x2 || y1 == y2 || x == y
}

// processDiagonalLine adds a diagonal line from x1,y2 to x2,y2 to the point map
func processDiagonalLine(pointMap map[point]int16, x1 int16, y1 int16, x2 int16, y2 int16) {
	if x2 > x1 {
		x := x2 - x1
		if y2 > y1 {
			for i := int16(1); i < x; i++ {
				incrementPoint(pointMap, x1+i, y1+i)
			}
		} else {
			for i := int16(1); i < x; i++ {
				incrementPoint(pointMap, x1+i, y1-i)
			}
		}
	} else {
		x := x1 - x2
		if y2 > y1 {
			for i := int16(1); i < x; i++ {
				incrementPoint(pointMap, x1-i, y1+i)
			}
		} else {
			for i := int16(1); i < x; i++ {
				incrementPoint(pointMap, x1-i, y1-i)
			}
		}
	}
}
