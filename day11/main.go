package main

import (
	_ "embed"
	"strings"
)

func main() {
	password := parseInput("cqjxxyzz")
	for i := 0; i < 10000000; i++ {
		candidate := nextIteration(password)
		if validPassword(candidate) {
			println("New password:", createOutput(password), "\nIterations required:", i)
			break
		}
	}
	println("End")
}

func parseInput(input string) (output []int) {
	for _, char := range input {
		output = append(output, int(char))
	}
	return
}

func createOutput(input []int) string {
	pwBuilder := strings.Builder{}
	for _, char := range input {
		pwBuilder.WriteRune(rune(char))
	}
	return pwBuilder.String()
}

func nextIteration(password []int) []int {
	minRune := int('a')
	maxRune := int('z')
	carryOver := 1

	for i := len(password) - 1; i >= 0; i-- {
		password[i] = password[i] + carryOver
		carryOver = 0
		if password[i] > maxRune {
			carryOver = password[i] - maxRune
			password[i] = minRune
		}
	}
	return password
}
