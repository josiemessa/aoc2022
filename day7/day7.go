package main

import (
	"aoc2022/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

const (
	cmdPrompt = "$ "
	cdCmd     = "cd "
	upDir     = ".."
)

type fileNode struct {
	name     string
	parent   *fileNode
	children []*fileNode
	size     int
}

func (f *fileNode) String() string {
	return f.printFileNode("", 0)
}

func (f *fileNode) printFileNode(path string, padding int) string {
	var b strings.Builder
	newPath := fmt.Sprintf("%s/%s", path, f.name)

	fmt.Fprintf(&b, "%*s %s (size=%d)\n", padding, "-", newPath, f.size)

	padding += 2
	for _, child := range f.children {
		b.WriteString(child.printFileNode(newPath, padding))
	}
	return b.String()
}

var totalSize int
var dirSizes = make([]int, 0)

func main() {
	lines := utils.ReadFileAsLines("D:\\dev\\src\\github.com\\josiemessa\\aoc2022\\day7\\input")
	root := parseInput(lines)
	fmt.Println(root)

	const (
		totalDiskSpace = 70000000
		requiredSpace  = 30000000
	)

	unusedSpace := totalDiskSpace - root.size
	requiredDelete := requiredSpace - unusedSpace

	dirSize(root, requiredDelete)
	sort.Ints(dirSizes)

	fmt.Println("Part 1:", totalSize)
	fmt.Println("Part 2:", dirSizes[0])
}

func dirSize(f *fileNode, requiredDelete int) {
	if f.size <= 100000 {
		fmt.Println(f.name, f.size)
		totalSize += f.size
	}

	if f.size >= requiredDelete {
		dirSizes = append(dirSizes, f.size)
	}

	for _, child := range f.children {
		dirSize(child, requiredDelete)
	}
}

func parseInput(lines []string) *fileNode {
	var root *fileNode
	var currentDir *fileNode
	for _, line := range lines {
		// process commands
		if strings.HasPrefix(line, cmdPrompt) {
			_, dirName, ok := strings.Cut(line, cdCmd)
			if !ok {
				if line != "$ ls" {
					log.Fatalln("Unexpected command line", line)
				}
				continue
			}

			// Change directory parsing
			// Navigating back up the tree
			if dirName == upDir || (dirName == "/" && root != nil) {
				currentDir.parent.children = append(currentDir.parent.children, currentDir)
				currentDir.parent.size += currentDir.size
				if dirName == upDir {
					currentDir = currentDir.parent
				} else {
					currentDir = root
				}
				continue
			}
			// navigating down the tree
			f := &fileNode{
				name:   dirName,
				parent: currentDir,
			}
			if dirName == "/" && root == nil {
				root = f
			}
			currentDir = f
			continue
		}

		// process ls output
		if strings.HasPrefix(line, "dir") {
			continue
		}
		size, _, ok := strings.Cut(line, " ")
		if !ok {
			log.Fatalln("Unexpected line", line)
		}
		sizeInt, err := strconv.Atoi(size)
		if err != nil {
			log.Fatalln("Unexpected size from line", line)
		}
		currentDir.size += sizeInt
	}
	// do one last roll up of the current directory, and cry for my lost morning
	for currentDir.parent != nil {
		currentDir.parent.children = append(currentDir.parent.children, currentDir)
		currentDir.parent.size += currentDir.size
		currentDir = currentDir.parent
	}
	return root
}
