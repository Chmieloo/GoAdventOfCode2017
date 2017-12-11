package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("day11input")
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	var line string

	// Read input file contents
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	var distances []int
	directions := strings.Split(line, ",")
	var x, y, z int

	for _, elem := range directions {
		switch elem {
		case "n":
			y++
			z--
		case "ne":
			x++
			z--
		case "nw":
			x--
			y++
		case "s":
			y--
			z++
		case "se":
			x++
			y--
		case "sw":
			x--
			z++
		}
		distances = append(distances, (abs(x)+abs(y)+abs(z))/2)
	}

	solution1 := (abs(x) + abs(y) + abs(z)) / 2
	solution2 := max(distances)

	fmt.Println(solution1)
	fmt.Println(solution2)

	elapsed := time.Since(start)
	log.Printf("Run time: %s", elapsed)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(array []int) int {
	max := array[0]
	for _, elem := range array {
		if elem >= max {
			max = elem
		}
	}
	return max
}
