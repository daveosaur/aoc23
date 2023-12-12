package main

import (
	"fmt"
	"testing"
)

var (
	input = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`
)

func TestSolve(t *testing.T) {
	fmt.Println(solve(input))
}
