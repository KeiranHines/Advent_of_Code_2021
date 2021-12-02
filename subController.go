package main

import (
	"strconv"
	"strings"
)

func ProcessCommands(data []string, x int, y int) (int, int) {
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
