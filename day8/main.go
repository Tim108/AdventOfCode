package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	unquoteDiff := 0 // part 1
	quoteDiff := 0   // part 2
	for _, line := range strings.Split(input, "\n") {
		codeLen := len(line)

		unquotedLine, _ := strconv.Unquote(line)
		unquotedLen := len(unquotedLine)
		unquoteDiff += codeLen - unquotedLen

		quotedLine := strconv.Quote(line)
		quotedLen := len(quotedLine)
		quoteDiff += quotedLen - codeLen
	}
	println("Unquote diff ->", strconv.Itoa(unquoteDiff))
	println("Quote diff ->", strconv.Itoa(quoteDiff))
}
