package main

import (
	"fmt"
	"sort"
	"strconv"
)

var input = "4x23x21\n22x29x19\n11x4x11\n8x10x5\n24x18x16\n11x25x22\n2x13x20\n24x15x14\n14x22x2\n30x7x3\n30x22x25\n29x9x9\n29x29x26\n14x3x16\n1x10x26\n29x2x30\n30x10x25\n10x26x20\n1x2x18\n25x18x5\n21x3x24\n2x5x7\n22x11x21\n11x8x8\n16x18x2\n13x3x8\n1x16x19\n19x16x12\n21x15x1\n29x9x4\n27x10x8\n2x7x27\n2x20x23\n24x11x5\n2x8x27\n10x28x10\n24x11x10\n19x2x12\n27x5x10\n1x14x25\n5x14x30\n15x26x12\n23x20x22\n5x12x1\n9x26x9\n23x25x5\n28x16x19\n17x23x17\n2x27x20\n18x27x13\n16x7x18\n22x7x29\n17x28x6\n9x22x17\n10x5x6\n14x2x12\n25x5x6\n26x9x10\n19x21x6\n19x4x27\n23x16x14\n21x17x29\n24x18x10\n7x19x6\n14x15x10\n9x10x19\n20x18x4\n11x14x8\n30x15x9\n25x12x24\n3x12x5\n12x21x28\n8x23x10\n18x26x8\n17x1x8\n2x29x15\n3x13x28\n23x20x11\n27x25x6\n19x21x3\n30x22x27\n28x24x4\n26x18x21\n11x7x16\n22x27x6\n27x5x26\n4x10x4\n4x2x27\n2x3x26\n26x29x19\n30x26x24\n8x25x12\n16x17x5\n13x2x3\n1x30x22\n20x9x1\n24x26x19\n26x18x1\n18x29x24\n1x6x9\n20x27x2\n3x22x21\n4x16x8\n29x18x16\n7x16x23\n13x8x14\n19x25x10\n23x29x6\n23x21x1\n22x26x10\n14x4x2\n18x29x17\n9x4x18\n7x22x9\n19x5x26\n27x29x19\n7x13x14\n19x10x1\n6x22x3\n12x21x5\n24x20x12\n28x2x11\n16x18x23\n2x13x25\n11x7x17\n27x21x4\n2x10x25\n22x16x17\n23x22x15\n17x13x13\n23x24x26\n27x18x24\n24x7x28\n30x12x15\n14x28x19\n2x15x29\n12x13x5\n17x22x21\n27x10x27\n17x6x25\n22x2x1\n1x10x9\n9x7x2\n30x28x3\n28x11x10\n8x23x15\n23x4x20\n12x5x4\n13x17x14\n28x11x2\n21x11x29\n10x23x22\n27x23x14\n7x15x23\n20x2x13\n8x21x4\n10x20x11\n23x28x11\n21x22x25\n23x11x17\n2x29x10\n28x16x5\n30x26x10\n17x24x16\n26x27x25\n14x13x25\n22x27x5\n24x15x12\n5x21x25\n4x27x1\n25x4x10\n15x13x1\n21x23x7\n8x3x4\n10x5x7\n9x13x30\n2x2x30\n26x4x29\n5x14x14\n2x27x9\n22x16x1\n4x23x5\n13x7x26\n2x12x10\n12x7x22\n26x30x26\n28x16x28\n15x19x11\n4x18x1\n20x14x24\n6x10x22\n9x20x3\n14x9x27\n26x17x9\n10x30x28\n6x3x29\n4x16x28\n8x24x11\n23x10x1\n11x7x7\n29x6x15\n13x25x12\n29x14x3\n26x22x21\n8x3x11\n27x13x25\n27x6x2\n8x11x7\n25x12x9\n24x30x12\n13x1x30\n25x23x16\n9x13x29\n29x26x16\n11x15x9\n11x23x6\n15x27x28\n27x24x21\n6x24x1\n25x25x5\n11x1x26\n21x4x24\n10x5x12\n4x30x13\n24x22x5\n26x7x21\n23x3x17\n22x18x2\n25x1x14\n23x25x30\n8x7x7\n30x19x8\n17x6x15\n2x11x20\n8x3x22\n23x14x26\n8x22x25\n27x1x2\n10x26x2\n28x30x7\n5x30x7\n27x16x30\n28x29x1\n8x25x18\n20x12x29\n9x19x9\n7x25x15\n25x18x18\n11x8x2\n4x20x6\n18x5x20\n2x3x29\n25x26x22\n18x25x26\n9x12x16\n18x7x27\n17x20x9\n6x29x26\n17x7x19\n21x7x5\n29x15x12\n22x4x1\n11x12x11\n26x30x4\n12x24x13\n13x8x3\n26x25x3\n21x26x10\n14x9x26\n20x1x7\n11x12x3\n12x11x4\n11x15x30\n17x6x25\n20x22x3\n1x16x17\n11x5x20\n12x12x7\n2x14x10\n14x27x3\n14x16x18\n21x28x24\n14x20x1\n29x14x1\n10x10x9\n25x23x4\n17x15x14\n9x20x26\n16x2x17\n13x28x25\n16x1x11\n19x16x8\n20x21x2\n27x9x22\n24x18x3\n23x30x6\n4x18x3\n30x15x8\n27x20x19\n28x29x26\n2x21x18\n1x23x30\n1x9x12\n4x11x30\n1x28x4\n17x10x10\n12x14x6\n8x9x24\n8x3x3\n29x8x20\n26x29x2\n29x25x25\n11x17x23\n6x30x21\n13x18x29\n2x10x8\n29x29x27\n27x15x15\n16x17x30\n3x3x22\n21x12x6\n22x1x5\n30x8x20\n6x28x13\n11x2x23\n14x18x27\n6x26x13\n10x24x24\n4x24x6\n20x8x3\n23x11x5\n29x5x24\n14x15x22\n21x17x13\n10x10x8\n1x11x23\n21x19x24\n19x9x13\n21x26x28\n25x11x28\n2x17x1\n18x9x8\n5x21x6\n12x5x2\n23x8x15\n30x16x24\n7x9x27\n16x30x7\n2x21x28\n5x10x6\n8x7x1\n28x13x5\n11x5x14\n26x22x29\n23x15x13\n14x2x16\n22x21x9\n4x20x3\n18x17x19\n12x7x9\n6x12x25\n3x30x27\n8x19x22\n1x9x27\n23x20x12\n14x7x29\n9x12x12\n30x2x6\n15x7x16\n19x13x18\n11x8x13\n16x5x3\n19x26x24\n26x8x21\n21x20x7\n15x1x25\n29x15x21\n22x17x7\n16x17x10\n6x12x24\n8x13x27\n30x25x14\n25x7x10\n15x2x2\n18x15x19\n18x13x24\n19x30x1\n17x1x3\n26x21x15\n10x10x18\n9x16x6\n29x7x30\n11x10x30\n6x11x2\n7x29x23\n13x2x30\n25x27x13\n5x15x21\n4x8x30\n15x27x11\n27x1x6\n2x24x11\n16x20x19\n25x28x20\n6x8x4\n27x16x11\n1x5x27\n12x19x26\n18x24x14\n4x25x17\n24x24x26\n28x3x18\n8x20x28\n22x7x21\n24x5x28\n23x30x29\n25x16x27\n28x10x30\n9x2x4\n30x2x23\n21x9x23\n27x4x26\n2x23x16\n24x26x30\n26x1x30\n10x4x28\n11x29x12\n28x13x30\n24x10x28\n8x12x12\n19x27x11\n11x28x7\n14x6x3\n6x27x5\n6x17x14\n24x24x17\n18x23x14\n17x5x7\n11x4x23\n5x1x17\n26x15x24\n3x9x24\n5x3x15\n5x20x19\n5x21x2\n13x5x30\n19x6x24\n19x17x6\n23x7x13\n28x23x13\n9x1x6\n15x12x16\n21x19x9\n25x5x5\n9x7x9\n6x5x8\n3x11x18\n23x25x11\n25x4x6\n4x27x1\n4x3x3\n30x11x5\n9x17x12\n15x6x24\n10x22x15\n29x27x9\n20x21x11\n18x10x5\n11x2x2\n9x8x8\n1x26x21\n11x11x16\n2x18x30\n29x27x24\n27x8x18\n19x3x17\n30x21x26\n25x13x25\n20x22x1\n10x1x12\n11x17x15\n29x11x30\n17x30x27\n21x22x17\n13x6x22\n22x16x12\n27x18x19\n4x13x6\n27x29x10\n3x23x10\n26x16x24\n18x26x20\n11x28x16\n21x6x15\n9x26x17\n8x15x8\n3x7x10\n2x28x8\n1x2x24\n7x8x9\n19x4x22\n11x20x9\n12x22x16\n26x8x19\n13x28x24\n4x10x16\n12x8x10\n14x24x24\n19x19x28\n29x1x15\n10x5x14\n20x19x23\n10x7x12\n1x7x13\n5x12x13\n25x21x8\n22x28x8\n7x9x4\n3x20x15\n15x27x19\n18x24x12\n16x10x16\n22x19x8\n15x4x3\n9x30x25\n1x1x6\n24x4x25\n13x18x29\n10x2x8\n21x1x17\n29x14x22\n17x29x11\n10x27x16\n25x16x15\n14x2x17\n12x27x3\n14x17x25\n24x4x1\n18x28x18\n9x14x26\n28x24x17\n1x26x12\n2x18x20\n12x19x22\n19x25x20\n5x17x27\n17x29x16\n29x19x11\n16x2x4\n23x24x1\n19x18x3\n28x14x6\n18x5x23\n9x24x12\n15x4x6\n15x7x24\n22x15x8\n22x1x22\n6x4x22\n26x1x30\n8x21x27\n7x1x11\n9x8x18\n20x27x12\n26x23x20\n26x22x30\n24x3x16\n8x24x28\n13x28x5\n4x29x23\n22x5x8\n20x22x3\n9x9x17\n28x3x30\n10x13x10\n10x25x13\n9x20x3\n1x21x25\n24x21x15\n21x5x14\n13x8x20\n29x17x3\n5x17x28\n16x12x7\n23x1x24\n4x24x29\n23x25x14\n8x27x2\n23x11x13\n13x4x5\n24x1x26\n21x1x23\n10x12x12\n21x29x25\n27x25x30\n24x23x4\n1x30x23\n29x28x14\n4x11x30\n9x25x10\n17x11x6\n14x29x30\n23x5x5\n25x18x21\n8x7x1\n27x11x3\n5x10x8\n11x1x11\n16x17x26\n15x22x19\n16x9x6\n18x13x27\n26x4x22\n1x20x21\n6x14x29\n11x7x6\n1x23x7\n12x19x13\n18x21x25\n15x17x20\n23x8x9\n15x9x26\n9x12x9\n12x13x14\n27x26x7\n11x19x22\n16x12x21\n10x30x28\n21x2x7\n12x9x18\n7x17x14\n13x17x17\n3x21x10\n30x9x15\n2x8x15\n15x12x10\n23x26x9\n29x30x10\n30x22x17\n17x26x30\n27x26x20\n17x28x17\n30x12x16\n7x23x15\n30x15x19\n13x19x10\n22x10x4\n17x23x10\n2x28x18\n27x21x28\n24x26x5\n6x23x25\n17x4x16\n14x1x13\n23x21x11\n14x15x30\n26x13x10\n30x19x25\n26x6x26\n9x16x29\n15x2x24\n13x3x20\n23x12x30\n22x23x23\n8x21x2\n18x28x5\n21x27x14\n29x28x23\n12x30x28\n17x16x3\n5x19x11\n28x22x22\n1x4x28\n10x10x14\n18x15x7\n18x11x1\n12x7x16\n10x22x24\n27x25x6\n19x29x25\n10x1x26\n26x27x30\n4x23x19\n24x19x4\n21x11x14\n4x13x27\n9x1x11\n16x20x8\n4x3x11\n1x16x12\n14x6x30\n8x1x10\n11x18x7\n29x28x30\n4x21x8\n3x21x4\n6x1x5\n26x18x3\n28x27x27\n17x3x12\n6x1x22\n23x12x28\n12x13x2\n11x2x13\n7x1x28\n27x6x25\n14x14x3\n14x11x20\n2x27x7\n22x24x23\n7x15x20\n30x6x17\n20x23x25\n18x16x27\n2x9x6\n9x18x19\n20x11x22\n11x16x19\n14x29x23\n14x9x20\n8x10x12\n18x17x6\n28x7x16\n12x19x28\n5x3x16\n1x25x10\n4x14x10\n9x6x3\n15x27x28\n13x26x14\n21x8x25\n29x10x20\n14x26x30\n25x13x28\n1x15x23\n6x20x21\n18x2x1\n22x25x16\n23x25x17\n2x14x21\n14x25x16\n12x17x6\n19x29x15\n25x9x6\n19x17x13\n24x22x5\n19x4x13\n10x18x6\n6x25x6\n23x24x20\n8x22x13\n25x10x29\n5x12x25\n20x5x11\n7x16x29\n29x24x22\n28x20x1\n10x27x10\n6x9x27\n26x15x30\n26x3x19\n20x11x3\n26x1x29\n6x23x4\n6x13x21\n9x23x25\n15x1x10\n29x12x13\n7x8x24\n29x30x27\n3x29x19\n14x16x17\n4x8x27\n26x17x8\n10x27x17\n11x28x17\n17x16x27\n1x8x22\n6x30x16\n7x30x22\n20x12x3\n18x10x2\n20x21x26\n11x1x17\n9x15x15\n19x14x30\n24x22x20\n11x26x23\n14x3x23\n1x28x29\n29x20x4\n1x4x20\n12x26x8\n14x11x14\n14x19x13\n15x13x24\n16x7x26\n11x20x11\n5x24x26\n24x25x7\n21x3x14\n24x29x20\n7x12x1\n16x17x4\n29x16x21\n28x8x17\n11x30x25\n1x26x23\n25x19x28\n30x24x5\n26x29x15\n4x25x23\n14x25x19\n29x10x7\n29x29x28\n19x13x24\n21x28x5\n8x15x24\n1x10x12\n2x26x6\n14x14x4\n10x16x27\n9x17x25\n25x8x7\n1x9x28\n10x8x17\n4x12x1\n17x26x29\n23x12x26\n2x21x22\n18x23x13\n1x14x5\n25x27x26\n4x30x30\n5x13x2\n17x9x6\n28x18x28\n7x30x2\n28x22x17\n14x15x14\n10x14x19\n6x15x22\n27x4x17\n28x21x6\n19x29x26\n6x17x17\n20x13x16\n25x4x1\n2x9x5\n30x3x1\n24x21x2\n14x19x12\n22x5x23\n14x4x21\n10x2x17\n3x14x10\n17x5x3\n22x17x13\n5x19x3\n29x22x6\n12x28x3\n9x21x25\n10x2x14\n13x26x7\n18x23x2\n9x14x17\n21x3x13\n13x23x9\n1x20x4\n11x4x1\n19x5x30\n9x9x29\n26x29x14\n1x4x10\n7x27x30\n8x3x23\n1x27x27\n7x27x27\n1x26x16\n29x16x14\n18x6x12\n24x24x24\n26x2x19\n15x17x4\n11x7x14\n14x19x10\n9x10x1\n14x17x9\n20x19x13\n25x20x8\n24x20x21\n26x30x2\n24x2x10\n28x4x13\n27x17x11\n15x3x8\n11x29x10\n26x15x16\n4x28x22\n7x5x22\n10x28x9\n6x28x13\n10x5x6\n20x12x6\n25x30x30\n17x16x14\n14x20x3\n16x10x8\n9x28x14\n16x12x12\n11x13x25\n21x16x28\n10x3x18\n5x9x20\n17x23x5\n3x13x16\n29x30x17\n2x2x8\n15x8x30\n20x1x16\n23x10x29\n4x5x4\n6x18x12\n26x10x22\n21x10x17\n26x12x29\n7x20x21\n18x9x15\n10x23x20\n20x1x27\n10x10x3\n25x12x23\n30x11x15\n16x22x3\n22x10x11\n15x10x20\n2x20x17\n20x20x1\n24x16x4\n23x27x7\n7x27x22\n24x16x8\n20x11x25\n30x28x11\n21x6x24\n15x2x9\n16x30x24\n21x27x9\n7x19x8\n24x13x28\n12x26x28\n16x21x11\n25x5x13\n23x3x17\n23x1x17\n4x17x18\n17x13x18\n25x12x19\n17x4x19\n4x21x26\n6x28x1\n23x22x15\n6x23x12\n21x17x9\n30x4x23\n2x19x21\n28x24x7\n19x24x14\n13x20x26\n19x24x29\n8x26x3\n16x12x14\n17x4x21\n8x4x20\n13x27x17\n9x21x1\n29x25x6\n7x9x26\n13x25x5\n6x9x21\n12x10x11\n30x28x21\n15x6x2\n8x18x19\n26x20x24\n26x17x14\n27x8x1\n19x19x18\n25x24x27\n14x29x15\n22x26x1\n14x17x9\n2x6x23\n29x7x5\n14x16x19\n14x21x18\n10x15x23\n21x29x14\n20x29x30\n23x11x5\n"

