package main

import (
	"aoc2022/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadFileAsLines("input")
	caloriesByElf := make([]int, 0)
	currentCalories := 0
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			caloriesByElf = append(caloriesByElf, currentCalories)
			currentCalories = 0
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Expected calory value, found %s\n", line)
		}
		currentCalories += calories
	}

	// PART 1
	maxCalories := 0
	for _, calories := range caloriesByElf {
		if calories > maxCalories {
			maxCalories = calories
		}
	}
	fmt.Println("PART 1:", maxCalories)

	// PART 2
	sort.Ints(caloriesByElf)
	length := len(caloriesByElf)
	if length < 3 {
		log.Fatalf("Expected more than three elves, found", length)
	}
	fmt.Println("PART2", caloriesByElf[length-1]+caloriesByElf[length-2]+caloriesByElf[length-3])
}
