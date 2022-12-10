package main

import (
	"aoc2022/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	noop = "noop"
	addx = "addx "
)

var (
//	instructions = map[string]int{
//		noop: 1,
//		addx: 2,
//	}
)

func main() {
	lines := utils.ReadFileAsLines("D:\\dev\\src\\github.com\\josiemessa\\aoc2022\\day10\\input")

	during, _ := execute(lines)

	// PART 1
	cycles := [...]int{20, 60, 100, 140, 180, 220}
	var strength int
	for _, cycle := range cycles {
		strength += during[cycle] * cycle
		//fmt.Printf("%dth cycle, register: %d, strength %d\n", cycle, register[cycle-1], register[cycle-1]*cycle)
	}
	fmt.Println("Part 1: ", strength)
	fmt.Println("Part 2\n", render(during))

}

func render(register []int) string {
	var b strings.Builder
	for i := 0; i < 240; i++ {
		pixel := i % 40
		if pixel == 0 {
			if _, err := fmt.Fprint(&b, "\n"); err != nil {
				panic(err)
			}
		}
		// during cycle i
		// CRT is going to draw a pixel in position i, need to see if sprite is in position i

		x := register[i]
		if pixel >= x-1 && pixel <= x+1 {
			fmt.Fprint(&b, "# ")
		} else {
			fmt.Fprint(&b, ". ")
		}
	}
	return b.String()
}

// execute runs each command in program and returns a list of the value of X *after* each cycle
func execute(program []string) (during []int, after []int) {
	var (
		X = 1
	)
	during = []int{}
	after = []int{}

	for _, command := range program {
		during = append(during, X)
		after = append(after, X)
		switch {
		case command == noop:
			// ...noop
		case strings.HasPrefix(command, addx):
			val := strings.TrimPrefix(command, addx)
			intVal, err := strconv.Atoi(val)
			if err != nil {
				log.Fatalln(err)
			}

			during = append(during, X)
			X += intVal
			after = append(after, X)
		default:
			panic("unexpected line" + command)
		}
	}
	return during, after
}
