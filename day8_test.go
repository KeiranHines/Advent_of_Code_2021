package main

import (
	"testing"
)

func TestDay8Example1(t *testing.T) {
	input := FileToStringArray("inputs/d8t", "\n")
	result := FindCountOfUniqueNumbers(input)
	if result != 26 {
		t.Fatal("Exepected 26 got ", result)
	}
}

func TestDay8Part1(t *testing.T) {
	ans := 365
	input := FileToStringArray("inputs/d8a", "\n")
	result := FindCountOfUniqueNumbers(input)
	if result != ans {
		t.Fatalf("Exepected %v got %v", ans, result)
	}
}

func TestAbusingOne(t *testing.T) {
	zero := "abcefg"
	one := "cf"
	six := "abdefg"
	nine := "abcdfg"

	sixes := make([]string, 3)
	sixes[0] = zero
	sixes[1] = six
	sixes[2] = nine

	s, c, f := find6cf(one, sixes)
	if s != six {
		t.Fatal("expected to get six back not", s)
	}
	if c != 'c' {
		t.Fatal("expected to get c back not", c)
	}
	if f != 'f' {
		t.Fatal("expected to get f back not", f)
	}
}

func TestParseFives(t *testing.T) {
	fives := make([]string, 3)
	two := "acdeg"
	three := "acdfg"
	five := "abdfg"

	fives[0] = two
	fives[1] = three
	fives[2] = five
	tw, th, f := parseFives(fives, 'c', 'f')
	if tw != two {
		t.Fatal("expected to get two back not", tw)
	}
	if th != three {
		t.Fatal("expected to get three back not", th)
	}
	if f != five {
		t.Fatal("expected to get five back not", f)
	}
}

func TestZeroNine(t *testing.T) {
	zero := "abcefg"
	nine := "abcdfg"
	five := "abdfg"

	sixes := make([]string, 2)
	sixes[0] = zero
	sixes[1] = nine

	z, n := parseZeroNine(sixes, five)
	if z != zero {
		t.Fatal("expected to get zero back not", z)
	}
	if n != nine {
		t.Fatal("expected to get nine back not", n)
	}
}

func TestDay8Example2(t *testing.T) {
	input := FileToStringArray("inputs/d8t", "\n")
	result := ProcessInputsAndOutputs(input)
	if result != 61229 {
		t.Fatal("Exepected 61229 got ", result)
	}
}

func TestDay8Part2(t *testing.T) {
	ans := 975706
	input := FileToStringArray("inputs/d8a", "\n")
	result := ProcessInputsAndOutputs(input)
	if result != ans {
		t.Fatalf("Exepected %v got %v", ans, result)
	}
}
