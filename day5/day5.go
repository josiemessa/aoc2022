package main

import (
	"aoc2022/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type crateStack [9][]rune

func (c crateStack) String() (s string) {
	for i, runes := range c {
		s += fmt.Sprintf("%d: ", i)
		for _, r := range runes {
			s += fmt.Sprint(string(r))
		}
		s += "\n"
	}
	return s
}

type instruction struct {
	quantity int
	src      int
	dest     int
}

// expected to not have to return cratestack here but for some reason it doesn't make these changes in place and it's too early for me to remember why
func (in *instruction) apply(crates crateStack) crateStack {
	fmt.Printf("Move %d from %d to %d\n", in.quantity, in.src, in.dest)
	for i := 0; i < in.quantity; i++ {
		// put the first elements of the source onto the first element of dest
		// dance for the reallocation gods, the GC needs its sacrifices
		crates[in.dest] = append([]rune{crates[in.src][0]}, crates[in.dest]...)
		// remove the first element of src
		crates[in.src] = crates[in.src][1:]
	}
	fmt.Println(crates.String())
	return crates
}

func (in *instruction) apply2(crates crateStack) crateStack {
	fmt.Printf("Move %d from %d to %d\n", in.quantity, in.src, in.dest)

	top := make([]rune, in.quantity)
	if copied := copy(top, crates[in.src][0:in.quantity]); copied != in.quantity {
		log.Fatalf("copy failed. Copied %d wanted %d", copied, in.quantity)
	}
	crates[in.dest] = append(top, crates[in.dest]...)
	crates[in.src] = crates[in.src][in.quantity:]
	fmt.Println(crates.String())
	return crates
}

func main() {
	lines := utils.ReadFileAsLines("input")
	crates, instructions := parseInput(lines)
	fmt.Println(crates.String())

	// part 1
	//for _, in := range instructions {
	//	crates = in.apply(crates)
	//}
	//
	//fmt.Print("Part 1: ")
	//for _, stack := range crates {
	//	fmt.Print(string(stack[0]))
	//}
	//fmt.Println()

	// part 2
	for _, in := range instructions {
		crates = in.apply2(crates)
	}
	fmt.Print("Part 2: ")
	for _, stack := range crates {
		fmt.Print(string(stack[0]))
	}
	fmt.Println()

}

func parseInput(lines []string) (crateStack, [501]*instruction) {
	//crates := make([][]rune, 9)
	// parse crate stacks
	//for i, line := range lines {
	//	// there are 8 lines of crate stacks in input
	//	if i > 7 {
	//		break
	//	}
	//	for j, char := range line {
	//		if char != '[' {
	//			continue
	//		}
	//		stack := int(math.Floor(float64(j / 4)))               // 4 characters per stack, "[A] "
	//		crates[stack] = append(crates[stack], rune(line[j+1])) // this goes in first so 0 will always be the top of the stack
	//	}
	//}
	crates := [9][]rune{
		{'D', 'T', 'W', 'N', 'L'},
		{'H', 'P', 'C'},
		{'J', 'M', 'G', 'D', 'N', 'H', 'P', 'W'},
		{'L', 'Q', 'T', 'N', 'S', 'W', 'C'},
		{'N', 'C', 'H', 'P'},
		{'B', 'Q', 'W', 'M', 'D', 'N', 'H', 'T'},
		{'L', 'S', 'G', 'J', 'R', 'B', 'M'},
		{'T', 'R', 'B', 'V', 'G', 'W', 'N', 'Z'},
		{'L', 'P', 'N', 'D', 'G', 'W'},
	}

	instructions := [501]*instruction{}
	for i := 0; i < 501; i++ {
		line := lines[i+10]
		split := strings.Split(strings.Trim(line, "move "), " from ")
		quant, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatalln("Cannot parse quantity", err)
		}
		srcDest := strings.Split(split[1], " to ")
		srcDestInt := utils.SliceAtoi(srcDest)
		instructions[i] = &instruction{
			quantity: quant,
			src:      srcDestInt[0] - 1,
			dest:     srcDestInt[1] - 1,
		}
	}
	return crates, instructions
}
