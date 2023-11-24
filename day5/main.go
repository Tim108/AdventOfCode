package main

import (
	_ "embed"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	niceWordcount := 0
	for _, word := range strings.Split(input, "\n") {
		if property4(word) && property5(word) {
			niceWordcount++
		}
	}
	println("Nice words:", niceWordcount)
}

// It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
func property1(word string) (result bool) {
	vowels := "aeiou"
	vowelCount := 0
	for _, char := range word {
		if strings.Contains(vowels, string(char)) {
			vowelCount++
		}
	}
	return vowelCount >= 3
}

// It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
func property2(word string) (result bool) {
	for i, char := range word {
		if i < len(word)-1 {
			nextChar := rune(word[i+1])
			if char == nextChar {
				return true
			}
		}
	}
	return false
}

// It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
func property3(word string) (result bool) {
	bannedSequences := []string{"ab", "cd", "pq", "xy"}
	for i, char := range word {
		if i < len(word)-1 {
			nextChar := rune(word[i+1])
			sequence := string(char) + string(nextChar)
			if slices.Contains(bannedSequences, sequence) {
				return false
			}
		}
	}
	return true
}

// It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
func property4(word string) (result bool) {
	pairs := []string{}

	for i, char := range word {
		if i <= len(word)-2 {
			nextChar := rune(word[i+1])
			pair := string(char) + string(nextChar)

			if len(pairs) > 0 && slices.Contains(pairs[:len(pairs)-1], pair) {
				return true
			} else {
				pairs = append(pairs, pair)
			}
		}
	}
	return false
}

// It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.
func property5(word string) (result bool) {
	for i, char := range word {
		if i <= len(word)-3 {
			nextCharOver := rune(word[i+2])
			if char == nextCharOver {
				return true
			}
		}
	}
	return false
}
