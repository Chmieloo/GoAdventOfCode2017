package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2input")
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	var instructions []string
	var diffSum int
	var evenSum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = strings.Split(line, "\t")

		ary := make([]int, len(instructions))
		for j := range ary {
			ary[j], _ = strconv.Atoi(instructions[j])
		}

		max, min := getMaxMinNumbers(ary)
		difference := max - min
		diffSum += difference
		evenSum += getEvenlyDivisible(ary)
	}

	fmt.Println("Part 1 solution: " + strconv.Itoa(diffSum))
	fmt.Println("Part 2 solution: " + strconv.Itoa(evenSum))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getMaxMinNumbers(inputArray []int) (int, int) {
	var maxNumber = inputArray[0]
	var minNumber = inputArray[0]

	for i := 0; i < len(inputArray); i++ {
		if inputArray[i] >= maxNumber {
			maxNumber = inputArray[i]
		}
		if inputArray[i] <= minNumber {
			minNumber = inputArray[i]
		}
	}

	return maxNumber, minNumber
}

func getEvenlyDivisible(inputArray []int) int {
	var i int
	var j int

	for i = 0; i < len(inputArray)-1; i++ {
		for j = i + 1; j < len(inputArray); j++ {
			current := inputArray[i]
			next := inputArray[j]

			bigger := max(current, next)
			smaller := min(current, next)

			if bigger%smaller == 0 {
				return bigger / smaller
			}
		}
	}
	return 0
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
