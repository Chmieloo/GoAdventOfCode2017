package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("day5input")
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	var array, array2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		number, _ := strconv.Atoi(line)
		array = append(array, number)
		array2 = append(array2, number)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	startIndex := 0
	endIndex := len(array) - 1
	currentIndex := 0
	var sum int

	for currentIndex <= endIndex && currentIndex >= startIndex {
		steps := array[currentIndex]
		oldIndex := currentIndex
		currentIndex = currentIndex + steps
		array[oldIndex]++
		sum++
	}

	fmt.Println(sum)
	elapsed := time.Since(start)
	log.Printf("Run time part 1: %s", elapsed)

	start = time.Now()
	currentIndex = 0
	sum = 0

	for currentIndex <= endIndex && currentIndex >= startIndex {
		steps := array2[currentIndex]
		oldIndex := currentIndex
		currentIndex = currentIndex + steps
		if array2[oldIndex] >= 3 {
			array2[oldIndex]--
		} else {
			array2[oldIndex]++
		}
		sum++
	}

	fmt.Println(sum)
	elapsed = time.Since(start)
	log.Printf("Run time part 2: %s", elapsed)
}
