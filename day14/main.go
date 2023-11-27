package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Reindeer struct {
	name           string
	speed          int
	endurance      int
	rest           int
	enduranceTimer int
	restTimer      int
	distance       int
}

var time int
var herd = []Reindeer{}

func main() {
	createHerd()
	for i := 0; i < 2503; i++ {
		elapseSecond()
	}

	bestReindeer := Reindeer{}
	for _, reindeer := range herd {
		if reindeer.distance > bestReindeer.distance {
			bestReindeer = reindeer
		}
	}

	fmt.Printf("Best reindeer is %s with a distance of %d", bestReindeer.name, bestReindeer.distance)
}

func elapseSecond() {
	time++
	for rId := range herd {
		elapseSecondForReindeer(&herd[rId])
	}
}

func elapseSecondForReindeer(reindeer *Reindeer) {
	if reindeer.enduranceTimer > 0 {
		reindeer.enduranceTimer--
		reindeer.distance += reindeer.speed
		if reindeer.enduranceTimer == 0 {
			reindeer.restTimer = reindeer.rest
		}
	} else if reindeer.restTimer > 0 {
		reindeer.restTimer--
		if reindeer.restTimer == 0 {
			reindeer.enduranceTimer = reindeer.endurance
		}
	} else {
		reindeer.enduranceTimer = reindeer.endurance
	}
}

func createHerd() {
	for _, line := range strings.Split(input, "\n") {
		words := strings.Split(line, " ")
		name := words[0]
		speed, _ := strconv.Atoi(words[3])
		endurance, _ := strconv.Atoi(words[6])
		rest, _ := strconv.Atoi(words[13])
		reindeer := Reindeer{name, speed, endurance, rest, 0, 0, 0}
		herd = append(herd, reindeer)
	}
}
