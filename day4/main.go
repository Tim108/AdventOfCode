package main

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strconv"
)

var input = "ckczppom"
var stringToFind = "000000"
var batchSize = 100000
var batches = 100

func main() {
	resultChannel := make(chan int)
	results := []int{}

	for i := 0; i < batches; i++ {
		go checkMillion(i, resultChannel)
	}

	for i := 0; i < batches; i++ {
		results = append(results, <-resultChannel)
	}

	sort.Ints(results)
	for _, result := range results {
		if result != 0 {
			println("The answer is:", result)
			return
		}
	}

	println("Found nothing :(")
}

func checkMillion(j int, resultChannel chan int) {
	from := j * batchSize
	to := (j + 1) * batchSize

	for i := from; i < to; i++ {
		str := input + strconv.Itoa(i)
		hash := createHash(str)

		if hash[:len(stringToFind)] == stringToFind {
			resultChannel <- i
			return
		}

	}
	resultChannel <- 0
}

func createHash(key string) (hash string) {
	data := []byte(key)
	checksum := md5.Sum(data)
	hash = hex.EncodeToString(checksum[:])
	return
}
