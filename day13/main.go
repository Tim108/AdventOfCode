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

func main() {
	preferences := parseInput(input)
	preferences = addSelf(preferences) // part 2

	permutations := permutations([]string{}, preferences)

	bestOrder := []string{}
	maxHappiness := -1
	for _, permutation := range permutations {
		happiness := totalHappiness(permutation, preferences)
		if happiness > maxHappiness {
			maxHappiness = happiness
			bestOrder = permutation
		}
	}
	fmt.Printf("Best order is %v with a total happiness of %d", bestOrder, maxHappiness)
}

func permutations(order []string, preferences map[string]map[string]int) (perms [][]string) {
	if len(order) == len(preferences) {
		return [][]string{order}
	}

	// We know there's preferences from everyone to everyone so no need to check our options here
	for person := range preferences {
		if !slices.Contains(order, person) {
			perms = append(perms, permutations(append(order, person), preferences)...)
		}
	}

	return
}

func totalHappiness(order []string, preferences map[string]map[string]int) (total int) {
	order = append(order, order[0])

	for i := 0; i < len(order)-1; i++ {
		personA := order[i]
		personB := order[i+1]
		total += preferences[personA][personB]
		total += preferences[personB][personA]
	}

	return
}

func parseInput(input string) map[string]map[string]int {
	preferences := map[string]map[string]int{}

	for _, line := range strings.Split(input, "\n") {
		// Parse args
		words := strings.Split(line, " ")
		personA := words[0]
		personB := strings.TrimSuffix(words[10], ".")
		value, _ := strconv.Atoi(words[3])
		if words[2] == "lose" {
			value = value * -1
		}

		// Add to preferences map
		_, ok := preferences[personA]
		if !ok {
			preferences[personA] = map[string]int{}
		}
		preferences[personA][personB] = value
	}

	return preferences
}

func addSelf(preferences map[string]map[string]int) map[string]map[string]int {
	preferences["self"] = map[string]int{}
	for person := range preferences {
		preferences["self"][person] = 0
	}
	return preferences
}
