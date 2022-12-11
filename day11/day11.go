package main

import (
	"aoc2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items []int
	//itemFactors []map[int][2]int
	operation func(int) int
	//factorOperation factorOperation
	testDivisible  int
	trueRecipient  int
	falseRecipient int
	itemsInspected int
}

type factorOperation func(int, int, int) (int, int)

func (m *monkey) takeTurn(monkeys []*monkey) {
	m.itemsInspected += len(m.items)
	// PART 1
	//for _, item := range m.items {
	//	item = m.operation(item)
	//	//item = item / 3  // PART 1 ONLY
	//	if item%m.testDivisible == 0 {
	//		monkeys[m.trueRecipient].items = append(monkeys[m.trueRecipient].items, item)
	//	} else {
	//		monkeys[m.falseRecipient].items = append(monkeys[m.falseRecipient].items, item)
	//	}
	//}
	//m.items = []int{}

	for _, item := range m.items {
		// only perform the operation on the remainder and then fold into the other values
		newItem := m.operation(item)
		newItem = newItem % horribleBigNumer

		if newItem%m.testDivisible == 0 {
			monkeys[m.trueRecipient].items = append(monkeys[m.trueRecipient].items, newItem)
		} else {
			monkeys[m.falseRecipient].items = append(monkeys[m.falseRecipient].items, newItem)
		}
	}
	m.items = []int{}
}

func adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func multiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func squared(x int) int {
	return x * x
}

//func adder(x int) factorOperation {
//	return func(q int, r int, _ int) (int, int) {
//		return q, r + x
//	}
//}
//
//func multiplier(x int) factorOperation {
//	return func(q int, r int, _ int) (int, int) {
//		return q * x, r * x
//	}
//}
//
//func squared(q int, r int, d int) (int, int) {
//	return (d * q * q) + 2*q*r, r ^ 2
//}

var horribleBigNumer int

func main() {
	lines := utils.ReadFileAsLines("D:\\dev\\src\\github.com\\josiemessa\\aoc2022\\day11\\input")

	monkeys := parseInput(lines)

	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			m.takeTurn(monkeys)
		}
	}

	inspected := make([]int, len(monkeys))
	for i, m := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times\n", i, m.itemsInspected)
		//fmt.Printf("Monkey %d has items: ", i)
		//fmt.Print("\n")
		inspected[i] = m.itemsInspected
	}
	sort.Ints(inspected)

	fmt.Println("Part 1:", inspected[len(monkeys)-1]*inspected[len(monkeys)-2])
}

func parseInput(lines []string) []*monkey {
	var (
		currentMonkey = &monkey{}
		monkeys       = make([]*monkey, 0)
		divisors      = []int{}
	)

	for _, line := range lines {
		if line == "" {
			monkeys = append(monkeys, currentMonkey)
			currentMonkey = &monkey{}
			continue
		}

		// Parse starting items
		_, items, ok := strings.Cut(line, "Starting items: ")
		if ok {
			split := strings.Split(items, ", ")
			currentMonkey.items = utils.SliceAtoi(split)
			continue
		}

		// Parse operation
		_, operation, ok := strings.Cut(line, "Operation: new = old ")
		if ok {
			if operation[0] == '+' {
				// ADDITION
				val := strings.TrimPrefix(operation, "+ ")
				if val == "old" {
					currentMonkey.operation = multiplier(2)
				} else {
					valInt, err := strconv.Atoi(val)
					if err != nil {
						panic(line + err.Error())
					}
					currentMonkey.operation = adder(valInt)
				}
			} else {
				// MULTIPLICATION
				val := strings.TrimPrefix(operation, "* ")
				if val == "old" {
					currentMonkey.operation = squared
				} else {
					valInt, err := strconv.Atoi(val)
					if err != nil {
						panic(line + err.Error())
					}
					currentMonkey.operation = multiplier(valInt)
				}
			}
			continue
		}

		// Parse Test
		_, divisible, ok := strings.Cut(line, "Test: divisible by ")
		if ok {
			valInt, err := strconv.Atoi(divisible)
			if err != nil {
				panic(line + err.Error())
			}
			currentMonkey.testDivisible = valInt
			divisors = append(divisors, valInt)
			continue
		}

		// Parse Test results
		_, ifTrue, ok := strings.Cut(line, "If true: throw to monkey ")
		if ok {
			valInt, err := strconv.Atoi(ifTrue)
			if err != nil {
				panic(line + err.Error())
			}
			currentMonkey.trueRecipient = valInt
			continue
		}
		_, ifFalse, ok := strings.Cut(line, "If false: throw to monkey ")
		if ok {
			valInt, err := strconv.Atoi(ifFalse)
			if err != nil {
				panic(line + err.Error())
			}
			currentMonkey.falseRecipient = valInt
			continue
		}
	}
	monkeys = append(monkeys, currentMonkey)

	horribleBigNumer = 1
	for _, divisor := range divisors {
		horribleBigNumer *= divisor
	}

	// madness starts here
	//for i, m := range monkeys {
	//	m.itemFactors = make([]map[int][2]int, len(m.items))
	//	for j, item := range m.items {
	//		m.itemFactors[j] = make(map[int][2]int)
	//		for _, d := range divisors {
	//			quotient := int(item / d)
	//			remainder := item % d
	//			m.itemFactors[j][d] = [2]int{quotient, remainder}
	//		}
	//	}
	//	monkeys[i] = m
	//	monkeys[i].items = nil
	//}

	return monkeys
}
