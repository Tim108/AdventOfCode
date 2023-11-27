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
	points         int
}

var herd = []Reindeer{}

func main() {
	createHerd()
	for i := 0; i <= 2503; i++ {
		elapseSecond()
		awardPoint()
	}

	bestReindeer := Reindeer{}
	for _, reindeer := range herd {
		if reindeer.points > bestReindeer.points {
			bestReindeer = reindeer
		}
	}

	fmt.Printf("Best reindeer is %s with %d points", bestReindeer.name, bestReindeer.points)
}

func awardPoint() {
	indices := []int{}
	furthestDistance := -1

	for i, reindeer := range herd {
		if reindeer.distance > furthestDistance {
			indices = []int{i}
			furthestDistance = herd[i].distance
		} else if reindeer.distance == furthestDistance {
			indices = append(indices, i)
		}
	}

	for _, i := range indices {
		herd[i].points++
	}
}

func elapseSecond() {
	for rId := range herd {
		herd[rId].elapseSecond()
	}
}

func (reindeer *Reindeer) elapseSecond() {
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
	}
}

func createHerd() {
	for _, line := range strings.Split(input, "\n") {
		words := strings.Split(line, " ")
		name := words[0]
		speed, _ := strconv.Atoi(words[3])
		endurance, _ := strconv.Atoi(words[6])
		rest, _ := strconv.Atoi(words[13])
		reindeer := Reindeer{name, speed, endurance, rest, endurance, 0, 0, 0}
		herd = append(herd, reindeer)
	}
}
