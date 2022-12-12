package main

import (
	"aoc2022/utils"
	"fmt"
	"math"
)

type point struct {
	i int // i is row
	j int // j is column
}

func (p point) String() string {
	return fmt.Sprintf("(%d,%d)", p.i, p.j)
}

var (
	maxRow int
	maxCol int
	grid   [][]rune
)

func main() {
	lines := utils.ReadFileAsLines("/home/josie/personal/aoc2022/day12/input")
	maxRow = len(lines)
	grid = make([][]rune, maxRow)

	for i, line := range lines {
		if maxCol == 0 {
			maxCol = len(line)
		}
		grid[i] = make([]rune, maxCol)
		for j, r := range line {
			grid[i][j] = r
		}
	}
	// exampleInput
	//source := point{0, 0}
	//target := point{2, 5}

	source := point{20, 0}
	target := point{20, 43}
	fmt.Printf("Part 1: %d\n", dijkstra(source, target))

	source = target
	target = point{math.MaxInt, math.MaxInt}
	fmt.Printf("Part 2: %d\n", dijkstra(source, target))
}

func dijkstra(source, target point) int {
	distance := make(map[point]int)
	prev := make(map[point]*point)
	q := make(map[point]struct{})
	for i, row := range grid {
		for j := range row {
			v := point{i, j}
			distance[v] = math.MaxInt
			prev[v] = nil
			q[v] = struct{}{}
		}
	}

	distance[source] = 0

	for len(q) != 0 {
		u := minDistance(q, distance)
		if _, ok := q[u]; !ok {
			break
		}
		if target.i != math.MaxInt && u == target {
			break
		}
		delete(q, u)

		neighbours := calculateNeighboursp2(u)
		if len(neighbours) == 0 {
			break
		}
		for _, neighbour := range neighbours {
			alt := distance[u] + 1
			if alt < distance[neighbour] {
				distance[neighbour] = alt
				prev[neighbour] = &u
			}
		}
	}

	//p := target
	//previous := prev[p]
	//for previous != nil {
	//	fmt.Printf("(%d, %d):%s -> (%d,%d):%s\n", previous.i, previous.j, string(grid[previous.i][previous.j]), p.i, p.j, string(grid[p.i][p.j]))
	//	p = *previous
	//	previous = prev[p]
	//}

	if target.i == math.MaxInt {
		// we need to calculate the target
		lowestElevations := []point{}
		for i, row := range grid {
			for j, height := range row {
				if height == 'a' || height == 'S' {
					lowestElevations = append(lowestElevations, point{i, j})
				}
			}
		}
		min := math.MaxInt
		//newSource := point{math.MaxInt, math.MaxInt}
		for _, p := range lowestElevations {
			if distance[p] < min {
				min = distance[p]
				//newSource = p
			}
		}
		return min
	}

	return distance[target]
}

func minDistance(q map[point]struct{}, distances map[point]int) (result point) {
	min := math.MaxInt
	for p := range q {
		if distances[p] < min {
			min = distances[p]
			result = p
		}
	}
	return
}

func calculateNeighboursp1(p point) []point {
	result := make([]point, 0)
	height := getHeight(p)

	if p.i > 0 {
		n := point{p.i - 1, p.j}
		if getHeight(n) <= height+1 {
			result = append(result, n)
		}
	}
	if p.i < maxRow-1 {
		n := point{p.i + 1, p.j}
		if getHeight(n) <= height+1 {
			result = append(result, n)
		}
	}
	if p.j > 0 {
		n := point{p.i, p.j - 1}
		if getHeight(n) <= height+1 {
			result = append(result, n)
		}
	}
	if p.j < maxCol-1 {
		n := point{p.i, p.j + 1}
		if getHeight(n) <= height+1 {
			result = append(result, n)
		}
	}
	return result
}

func calculateNeighboursp2(p point) []point {
	result := make([]point, 0)
	height := getHeight(p)

	if p.i > 0 {
		n := point{p.i - 1, p.j}
		if height <= getHeight(n)+1 {
			result = append(result, n)
		}
	}
	if p.i < maxRow-1 {
		n := point{p.i + 1, p.j}
		if height <= getHeight(n)+1 {
			result = append(result, n)
		}
	}
	if p.j > 0 {
		n := point{p.i, p.j - 1}
		if height <= getHeight(n)+1 {
			result = append(result, n)
		}
	}
	if p.j < maxCol-1 {
		n := point{p.i, p.j + 1}
		if height <= getHeight(n)+1 {
			result = append(result, n)
		}
	}
	return result
}

func getHeight(p point) rune {
	height := grid[p.i][p.j]
	if height == 'S' {
		height = 'a'
	} else if height == 'E' {
		height = 'z'
	}
	return height
}
