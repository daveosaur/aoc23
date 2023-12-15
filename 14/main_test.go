package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	input := `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

	result := solve(input)
	if result != 136 {
		t.Errorf("result is %d, should be 136\n", result)
	}
	fmt.Println(result)
}
