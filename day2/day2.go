package main

import (
	"aoc2022/utils"
	"fmt"
	"log"
	"strings"
)

// rules is a map that shows which rune beats another rune
var lose = map[rune]rune{
	'A': 'Z', // rock beats scissors
	'B': 'X', // paper beats rock
	'C': 'Y', // scissors beats paper
}
var draw = map[rune]rune{
	'A': 'X',
	'B': 'Y',
	'C': 'Z',
}
var win = map[rune]rune{
	'A': 'Y',
	'B': 'Z',
	'C': 'X',
}

var score = map[rune]int{
	'X': 1,
	'Y': 2,
	'Z': 3,
}

var result = map[rune]int{
	'X': 0,
	'Y': 3,
	'Z': 6,
}

func main() {
	// 1 for rock, 2 for paper, 3 for scissors
	// 0 for loss, 3 for draw, 6 for win
	lines := utils.ReadFileAsLines("input")

	// PART 1
	part1(lines)

	// PART 2
	var runningScore int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) != 3 {
			log.Fatalln("expected three characters per line, found", len(line))
		}
		opp := rune(line[0])
		expResult := rune(line[2])

		runningScore += result[expResult]
		if expResult == 'Z' {
			runningScore += score[win[opp]]
			continue
		}
		if expResult == 'Y' {
			runningScore += score[draw[opp]]
			continue
		}
		if expResult == 'X' {
			runningScore += score[lose[opp]]
			continue
		}
		log.Fatalln("Unexpected result", expResult)
	}

	fmt.Println("PART 2:", runningScore)

}

func part1(lines []string) {
	var runningScore int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) != 3 {
			log.Fatalln("expected three characters per line, found", len(line))
		}
		opp := rune(line[0])
		mine := rune(line[2])

		runningScore += score[mine]
		if lose[opp] == mine {
			// lose
			continue
		}
		if draw[opp] == mine {
			runningScore += 3
			continue
		}
		// win
		runningScore += 6
	}

	fmt.Println("PART 1:", runningScore)
}
