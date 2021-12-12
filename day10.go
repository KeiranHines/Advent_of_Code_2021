package main

import "sort"

type Stack []rune

var bracket = map[rune]rune{
	'>': '<',
	']': '[',
	'}': '{',
	')': '(',
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return rune(0), false
	}
	i := len(*s) - 1
	e := (*s)[i]
	*s = (*s)[:i]
	return e, true
}

func FindCorruptedAndIncompleteLines(input []string) (map[string]int, map[string]int) {
	incomplete := make(map[string]int, 0)
	corrupt := make(map[string]int, 0)
	for _, line := range input {
		parseLine(line, incomplete, corrupt)
	}
	return incomplete, corrupt
}

func parseLine(line string, incomplete map[string]int, corrupt map[string]int) {
	var stack Stack
	for i, r := range line {
		switch r {
		case '(', '[', '{', '<':
			stack.Push(r)
		case ')', ']', '}', '>':
			o, _ := stack.Pop()
			if o != bracket[r] {
				corrupt[line] = i
				return
			}
		}
	}
	// If Stack is not empty the line as not been fully closed.
	if !stack.IsEmpty() {
		incomplete[line] = CompleteLine(stack)
	}
}

func CalculateScore(corrupt map[string]int) int {
	count := 0
	for k, v := range corrupt {
		b := k[v]
		r := rune(b)
		switch r {
		case ')':
			count += 3
		case ']':
			count += 57
		case '}':
			count += 1197
		case '>':
			count += 25137
		}
	}
	return count
}

func CompleteLine(stack Stack) int {
	score := 0
	for {
		o, f := stack.Pop()
		if !f {
			return score
		} else {
			score *= 5
			switch o {
			case '(':
				score += 1
			case '[':
				score += 2
			case '{':
				score += 3
			case '<':
				score += 4
			}
		}
	}
}

func FindCompletingWinner(inc map[string]int) int {
	scoreList := make([]int, 0)
	for _, v := range inc {
		scoreList = append(scoreList, v)
	}
	sort.Ints(scoreList)
	winner := len(inc) / 2
	return scoreList[winner]
}
