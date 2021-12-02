package main

import (
	"strconv"
	"strings"
)

// ProcessCommandsCorrectly Processes the submarine commands based on the expected insructions
// x : The initial x position
// y : The initial y position
// Returns The new x and y value after the sub has executed all the instructions
func ProcessCommandsBasic(data []string, x int, y int) (int, int) {
	for _, command := range data {
		split := strings.Split(command, " ")
		action := split[0]
		mag, _ := strconv.Atoi(split[1])
		switch action {
		case "forward":
			x += mag
			break
		case "down":
			y += mag
			break
		case "up":
			y -= mag
			break
		}
	}
	return x, y
}

// ProcessCommandsCorrectly Processes the submarine commands as per the instruction manual
// x : The initial x position
// y : The initial y position
// aim : The initial tragectory of the submarine
// Returns The new x and y value after the sub has executed all the instructions
func ProcessCommandsCorrectly(data []string, x int, y int, aim int) (int, int) {
	for _, command := range data {
		split := strings.Split(command, " ")
		action := split[0]
		mag, _ := strconv.Atoi(split[1])
		switch action {
		case "forward":
			x += mag
			y += mag * aim
			break
		case "down":
			aim += mag
			break
		case "up":
			aim -= mag
			break
		}
	}
	return x, y
}
