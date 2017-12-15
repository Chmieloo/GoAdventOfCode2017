package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	var prevA int64 = 722
	var prevB int64 = 354

	var aFactor int64 = 16807
	var bFactor int64 = 48271

	var divider int64 = 2147483647

	var solution int

	for i := 0; i < 5000000; i++ {
		nextA := getNextA(prevA, aFactor, divider)
		nextB := getNextB(prevB, bFactor, divider)

		last16A := nextA & 65535
		last16B := nextB & 65535

		if last16A == last16B {
			solution++
		}

		prevA = nextA
		prevB = nextB
	}

	fmt.Println(solution)

	elapsed := time.Since(start)
	log.Printf("Run time: %s", elapsed)
}

func getNextA(prevA int64, factor int64, divider int64) int64 {
	var nextA int64
	for ok := true; ok; ok = (prevA&0x3 != 0) {
		nextA = (prevA * factor) % divider
		prevA = nextA
	}

	return nextA
}

func getNextB(prevB int64, factor int64, divider int64) int64 {
	var nextB int64
	for ok := true; ok; ok = (prevB&0x7 != 0) {
		nextB = (prevB * factor) % divider
		prevB = nextB
	}

	return nextB
}
