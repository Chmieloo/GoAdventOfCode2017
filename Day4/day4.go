package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("day4input")
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	var instructions []string
	var validPasswords int
	var validPasswordWithoutAnagrams int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = strings.Split(line, " ")

		lineLength := len(instructions)

		m := make(map[string]int)
		for i := 0; i < lineLength; i++ {
			m[instructions[i]] = 1
		}

		uniquesLength := len(m)

		if uniquesLength == lineLength {
			validPasswords++
		}

		noAnagrams := true
	Loop:
		for i := 0; i < len(instructions)-1; i++ {
			for j := i + 1; j < len(instructions); j++ {
				if checkAnagrams(instructions[i], instructions[j]) {
					noAnagrams = false
					break Loop
				}
			}
		}

		if noAnagrams {
			validPasswordWithoutAnagrams++
		}
	}

	fmt.Println("Part 1 solution: " + strconv.Itoa(validPasswords))
	fmt.Println("Part 2 solution: " + strconv.Itoa(validPasswordWithoutAnagrams))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("Run time: %s", elapsed)
}

func checkAnagrams(string1 string, string2 string) bool {
	map1 := make(map[string]int)
	map2 := make(map[string]int)

	split1 := strings.Split(string1, "")
	split2 := strings.Split(string2, "")

	for i := 0; i < len(string1); i++ {
		map1[split1[i]] = strings.Count(string1, split1[i])
	}

	for i := 0; i < len(string2); i++ {
		map2[split2[i]] = strings.Count(string2, split2[i])
	}

	// m1 and m2 are the maps we want to compare
	eq := reflect.DeepEqual(map1, map2)
	if eq {
		// there is an anagram
		return true
	}

	return false
}
