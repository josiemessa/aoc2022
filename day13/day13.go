package main

import (
	"aoc2022/utils"
	"fmt"
	"sort"
	"strings"
)

type packetData struct {
	number int
	list   []*packetData
}

func (p *packetData) String() string {
	var result string
	if p.list != nil {
		result += fmt.Sprint("[")
		for _, data := range p.list {
			result += fmt.Sprintf("%s,", data)
		}
		result += fmt.Sprint("]")
	} else {
		result += fmt.Sprintf("%d", p.number)
	}
	return result
}

type pair struct {
	left  *packetData
	right *packetData
}

func main() {
	lines := utils.ReadFileAsLines("D:\\dev\\src\\github.com\\josiemessa\\aoc2022\\day13\\input")
	pairs := parseInput(lines)
	fmt.Println("Part 1: ", part1(pairs))

	list := make([]*packetData, len(pairs)*2)
	for i, p := range pairs {
		list[i*2] = p.left
		list[i*2+1] = p.right
	}
	fmt.Println("Part 2: ", part2(list))

	//for _, pr := range pairs {
	//	fmt.Println(pr.left)
	//	fmt.Println(pr.right)
	//	fmt.Println()
	//}
}

func part1(pairs []pair) int {
	var result int
	for i, p := range pairs {
		if test(p.left.list, p.right.list) == 1 {
			fmt.Printf("Pair %d in the right order\n", i+1)
			result += i + 1
		}
	}
	return result
}

func part2(list []*packetData) int {
	two := &packetData{list: []*packetData{{number: 2}}}
	six := &packetData{list: []*packetData{{number: 6}}}
	list = append(list, two, six)
	sort.Slice(list, func(i, j int) bool {
		return test(list[i].list, list[j].list) == 1
	})

	var result int
	for i, data := range list {
		if data == two {
			result = i + 1
		}
		if data == six {
			result *= i + 1
		}
	}
	return result
}

// test returns 0 if false, 1 is true, -1 is undetermined
func test(leftList []*packetData, rightList []*packetData) (rightOrder int) {
	for i := 0; i < len(rightList); i++ {
		if i >= len(leftList) {
			return 1
		}

		leftValue := leftList[i]
		rightValue := rightList[i]

		// Both values are integers
		if leftValue.list == nil && rightValue.list == nil {
			if leftValue.number < rightValue.number {
				return 1
			} else if leftValue.number > rightValue.number {
				return 0
			} else {
				continue
			}
		}

		// both values are lists
		if leftValue.list != nil && rightValue.list != nil {
			res := test(leftValue.list, rightValue.list)
			if res == -1 {
				continue
			} else {
				return res
			}
		}

		// mismatching types
		left := leftValue.list
		right := rightValue.list
		if leftValue.list == nil {
			left = []*packetData{{number: leftValue.number}}
		} else {
			right = []*packetData{{number: rightValue.number}}
		}
		res := test(left, right)
		if res == -1 {
			continue
		} else {
			return res
		}
	}

	if len(rightList) == len(leftList) {
		return -1
	} else {
		return 0
	}
}

func parseInput(lines []string) []pair {
	pairs := make([]pair, 0, len(lines)/2)
	currentPair := pair{}
	for i, line := range lines {
		if line == "" {
			pairs = append(pairs, currentPair)
			currentPair = pair{}
			continue
		}

		line = strings.TrimSuffix(strings.TrimPrefix(line, "["), "]")
		p, _ := parseLine(line)
		if i%3 == 0 {
			currentPair.left = p
		} else if i%3 == 1 {
			currentPair.right = p
		}
	}
	pairs = append(pairs, currentPair)

	return pairs
}

func parseLine(line string) (*packetData, int) {
	p := &packetData{
		list: make([]*packetData, 0),
	}
	i := 0
	for i < len(line) {
		switch line[i] {
		case '[':
			res, n := parseLine(line[i+1:])
			i += n
			p.list = append(p.list, res)
		case ']':
			return p, i + 1 // return how many characters were read
		case ',':
			break
		default:
			// 10s exist, curse AoC and their families with off by 1 errors
			if line[i] == '1' && line[i+1] == '0' {
				p.list = append(p.list, &packetData{number: 10})
			} else {
				p.list = append(p.list, &packetData{number: int(line[i] - 48)})
			}
		}
		i++
	}
	return p, len(line) - 1
}

/* pair should look like
[1,1,3,1,1]
[1,1,5,1,1]
pair {
	left:  packetData{ [ {1}, {1}, {3}, {1}, {1} ] }
	right: packetData{ [ {1}, {1}, {3}, {1}, {1} ] }
}

[[1],[2,3,4]]
[[1],4]

pair {
	left:  packetData{ [ { [ {1} ] }, { [ {2},{3},{4} ] } ] }
	right: packetData{ [ { [ {1} ] }, {4} ] }
}

[[[]]]
[[]]

pair {
	left:  packetData{ [ packetData{ [ packetData{ [] } ] } ] ] }
	right: packetData{ [ packetData{ [] } ] ] }
}

*/
