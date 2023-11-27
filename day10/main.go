package main

import (
	_ "embed"
	"fmt"
	"strconv"
)

func main() {
	value := "1113222113"
	for i := 0; i < 40; i++ {
		value = lookAndSee(value)
	}
	fmt.Printf("%d\n", len(value))
}

func lookAndSee(input string) (output string) {
	group := []string{}

	for _, char := range input {
		if len(group) == 0 || group[len(group)-1] == string(char) {
			group = append(group, string(char))
		} else {
			output += groupToString(group)
			group = []string{string(char)}
		}
	}
	output += groupToString(group)

	return
}

func groupToString(input []string) (output string) {
	if len(input) == 0 {
		return ""
	}
	output += strconv.Itoa(len(input))
	output += input[0]
	return
}
