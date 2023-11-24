package main

import (
	_ "embed"
	"slices"
	"strings"
)

//go:embed input.txt
var input string
var vowels = "aeiou"
var bannedSequences = []string{"ab", "cd", "pq", "xy"}

func main() {
	niceWordcount := 0
	for _, word := range strings.Split(input, "\n") {
		if niceWord(word) {
			niceWordcount++
		}
	}
	println("Nice words:", niceWordcount)
}

func niceWord(word string) (result bool) {
	vowelCount := 0
	repeatedChar := false

	for i, char := range word {
		if strings.Contains(vowels, string(char)) {
			vowelCount++
		}
		if i < len(word)-1 {
			nextChar := rune(word[i+1])
			if char == nextChar {
				repeatedChar = true
			}
			sequence := string(char) + string(nextChar)
			if slices.Contains(bannedSequences, sequence) {
				return false
			}
		}
	}
	return vowelCount >= 3 && repeatedChar
}
