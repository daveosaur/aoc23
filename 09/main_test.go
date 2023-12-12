package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var (
	input = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
)

func TestParse(t *testing.T) {

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		nums := parseLine(line)
		fmt.Println(nums)

		num := processHistory(nums)
		fmt.Println("part 1: ", num)

		start := processHistoryP2(nums)
		fmt.Println("part 2: ", start)

	}
}

func TestPart1(t *testing.T) {
	fmt.Println(solve(input, 1))
}

func TestPart2(t *testing.T) {
	fmt.Println(solve(input, 2))
}

func BenchmarkProcess(b *testing.B) {
	inp, _ := os.ReadFile("input.txt")
	_ = inp

	for n := 0; n < b.N; n++ {
		// solve(input, 2)
		solve(string(input), 1)
	}
	b.ReportAllocs()

}
