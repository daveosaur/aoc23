package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	input := `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	answer := solve(input, -1, 0, RIGHT)
	fmt.Println(answer)
}
