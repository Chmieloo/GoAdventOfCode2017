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

	file, err := os.Open("day8input")
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	var instructions []string
	var line string

	var commands []string
	var conditions []string
	var registers []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		instructions = strings.Split(line, "if")
		command := strings.Trim(instructions[0], " ")
		condition := strings.Trim(instructions[1], " ")

		commands = append(commands, command)
		conditions = append(conditions, condition)

		parts := strings.Split(command, " ")
		register := strings.Trim(parts[0], " ")

		registers = appendIfMissing(registers, register)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(registers)
	elapsed := time.Since(start)
	log.Printf("Run time: %s", elapsed)
}

func appendIfMissing(slice []string, newString string) []string {
	for _, elem := range slice {
		if elem == newString {
			return slice
		}
	}
	return append(slice, newString)
}
