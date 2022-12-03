package main

import (
	"aoc2022/utils"
	"fmt"
	"log"
	"strings"
	"unicode"
)

func main() {
	lines := utils.ReadFileAsLines("../input")

	var itemPriorities int
	var badgePriorities int

	for i, line := range lines {
		// PART 1
		length := len(line)
		if length%2 != 0 {
			log.Fatalln("Expected length to be even, was ", length)
		}
		for _, item := range line[:length/2] {
			if ok := strings.ContainsRune(line[length/2:length], item); ok {
				// found item, calculate priority
				itemPriorities += calculatePriority(item)
				break
			}
		}

		// PART 2
		if i%3 != 0 {
			continue
		}
		for _, item := range lines[i] {
			if ok1 := strings.ContainsRune(lines[i+1], item); ok1 {
				if ok2 := strings.ContainsRune(lines[i+2], item); ok2 {
					badgePriorities += calculatePriority(item)
					break
				}
			}
		}
	}

	fmt.Println("PART 1:", itemPriorities)
	fmt.Println("PART 2:", badgePriorities)
}

func calculatePriority(item rune) int {
	if unicode.IsLower(item) {
		return int(item) - int('a') + 1
	}
	return int(item) - int('A') + 27
}
