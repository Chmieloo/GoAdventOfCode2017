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

	var sum int64

	for i := 0; i < len(instructions); i++ {
		nextIndex := (i + 1) % len(instructions)
		current, errCurrent := strconv.ParseInt(instructions[i], 10, 64)
		next, errNext := strconv.ParseInt(instructions[nextIndex], 10, 64)

		if errCurrent != nil {
			log.Fatal(errCurrent)
		}

		if errNext != nil {
			log.Fatal(errNext)
		}

		if next == current {
			sum += current
		}
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func abs(n int) int64 {
	if n < 0 {
		return int64(-1 * n)
	}

	return int64(n)
}
