package main

import (
	"slices"
)

func validPassword(password []int) bool {
	return rule1(password) && rule2(password) && rule3(password)
}

// Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
func rule1(password []int) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1]-1 && password[i] == password[i+2]-2 {
			return true
		}
	}
	return false
}

// Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
func rule2(password []int) bool {
	illegalRunes := []int{int('i'), int('o'), int('l')}
	for _, runeInt := range password {
		if slices.Contains(illegalRunes, runeInt) {
			return false
		}
	}
	return true
}

// Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.
func rule3(password []int) bool {
	foundIndices := []int{}
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			// we have found a pair before and that pair didn't start with the previous index
			if len(foundIndices) > 0 && !slices.Contains(foundIndices, i-1) {
				return true
			}
			foundIndices = append(foundIndices, i)
		}
	}
	return false
}
