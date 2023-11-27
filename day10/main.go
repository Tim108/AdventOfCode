package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	value := "1113222113"
	for i := 0; i < 50; i++ {
		value = lookAndSee(value)
	}
	fmt.Printf("%d\n", len(value))
}

func lookAndSee(input string) string {
	outputBuilder := strings.Builder{}
	var group []rune

	for _, char := range input {
		if len(group) == 0 || group[len(group)-1] == char {
			group = append(group, char)
		} else {
			outputBuilder.WriteString(strconv.Itoa(len(group)))
			outputBuilder.WriteRune(group[0])
			group = []rune{char}
		}
	}
	outputBuilder.WriteString(strconv.Itoa(len(group)))
	outputBuilder.WriteRune(group[0])

	return outputBuilder.String()
}
