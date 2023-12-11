package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`
	// _ = input
	// 	inp := `....
	// ...#
	// .#..
	// #...`
	result := solveP2(input)
	// result := solve(inp)
	fmt.Println(result)
}
