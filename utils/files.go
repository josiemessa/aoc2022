package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SliceAtoi(input []string) []int {
	var result []int
	for i, s := range input {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		x, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Could not parse line %d: %q", i, s)
		}
		result = append(result, x)
	}
	return result
}

func ReadFileAsLines(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer f.Close()
	fmt.Println(f.Name())

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("scanner error", err)
	}
	return lines
}
