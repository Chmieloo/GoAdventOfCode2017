package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("day6input")
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	var instructions []string
	var line string
	var array []int

	var solution1 = 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	instructions = strings.Split(line, "\t")
	for _, number := range instructions {
		currentNumber, _ := strconv.Atoi(number)
		array = append(array, currentNumber)
	}

	var knownBlocks []string
	var index int

Loop:
	for true {
		maxValue := array[0]
		maxIndex := 0
		for i := 1; i < len(array); i++ {
			if array[i] > maxValue {
				maxValue = array[i]
				maxIndex = i
			}
		}

		array[maxIndex] = 0
		for j := 1; j <= maxValue; j++ {
			insertIndex := (maxIndex + j) % len(array)
			array[insertIndex]++
		}

		joinedBlocks := arrayToString(array, "")
		if stringInSlice(joinedBlocks, knownBlocks) {
			for i := 0; i < len(knownBlocks); i++ {
				if knownBlocks[i] == joinedBlocks {
					index = i
				}
			}
			break Loop
		} else {
			solution1++
			knownBlocks = append(knownBlocks, joinedBlocks)
		}
	}

	// Increase to substract 1 more because of 0 index
	index++
	solution2 := solution1 - index

	fmt.Println(solution1)
	fmt.Println(solution2)
	elapsed := time.Since(start)
	log.Printf("Run time: %s", elapsed)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
