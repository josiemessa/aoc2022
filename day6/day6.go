package main

import (
	"aoc2022/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFileAsLines("input")
	line := lines[0]

	for i, _ := range line {
		//fourChars := map[int32]struct{}{}
		fourteenChars := map[int32]struct{}{}
		//for _, char := range line[i : i+4] {
		//	fourChars[char] = struct{}{}
		//}
		for _, char := range line[i : i+14] {
			fourteenChars[char] = struct{}{}
		}
		//if len(fourChars) == 4 {
		//	fmt.Println("Part 1:", i+4)
		//	return
		//}
		if len(fourteenChars) == 14 {
			fmt.Println("Part 2:", i+14)
			return
		}
	}
}
