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
	file, err := os.Open("day1input")
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	var instructions []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = strings.Split(line, "")
	}

	middle := len(instructions) / 2

	var sum int64
	var sumMiddle int64

	for i := 0; i < len(instructions); i++ {
		nextIndex := (i + 1) % len(instructions)
		nextMiddleIndex := (i + middle) % len(instructions)
		current, errCurrent := strconv.ParseInt(instructions[i], 10, 64)
		next, errNext := strconv.ParseInt(instructions[nextIndex], 10, 64)
		nextMiddle, errNextMiddle := strconv.ParseInt(instructions[nextMiddleIndex], 10, 64)

		if errCurrent != nil {
			log.Fatal(errCurrent)
		}

		if errNext != nil {
			log.Fatal(errNext)
		}

		if errNextMiddle != nil {
			log.Fatal(errNextMiddle)
		}

		if next == current {
			sum += current
		}

		if nextMiddle == current {
			sumMiddle += current
		}
	}

	fmt.Println("Part 1 solution: " + strconv.Itoa(int(sum)))
	fmt.Println("Part 2 solution: " + strconv.Itoa(int(sumMiddle)))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
