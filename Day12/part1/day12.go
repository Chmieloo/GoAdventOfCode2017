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

var visited []int

func main() {
	start := time.Now()

	file, err := os.Open("day12input")
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	var line string
	var array = make(map[int][]int)

	// Read input file contents
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		data := strings.Split(line, " <-> ")
		children := strings.Split(data[1], ",")
		for _, elem := range children {
			index, _ := strconv.Atoi(strings.Trim(data[0], " "))
			child, _ := strconv.Atoi(strings.Trim(elem, " "))
			array[index] = append(array[index], child)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	dfs(0, array)
	solution1 := len(visited)

	fmt.Println(solution1)

	elapsed := time.Since(start)
	log.Printf("Run time: %s", elapsed)
}

func dfs(nodeIndex int, array map[int][]int) {
	children := array[nodeIndex]
	for _, elem := range children {
		if !intInSlice(elem, visited) {
			visited = append(visited, elem)
			dfs(elem, array)
		}
	}
}

func intInSlice(a int, slice []int) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
