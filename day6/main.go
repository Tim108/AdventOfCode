package main

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var instructionSeparator = " "
var setInstruction = "turn"
var onInstruction = "on"
var offInstruction = "off"
var toggleInstruction = "toggle"
var rangeInstruction = "through"
var coordinateSeparator = ","

var lights = [1000][1000]bool{}

type Coordinate struct {
	x int
	y int
}

func main() {
	followInstructions(strings.Split(input, "\n"))
	println(countLights(), "lights lit!")
}

func followInstructions(instructionsList []string) {
	for _, word := range instructionsList {
		instructions := strings.Split(word, instructionSeparator)
		rangeIndex := slices.Index(instructions, rangeInstruction)
		from := stringToCoordinate(instructions[rangeIndex-1])
		to := stringToCoordinate(instructions[rangeIndex+1])

		if instructions[0] == setInstruction {
			if instructions[1] == onInstruction {
				set(from, to, true)
			} else if instructions[1] == offInstruction {
				set(from, to, false)
			}
		} else if instructions[0] == toggleInstruction {
			toggle(from, to)
		}
	}
}

func set(from Coordinate, to Coordinate, state bool) {
	for x := from.x; x <= to.x; x++ {
		for y := from.y; y <= to.y; y++ {
			lights[x][y] = state
		}
	}
}

func toggle(from Coordinate, to Coordinate) {
	for x := from.x; x <= to.x; x++ {
		for y := from.y; y <= to.y; y++ {
			lights[x][y] = !lights[x][y]
		}
	}
}

func countLights() (count int) {
	for _, rowOfLights := range lights {
		for _, light := range rowOfLights {
			if light {
				count++
			}
		}
	}
	return
}

func stringToCoordinate(str string) Coordinate {
	numberStrings := strings.Split(str, coordinateSeparator)
	x, errX := strconv.Atoi(numberStrings[0])
	y, errY := strconv.Atoi(numberStrings[1])

	if errX != nil {
		panic(errX)
	} else if errY != nil {
		panic(errY)
	}

	return Coordinate{x, y}
}
