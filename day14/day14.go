package main

import (
	"aoc2022/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	lines := utils.ReadFileAsLines("/home/josie/personal/aoc2022/day14/input")
	grid := parseInput(lines)
	var (
		maxLeft  = point{math.MaxInt, 0}
		maxRight = point{0, 0}
		maxDown  = point{0, 0}
	)
	for p := range grid {
		if p.x < maxLeft.x {
			maxLeft = p
		} else if p.x > maxRight.x {
			maxRight = p
		}
		if p.y > maxDown.y {
			maxDown = p
		}
	}

	//var result int
	//for simulateSandp1(grid, maxLeft, maxRight, maxDown) {
	//	result++
	//}
	//fmt.Println("Part 1: ", result)

	//start := len(grid)
	//var p2result int
	for simulateSandp2(grid, maxDown) {
		//p2result++
	}
	p2Result := 0
	for _, sand := range grid {
		if sand {
			p2Result++
		}
	}
	printGrid(grid, maxLeft, maxRight, maxDown)
	fmt.Println("Part 2: ", p2Result)
}

func printGrid(grid map[point]bool, maxLeft point, maxRight point, maxDown point) {
	fmt.Printf("%d -> %d", maxLeft.x-10, maxRight.x+10)
	for j := 0; j < maxDown.y+2; j++ {
		fmt.Printf("\n%d ", j)
		for i := maxLeft.x - 10; i < maxRight.x+10; i++ {
			if sand, ok := grid[point{i, j}]; ok {
				if sand {
					fmt.Print("O ")
				} else {
					fmt.Print("# ")
				}
				delete(grid, point{i, j})
			} else {
				fmt.Print(". ")
			}
		}
	}
}

func simulateSandp1(grid map[point]bool, maxLeft point, maxRight point, maxDown point) bool {
	// sand always starts at 500,0
	sand := point{500, 0}

	// y increasing means that the point is moving down, x increasing means that the point is going right
	y := 0
	for {
		// This is a lookahead, we only let the sand "move" to the new position once we've checked it's not blocked
		y++
		newSand := point{sand.x, y}
		if newSand.x > maxRight.x || newSand.x < maxLeft.x || newSand.y > maxDown.y {
			return false
		}
		if _, ok := grid[newSand]; !ok {
			// nothing here, sand continues to move
			sand = newSand
			continue
		}
		// something here, sand needs to see if it can get around
		// one step down and to the left
		newSand = point{sand.x - 1, y}
		if _, ok := grid[newSand]; !ok {
			// nothing here, sand continues to move
			sand = newSand
			continue
		}
		// something to the left, one step down and to the right instead
		newSand = point{sand.x + 1, y}
		if _, ok := grid[newSand]; !ok {
			// nothing here, sand continues to move
			sand = newSand
			continue
		}

		// sand must be blocked, cannot move down, left or right. It settles
		grid[sand] = true
		return true
	}
}

func simulateSandp2(grid map[point]bool, maxDown point) bool {
	// sand always starts at 500,0
	sand := point{500, 0}

	// y increasing means that the point is moving down, x increasing means that the point is going right
	y := 0
	for y < maxDown.y+1 {
		y++
		newSand := point{sand.x, y}

		if _, ok := grid[newSand]; !ok {
			// nothing here, sand continues to move
			sand = newSand
			continue
		}

		// something here, sand needs to see if it can get around
		// one step down and to the left
		newSand = point{sand.x - 1, y}
		if _, ok := grid[newSand]; !ok {
			// nothing here, sand continues to move
			sand = newSand
			continue
		}
		// something to the left, one step down and to the right instead
		newSand = point{sand.x + 1, y}
		if _, ok := grid[newSand]; !ok {
			// nothing here, sand continues to move
			sand = newSand
			continue
		}

		// sand must be blocked, cannot move down, left or right. It settles
		grid[sand] = true
		if sand.x == 500 && sand.y == 0 {
			return false
		}
		return true
	}
	// if it gets here, there's always a rock blocking it at it's looking at the floor
	grid[sand] = true
	return true
}

func parseInput(lines []string) map[point]bool {
	grid := make(map[point]bool)
	for _, line := range lines {
		// 498,4 -> 498,6 -> 496,6
		coords := strings.Split(line, " -> ")
		var prevPoint point
		for i, coord := range coords {
			xStr, yStr, _ := strings.Cut(coord, ",")
			x, _ := strconv.Atoi(xStr)
			y, _ := strconv.Atoi(yStr)
			p := point{x, y}
			grid[p] = false

			if i == 0 {
				prevPoint = p
				continue
			}
			if prevPoint.x < p.x {
				for j := prevPoint.x; j < p.x; j++ {
					q := point{j, prevPoint.y}
					grid[q] = false
				}
			}
			if prevPoint.x > p.x {
				for j := prevPoint.x; j > p.x; j-- {
					q := point{j, prevPoint.y}
					grid[q] = false
				}
			}

			if prevPoint.y < p.y {
				for j := prevPoint.y; j < p.y; j++ {
					q := point{prevPoint.x, j}
					grid[q] = false
				}
			}

			if prevPoint.y > p.y {
				for j := prevPoint.y; j > p.y; j-- {
					q := point{prevPoint.x, j}
					grid[q] = false
				}
			}

			prevPoint = p
		}
	}
	return grid
}
