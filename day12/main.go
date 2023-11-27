package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
	data := parseInput(input)
	integers := findIntegers(valuesOfMap(data))

	sum := 0
	for _, integer := range integers {
		sum += integer
	}
	println("Total of all integers:", strconv.Itoa(sum))
}

func parseInput(jsonStr string) (result map[string]interface{}) {
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		panic(err)
	}
	return result
}

func findIntegers(data []interface{}) (integers []int) {
	for _, value := range data {
		switch value.(type) {
		case float64:
			integers = append(integers, int(value.(float64)))
		case map[string]interface{}:
			mapValues := valuesOfMap(value.(map[string]interface{}))
			integers = append(integers, findIntegers(mapValues)...)
		case []interface{}:
			integers = append(integers, findIntegers(value.([]interface{}))...)
		case string:
		default:
			fmt.Printf("type not caught for %v is type %T \n", value, value)
		}
	}
	return
}

func valuesOfMap(input map[string]interface{}) (output []interface{}) {
	for _, mapValue := range input {
		// If any of the values is "red" we ignore all values of the map for part 2
		if mapValue == "red" {
			return make([]interface{}, 0)
		}
		output = append(output, mapValue)
	}
	return
}
