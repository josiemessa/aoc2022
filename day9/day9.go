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

func (p point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func right(p point) point {
	return point{p.x + 1, p.y}
}
func left(p point) point {
	return point{p.x - 1, p.y}
}
func up(p point) point {
	return point{p.x, p.y + 1}
}
func down(p point) point {
	return point{p.x, p.y - 1}
}

var (
	head    = point{0, 0}
	tailPos = make(map[string]struct{})
)

func main() {
	lines := utils.ReadFileAsLines("D:\\dev\\src\\github.com\\josiemessa\\aoc2022\\day9\\input")

	tailPos[point{0, 0}.String()] = struct{}{}
	rope := make([]point, 9)

	for _, line := range lines {
		direction, size, _ := strings.Cut(line, " ")

		var move func(point) point
		switch direction {
		case "R":
			move = right
		case "L":
			move = left
		case "U":
			move = up
		case "D":
			move = down
		}

		sizeInt, _ := strconv.Atoi(size)

		rope = simulate(sizeInt, move, rope)

		//fmt.Println("Head:", head, " - Tail:", tail)
	}

	fmt.Println("Solution:", len(tailPos))
}

func simulate(sizeInt int, move func(point) point, rope []point) []point {
	for i := 0; i < sizeInt; i++ {
		head = move(head)

		prevKnot := head
		for j, knot := range rope {
			// move tail
			// diagonal case first
			if prevKnot.x != knot.x && prevKnot.y != knot.y {
				// check if move needed
				if math.Abs(float64(prevKnot.x-knot.x)) > 1 || math.Abs(float64(prevKnot.y-knot.y)) > 1 {
					if prevKnot.x > knot.x {
						knot = right(knot)
					} else {
						knot = left(knot)
					}
					if prevKnot.y > knot.y {
						knot = up(knot)
					} else {
						knot = down(knot)
					}
				}
			} else {
				// cardinal
				if prevKnot.x-knot.x > 1 {
					knot = right(knot)
				} else if knot.x-prevKnot.x > 1 {
					knot = left(knot)
				} else if prevKnot.y-knot.y > 1 {
					knot = up(knot)
				} else if knot.y-prevKnot.y > 1 {
					knot = down(knot)
				}
			}

			prevKnot = knot
			rope[j] = knot
		}

		tail := rope[len(rope)-1]

		tailPos[tail.String()] = struct{}{}
	}
	return rope
}
