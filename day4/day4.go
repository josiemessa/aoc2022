package main

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

func main() {
	lines := utils.ReadFileAsLines("input")

	var fullyContainsCount int
	var partiallyContainsCount int
	for _, line := range lines {
		split := strings.Split(line, ",")
		firstRangeStr := strings.Split(split[0], "-")
		secondRangeStr := strings.Split(split[1], "-")
		firstRange := utils.SliceAtoi(firstRangeStr)
		secondRange := utils.SliceAtoi(secondRangeStr)

		if secondRange[0] >= firstRange[0] {
			if secondRange[1] <= firstRange[1] {
				fullyContainsCount++
				continue
			}
			if secondRange[0] <= firstRange[1] {
				partiallyContainsCount++
				continue
			}
		}
		if firstRange[0] >= secondRange[0] {
			if firstRange[1] <= secondRange[1] {
				fullyContainsCount++
				continue
			}
			if firstRange[0] <= secondRange[1] {
				partiallyContainsCount++
				continue
			}
		}
		if secondRange[1] >= firstRange[0] && secondRange[1] <= firstRange[1] {
			partiallyContainsCount++
			continue
		}
		if firstRange[1] >= secondRange[0] && firstRange[1] <= secondRange[1] {
			partiallyContainsCount++
		}
	}

	fmt.Println("Part 1:", fullyContainsCount)
	fmt.Println("Part 2:", partiallyContainsCount+fullyContainsCount)
}
