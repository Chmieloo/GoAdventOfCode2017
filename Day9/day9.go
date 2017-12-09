package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("day9input")
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	var line string

	var solution1 int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	var ignored = regexp.MustCompile(`\!.`)
	var garbage = regexp.MustCompile(`\<([^\>]*)\>`)

	s := ignored.ReplaceAllString(line, ``)
	cleared := garbage.ReplaceAllString(s, ``)

	var level int
	for i := 0; i < len(cleared); i++ {
		if string(cleared[i]) == string('{') {
			level++
			solution1 += level
		} else if string(cleared[i]) == string('}') {
			level--
		}
	}

	var solution2 int
	var opened = false
	for i := 0; i < len(s); i++ {
		if opened {
			if string(s[i]) != string('>') {
				solution2++
			} else {
				opened = false
			}
		} else {
			if string(s[i]) == string('<') {
				opened = true
			}
		}
	}

	fmt.Println(solution1)
	fmt.Println(solution2)
	elapsed := time.Since(start)
	log.Printf("Run time: %s", elapsed)
}
