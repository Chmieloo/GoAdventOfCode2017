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

	file, err := os.Open("day10input")
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	var line string
	var solution1 int

	// Construct a list of ints [0:255]
	list := make([]int, 256)
	for i := 0; i < 256; i++ {
		list[i] = i
	}

	// Read input file contents
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	stringLengths := strings.Split(line, ",")
	intLengths := make([]int, len(stringLengths))
	for i := range intLengths {
		intLengths[i], _ = strconv.Atoi(stringLengths[i])
	}

	var currentPosition int
	var skipSize int
	var listLength = len(list)

	for _, elem := range intLengths {
		var chunk []int
		for i := currentPosition; i < (currentPosition + elem); i++ {
			index := i % listLength
			chunk = append(chunk, list[index])
		}

		for i, j := 0, len(chunk)-1; i < j; i, j = i+1, j-1 {
			chunk[i], chunk[j] = chunk[j], chunk[i]
		}

		var j int
		for i := currentPosition; i < (currentPosition + elem); i++ {
			index := i % listLength
			list[index] = chunk[j]
			j++
		}

		currentPosition += (elem + skipSize) % listLength
		skipSize++
	}

	solution1 = list[0] * list[1]

	fmt.Println(solution1)
	elapsed := time.Since(start)
	log.Printf("Run time: %s", elapsed)
}
