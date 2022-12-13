package main

import (
	"aoc2022/utils"
	"fmt"
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

func main() {
	lines := utils.ReadFileAsLines("/home/josie/personal/aoc2022/day13/exampleInput")

	pairs := make([]pair, 0, len(lines)/2)
	currentPair := pair{}
	for i, line := range lines {
		if line == "" {
			pairs = append(pairs, currentPair)
			currentPair = pair{}
			continue
		}

		p, _ := parseLine(line)
		if i%3 == 0 {
			currentPair.left = p
		} else if i%3 == 1 {
			currentPair.right = p
		}
	}
	pairs = append(pairs, currentPair)
	for _, pr := range pairs {
		fmt.Println(pr.left)
		fmt.Println(pr.right)
		fmt.Println()
	}

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
			continue
		default:
			p.list = append(p.list, &packetData{number: int(line[i] - 48)})
		}
		i++
	}
	return p, len(line) - 1
}
