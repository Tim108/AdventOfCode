package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

var input = "ckczppom"

func main() {
	for i := 0; i < 10000000; i++ {
		str := input + strconv.Itoa(i)
		hash := createHash(str)

		if hash[:6] == "000000" {
			println("The answer is", strconv.Itoa(i), "with hash", hash)
			return
		}

	}
	println("Found nothing :(")
}

func createHash(key string) (hash string) {
	data := []byte(key)
	checksum := md5.Sum(data)
	hash = hex.EncodeToString(checksum[:])
	return
}
