package main

import "strconv"

type heatPoint struct {
	x   int
	y   int
	num int
}

// FindLowPoints finds the low points in a heatmap.
// Returns an array of heatPoints being the lowest nearby points.
func FindLowPoints(input []string) []heatPoint {
	maxY := len(input) - 1
	maxX := len(input[0]) - 1
	lowPoints := make([]heatPoint, 0)
	for y, line := range input {
		for x, char := range line {
			i, _ := strconv.Atoi(string(char))
			if isLowPoint(input, line, maxX, maxY, x, y, i) {
				lowPoints = append(lowPoints, heatPoint{x, y, i})
			}
		}
	}
	return lowPoints
}

// isLowPoint
// Returns two if the point is the lowest of the adjacent points.
func isLowPoint(input []string, line string, maxX int, maxY int, x int, y int, i int) bool {
	left := Max(x-1, 0)
	right := Min(x+1, maxX)
	top := Max(y-1, 0)
	bottom := Min(y+1, maxY)

	leftNum, _ := strconv.Atoi(string(line[left]))
	rightNum, _ := strconv.Atoi(string(line[right]))
	topNum, _ := strconv.Atoi(string(input[top][x]))
	bottomNum, _ := strconv.Atoi(string(input[bottom][x]))

	return (x == 0 || i < leftNum) && (x == maxX || i < rightNum) && (y == 0 || i < topNum) && (y == maxY || i < bottomNum)
}

// CalculateRiskLevel Calculates the risk points of the low points.
// Returns the risk of each low point
func CalculateRiskLevel(heatPoints []heatPoint) int {
	sum := 0
	for _, point := range heatPoints {
		sum += (point.num + 1)
	}
	return sum
}

// CalculateBasinSize Calculates the basin size using a Breadth first search approach.
// Returns a map of the points that make up a basin starting at a given lowpoint point.
func CalculateBasinSize(input []string, point heatPoint) map[heatPoint]bool {
	maxY := len(input) - 1
	maxX := len(input[0]) - 1
	end := byte('9')
	seenPoints := make(map[heatPoint]bool)
	point.num = 0
	queue := []heatPoint{point}

	for {
		p := queue[0]
		queue = queue[1:]
		x := p.x
		y := p.y
		if !seenPoints[p] && input[y][x] != end {
			seenPoints[p] = true
			if x > 0 {
				queue = append(queue, heatPoint{x - 1, y, 0})
			}
			if x < maxX {
				queue = append(queue, heatPoint{x + 1, y, 0})
			}

			if y > 0 {
				queue = append(queue, heatPoint{x, y - 1, 0})
			}
			if y < maxY {
				queue = append(queue, heatPoint{x, y + 1, 0})
			}
		}
		if len(queue) == 0 {
			return seenPoints
		}
	}
}

/*
 1  procedure BFS(G, root) is
 2      let Q be a queue
 3      label root as explored
 4      Q.enqueue(root)
 5      while Q is not empty do
 6          v := Q.dequeue()
 7          if v is the goal then
 8              return v
 9          for all edges from v to w in G.adjacentEdges(v) do
10              if w is not labeled as explored then
11                  label w as explored
12                  Q.enqueue(w)
*/