// var input = "23x4x21\n"
var lineSeparator = '\n'
var dimensionSeparator = 'x'

func main() {
	parseInput(input)
	surface := 0
	for _, d := range dimensionsList {
		surface += surfaceForDimensions(d)
	}
	fmt.Printf("total surface: %v\n", surface)

	ribbonLength := 0
	for _, d := range dimensionsList {
		ribbonLength += ribbonLengthForDimensions(d)
	}
	fmt.Printf("total ribbon length: %v\n", ribbonLength)
}

func surfaceForDimensions(dimensions [3]int) (surface int) {
	h, l, w := dimensions[0], dimensions[1], dimensions[2]
	s1, s2, s3 := h*l, l*w, w*h
	sides := [3]int{s1, s2, s3}
	sort.Ints(sides[:])

	surface = s1*2 + s2*2 + s3*2 + sides[0]
	return
}

func ribbonLengthForDimensions(dimensions [3]int) (length int) {
	h, l, w := dimensions[0], dimensions[1], dimensions[2]
	p1, p2, p3 := h*2+l*2, l*2+w*2, w*2+h*2
	perimeters := [3]int{p1, p2, p3}
	sort.Ints(perimeters[:])
	length = perimeters[0] + h*l*w
	return
}

var numberStr = ""
var dimensions = [3]int{}
var dimensionsList = [][3]int{}

func parseInput(input string) {
	for _, char := range input {
		if char == lineSeparator {
			endNumber()
			endDimensions()
		} else if char == dimensionSeparator {
			endNumber()
		} else {
			numberStr = numberStr + string(char)
		}
	}
}

func endNumber() {
	number, err := strconv.Atoi(numberStr)
	numberStr = ""
	if err != nil {
		panic(err)
	}
	addToDimensions(number)
}

func endDimensions() {
	dimensionsList = append(dimensionsList, dimensions)
	dimensions = [3]int{}
}

func addToDimensions(number int) {
	for i, o := range dimensions {
		if o == 0 {
			dimensions[i] = number
			break
		}
	}
}
