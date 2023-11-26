package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var components = map[string]Component{}

type Component struct {
	name      string
	operation string
	argA      string
	argB      string
	value     uint16
	processed bool
}

func main() {
	for _, line := range strings.Split(input, "\n") {
		parseLine(line)
	}
	// part 1
	valA := getValue("a")
	println("Value of a is", valA)

	// part 2
	reset()
	componentB := components["b"]
	componentB.value = valA
	componentB.processed = true
	components[componentB.name] = componentB

	valA = getValue("a")
	println("Value of a is", valA)
}

func parseLine(line string) {
	parts := strings.Split(line, " ")
	arrowIndex := slices.Index(parts, "->")
	name := parts[arrowIndex+1]
	instructions := parts[:arrowIndex]

	switch len(instructions) {
	case 1:
		makeComponent(name, "WIRE", instructions[0], "")
	case 2:
		makeComponent(name, instructions[0], instructions[1], "")
	case 3:
		makeComponent(name, instructions[1], instructions[0], instructions[2])
	}
}

func makeComponent(name string, operation string, argA string, argB string) {
	component := Component{name, operation, argA, argB, 0, false}
	components[name] = component
}

func getValue(name string) uint16 {
	// return input value if string can be parsed to int
	intValue, err := strconv.Atoi(name)
	if err == nil {
		return uint16(intValue)
	}

	// return value if we already processed this component
	component := components[name]
	if component.processed {
		return component.value
	}

	// calculate value of the component recursively
	switch component.operation {
	case "WIRE":
		component.value = getValue(component.argA)
	case "NOT":
		component.value = ^getValue(component.argA)
	case "AND":
		component.value = getValue(component.argA) & getValue(component.argB)
	case "OR":
		component.value = getValue(component.argA) | getValue(component.argB)
	case "LSHIFT":
		component.value = getValue(component.argA) << getValue(component.argB)
	case "RSHIFT":
		component.value = getValue(component.argA) >> getValue(component.argB)
	default:
		panic(fmt.Sprintf("Unknown operation '%s'", component.operation))
	}

	component.processed = true
	components[name] = component
	return component.value
}

func reset() {
	for name, component := range components {
		component.processed = false
		components[name] = component
	}
}
