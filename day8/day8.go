package main

import (
	"aoc2022/utils"
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
)

type tree struct {
	x       int
	y       int
	height  int
	visible bool
}

func point(row, col int) string {
	return fmt.Sprintf("(%d,%d)", row, col)
}

var (
	grid [][]int
	// visible is a map of a coordinate string (in the format "(x,y)"
	visible     = make(map[string]struct{})
	scenicScore = make(map[string]int)
)

const (
	length = 99
)

func main() {
	lines := utils.ReadFileAsLines("D:\\dev\\src\\github.com\\josiemessa\\aoc2022\\day8\\input")
	grid = make([][]int, length)
	for i, line := range lines {
		grid[i] = make([]int, length)
		for j, r := range line {
			grid[i][j] = int(r - 48)
		}
	}

	// i is always the row
	// j is always the column

	// walk around the outside of the grid, looking at all trees along one axis

	// First walk on the left edge, looking right
	//fmt.Println("\nLEFT ->")
	for i := 0; i < length; i++ {
		// look across each row of trees from left to right
		for j := 0; j < length; j++ {
			row := grid[i][:j+1]
			p := point(i, j)
			// Part 1
			vis, vd := treeStats(row)
			if vis {
				fmt.Printf("%#v - visible - %s\n", row, p)
				visible[p] = struct{}{}
			}

			// Part 2 - this is the equivalent of being at tree (i,j) and looking left towards to the right edge
			if s, ok := scenicScore[p]; !ok {
				scenicScore[p] = vd
			} else {
				scenicScore[p] = s * vd
			}
		}
	}

	//fmt.Println("\n\nBOTTOM ^")
	// now walk on the bottom edge, looking up
	for j := 0; j < length; j++ {
		for i := 0; i < length; i++ {
			// create column from bottom looking up
			col := column(j, length-1, i)
			//fmt.Printf("\n%#v", col)
			//if isLastMaxAndDistinct(col) {
			//	//fmt.Print(" - visible ", point(i, j))
			//	visible[point(i, j)] = struct{}{}
			//}
			p := point(i, j)
			// Part 1
			vis, vd := treeStats(col)
			if vis {
				fmt.Printf("%#v - visible - %s\n", col, p)
				visible[p] = struct{}{}
			}

			// Part 2 - this is the equivalent of being at tree (i,j) and looking left towards to the right edge
			if s, ok := scenicScore[p]; !ok {
				scenicScore[p] = vd
			} else {
				scenicScore[p] = s * vd
			}
		}
	}

	//fmt.Println("\n\n<- RIGHT")
	// now walk on the right edge, looking left
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			reversed := utils.Reverse(grid[i][j:])
			//fmt.Printf("\n%#v", reversed)
			//if isLastMaxAndDistinct(reversed) {
			//	key := point(i, j)
			//	//fmt.Print(" - visible", key)
			//	visible[key] = struct{}{}
			//}
			p := point(i, j)
			// Part 1
			vis, vd := treeStats(reversed)
			if vis {
				fmt.Printf("%#v - visible - %s\n", reversed, p)
				visible[p] = struct{}{}
			}

			// Part 2 - this is the equivalent of being at tree (i,j) and looking left towards to the right edge
			if s, ok := scenicScore[p]; !ok {
				scenicScore[p] = vd
			} else {
				scenicScore[p] = s * vd
			}
		}
	}

	//fmt.Println("\n\nTOP v")
	// now walk on the top edge looking down
	for j := 0; j < length; j++ {
		for i := 0; i < length; i++ {
			// create column from top looking down
			col := column(j, 0, i)
			//fmt.Printf("\n%#v", col)
			//if isLastMaxAndDistinct(col) {
			//	//fmt.Print(" - visible", point(i, j))
			//	visible[point(i, j)] = struct{}{}
			//}
			p := point(i, j)
			// Part 1
			vis, vd := treeStats(col)
			if vis {
				fmt.Printf("%#v - visible - %s\n", col, p)
				visible[p] = struct{}{}
			}

			// Part 2 - this is the equivalent of being at tree (i,j) and looking left towards to the right edge
			if s, ok := scenicScore[p]; !ok {
				scenicScore[p] = vd
			} else {
				scenicScore[p] = s * vd
			}
		}
	}

	max := 0
	for _, score := range scenicScore {
		if score > max {
			max = score
		}
	}

	fmt.Println("\n\nPart 1:", len(visible))
	fmt.Println("\n\nPart 2:", max)
}

func treeStats[T constraints.Ordered](s []T) (bool, int) {
	if len(s) <= 1 {
		return true, 0
	}

	// figure out where the next highest tree is in this row
	treeHeight := s[len(s)-1]
	// iterate backwards through the list starting at the penultimate member
	for k := len(s) - 2; k >= 0; k-- {
		if s[k] >= treeHeight {
			return false, len(s) - 1 - k
		}
	}
	return true, len(s) - 1
}

func column(columnIndex, startIndex, endIndex int) []int {
	length := int(math.Abs(float64(endIndex-startIndex)) + 1)
	c := make([]int, length)
	i := 0
	j := startIndex
	for {
		if i > length-1 {
			break
		}
		c[i] = grid[j][columnIndex]
		if startIndex < endIndex {
			j++
		} else {
			j--
		}
		i++
	}
	return c
}

//// for a 5x5 grid, I want to traverse (x=0,y=0->4), (x=0->4, y=4), (x=4, y=4->0), (x=4->0, y=0) for the outer ring
//t := tree{
//	x: 0,
//	y: i,
//}
//t.height = grid[t.y][t.x] // note these are swapped as the grid is lines first then characters

// checkVisible is readonly
//func (t tree) checkVisible() bool {
//	if v, ok := visible[point(t.x, t.y)]; ok {
//		return v
//	}
//	// on the outer ring
//	if t.x == 0 || t.x == len(grid[t.y]) || t.y == 0 || t.y == len(grid) {
//		return true
//	}
//	// check cardinals
//	// t.x-1, t.y // left
//	// t.x+1, t.y // right
//	// t.x, t.y-1 // top
//	// t.x, t.y+1 // bottom
//
//	// if tree to one direction is not viible
//	if v, ok := visible[point(t.x-1, t.y)]; ok && !v {
//		return false
//	}
//	if v, ok := visible[point(t.x+1, t.y)]; ok && !v {
//		return false
//	}
//	if v, ok := visible[point(t.x, t.y-1)]; ok && !v {
//		return false
//	}
//	if v, ok := visible[point(t.x, t.y+1)]; ok && !v {
//		return false
//	}
//}
