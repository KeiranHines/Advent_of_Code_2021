package main

import (
	"strconv"
	"strings"
)

//point Typedef for an x,y pair
type point struct {
	x int
	y int
}

// ParsePipeMap parses a string representation of a map of pipes
// Accepts diagonal lines on a 45 degree angle if allowDiagonal is set
// Returns a map of points and the times they overlap.
func ParsePipeMap(input []string, allowDiagonal bool) map[point]int {
	pointMap := make(map[point]int)
	var validityCheck func(int, int, int, int) bool
	if allowDiagonal {
		validityCheck = diagonalCheck
	} else {
		validityCheck = nonDiagonalCheck
	}
	for _, line := range input {
		split := strings.Split(line, " -> ")
		start := strings.Split(split[0], ",")
		end := strings.Split(split[1], ",")
		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])
		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])

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
func incrementPoint(pointMap map[point]int, x int, y int) {
	p := point{x, y}
	v := pointMap[p]
	pointMap[p] = v + 1
}

// nonDiagonalCheck only allows lines that are horizontal or vertical.
func nonDiagonalCheck(x1 int, y1 int, x2 int, y2 int) bool {
	return x1 == x2 || y1 == y2
}

// diagonalCheck only allows line that are horizontal, vertical or
// on a 45 degree angle.
func diagonalCheck(x1 int, y1 int, x2 int, y2 int) bool {
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
func processDiagonalLine(pointMap map[point]int, x1 int, y1 int, x2 int, y2 int) {
	if x2 > x1 {
		x := x2 - x1
		if y2 > y1 {
			for i := 1; i < x; i++ {
				incrementPoint(pointMap, x1+i, y1+i)
			}
		} else {
			for i := 1; i < x; i++ {
				incrementPoint(pointMap, x1+i, y1-i)
			}
		}
	} else {
		x := x1 - x2
		if y2 > y1 {
			for i := 1; i < x; i++ {
				incrementPoint(pointMap, x1-i, y1+i)
			}
		} else {
			for i := 1; i < x; i++ {
				incrementPoint(pointMap, x1-i, y1-i)
			}
		}
	}
}
