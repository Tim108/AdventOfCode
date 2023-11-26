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

var distances = map[string]map[string]int{}

func main() {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		locationA := parts[0]
		locationB := parts[2]
		cost, _ := strconv.Atoi(parts[4])
		addPath(locationA, locationB, cost)
	}
	route, cost := findRoutes([]string{}, 0)
	fmt.Printf("%v costs %d\n", route, cost)
}

func addPath(locationA string, locationB string, cost int) {
	mapA, okA := distances[locationA]
	if !okA {
		mapA = map[string]int{}
	}
	mapA[locationB] = cost
	distances[locationA] = mapA

	mapB, okB := distances[locationB]
	if !okB {
		mapB = map[string]int{}
	}
	mapB[locationA] = cost
	distances[locationB] = mapB
}

func findRoutes(path []string, totalCost int) (bestPath []string, bestCost int) {
	bestPath = path
	bestCost = totalCost

	options := optionsForPath(path)

	foundRoute := false
	for target, cost := range options {
		potentialPath, potentialCost := findRoutes(append(path, target), totalCost+cost)
		if !foundRoute || potentialCost < bestCost { // change the sign here for part 1 or part 2
			foundRoute = true
			bestCost = potentialCost
			bestPath = potentialPath
		}
	}
	return
}

func optionsForPath(path []string) map[string]int {
	if len(path) == 0 {
		return startingOptions()
	}

	options := map[string]int{}
	currentLocation := path[len(path)-1]
	for target, cost := range distances[currentLocation] {
		if !slices.Contains(path, target) {
			options[target] = cost
		}
	}
	return options
}

func startingOptions() map[string]int {
	options := map[string]int{}
	for target := range distances {
		options[target] = 0
	}
	return options
}
