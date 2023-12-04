package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var auntDataMap map[int]map[string]int

func main() {
	parseInput()
	fmt.Printf("%v\n", auntDataMap)

	auntToFind := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	bestAunt, found := findClosestMatch(auntToFind)
	if found {
		println("Found an aunt")
		fmt.Printf("Best aunt: %d", bestAunt)
	} else {
		println("Did not find an aunt")
	}
}

func findClosestMatch(auntToFind map[string]int) (int, bool) {
	bestScore := -1
	bestAunt := -1

auntFinder:
	for auntNumber, auntData := range auntDataMap {
		currentScore := 0
		for attrName, attrValue := range auntToFind {
			foundValue, found := auntData[attrName]
			if !found {
				continue
			}
			if compareValue(attrName, attrValue, foundValue) {
				currentScore++
			} else {
				continue auntFinder
			}
			if currentScore > bestScore {
				bestScore = currentScore
				bestAunt = auntNumber
			}
		}
	}
	return bestAunt, bestAunt != -1
}

func compareValue(field string, found int, data int) bool {
	if field == "cats" || field == "trees" {
		return data > found
	} else if field == "pomeranians" || field == "goldfish" {
		return data < found
	} else {
		return data == found
	}
}

func parseInput() {
	auntDataMap = map[int]map[string]int{}
	for _, line := range strings.Split(input, "\n") {
		var numberStr, attr1, attr2, attr3 string
		var amount1, amount2, amount3 int
		_, err := fmt.Sscanf(line, "Sue %s %s %d, %s %d, %s %d",
			&numberStr, &attr1, &amount1, &attr2, &amount2, &attr3, &amount3)
		if err != nil {
			panic(err)
		}

		number, _ := strconv.Atoi(strings.TrimSuffix(numberStr, ":"))
		attr1 = strings.TrimSuffix(attr1, ":")
		attr2 = strings.TrimSuffix(attr2, ":")
		attr3 = strings.TrimSuffix(attr3, ":")

		auntDataMap[number] = map[string]int{
			attr1: amount1,
			attr2: amount2,
			attr3: amount3,
		}
	}
	return
}
