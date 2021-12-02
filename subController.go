package main

import (
	"strconv"
	"strings"
)

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
