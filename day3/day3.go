package main

import (
	"aoc2022/utils"
	"fmt"
	"log"
	"strings"
	"unicode"
)

func main() {
	lines := utils.ReadFileAsLines("input")

	part1(lines)

	// part 2
	var priorities int
	for i := range lines {
		if i%3 != 0 {
			continue
		}
		for _, item := range lines[i] {
			if ok1 := strings.ContainsRune(lines[i+1], item); ok1 {
				if ok2 := strings.ContainsRune(lines[i+2], item); ok2 {
					priorities += calculatePriority(item)
					break
				}
			}
		}
	}

	fmt.Println("PART 2: ", priorities)
}

func part1(lines []string) {
	var priorities int
	for _, line := range lines {
		length := len(line)
		if length%2 != 0 {
			log.Fatalln("Expected length to be even, was ", length)
		}
		first := line[:length/2]
		second := line[length/2 : length]
		for _, item := range first {
			if ok := strings.ContainsRune(second, item); ok {
				// found item, calculate priority
				priorities += calculatePriority(item)
				break
			}
		}
	}

	fmt.Println("PART 1: ", priorities)
}

func calculatePriority(item rune) int {
	if unicode.IsLower(item) {
		return int(item) - 96
	} else {
		return 26 + int(item) - 64
	}
}
