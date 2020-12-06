package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Could not read input file")
	}
	rows := strings.Fields(string(input))

	isTree := func(x, y int) bool {
		row := rows[y]
		return string(row[x%len(row)]) == "#"
	}

	tryPath := func(right, down int) int {
		numTrees := 0
		for x, y := 0, 0; y < len(rows); x, y = x+right, y+down {
			if isTree(x, y) {
				numTrees++
			}
		}
		return numTrees
	}

	// Part 1
	fmt.Println(tryPath(3, 1))

	// Part 2
	fmt.Println(
		tryPath(1, 1) *
			tryPath(3, 1) *
			tryPath(5, 1) *
			tryPath(7, 1) *
			tryPath(1, 2))
}
