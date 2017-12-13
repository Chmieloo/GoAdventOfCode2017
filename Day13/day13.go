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

	file, err := os.Open("day13input")
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	var line string
	var layers = make(map[int]int)
	var length int

	// Read input file contents
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		parts := strings.Split(line, ": ")
		index, _ := strconv.Atoi(parts[0])
		value, _ := strconv.Atoi(parts[1])

		layers[index] = value
		if index > length {
			length = index
		}
	}

	severity := 0

	for i := 0; i <= length; i++ {
		if _, ok := layers[i]; ok {
			if i%((layers[i]-1)*2) == 0 {
				severity += layers[i] * i
			}
		}
	}

	delay := 0
	scanning := true
	for scanning {
		scanning = false
		for i := 0; i <= length; i++ {
			if _, ok := layers[i]; ok {
				if (i+delay)%((layers[i]-1)*2) == 0 {
					scanning = true
					delay++
					break
				}
			}
		}
	}

	solution1 := severity
	solution2 := delay

	fmt.Println("Solution 1: ", solution1)
	fmt.Println("Solution 2: ", solution2)

	elapsed := time.Since(start)
	log.Printf("Run time: %s", elapsed)
}
