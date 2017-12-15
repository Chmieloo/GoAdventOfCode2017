package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	var prevA int64 = 722
	var prevB int64 = 354

	var aFactor int64 = 16807
	var bFactor int64 = 48271

	var divider int64 = 2147483647

	var solution1 int

	for i := 0; i < 40000000; i++ {
		nextA := (prevA * aFactor) % divider
		nextB := (prevB * bFactor) % divider

		last16A := nextA & 65535
		last16B := nextB & 65535

		if last16A == last16B {
			solution1++
		}

		prevA = nextA
		prevB = nextB
	}

	fmt.Println(solution1)

	elapsed := time.Since(start)
	log.Printf("Run time: %s", elapsed)
}

func leftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}
