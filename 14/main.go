package main

import (
	"fmt"
	"os"
	"strings"
)

type direction int

const (
	UP direction = iota
	LEFT
	DOWN
	RIGHT
)

var (
	points = make(map[[2]int]int, 1000)
)

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println(solve(string(input)))
}

func solve(inp string) int {
	grid := make([][]rune, 0, 100)
	lines := strings.Split(inp, "\n")
	// convert input string to 2d slice
	for _, line := range lines {
		if len(line) < 2 {
			break
		}
		arr := make([]rune, 0, 100)
		for _, c := range line {
			arr = append(arr, c)
		}
		grid = append(grid, arr)
	}

	// this is hardcoded for my solution.
	// spin 1000 times to get to a stable period
	count := 1000
	grid = spinny(grid, count)
	goal := 1_000_000_000
	// i manually checked, my input has a period of 72 spins.
	remaining := (goal - count) % 72
	grid = spinny(grid, remaining)

	return countStones(grid)
}

func moveStones(inp [][]rune) [][]rune {
	// count the stones, find where they should go
	for i := 0; i < len(inp); i++ {
		for j := 0; j < len(inp[0]); j++ {
			if inp[i][j] == 'O' {
				inp[i][j] = '.'
				// found stone, find whats above
				for f := i; ; f-- {
					if f < 0 || inp[f][j] == '#' {
						num, ok := points[[2]int{f, j}]
						if ok {
							points[[2]int{f, j}] = num + 1
						} else {
							points[[2]int{f, j}] = 1
						}
						break
					}
				}
			}
		}
	}
	// redraw the stones
	for pos, num := range points {
		for i := 0; i < num; i++ {
			inp[pos[0]+i+1][pos[1]] = 'O'
		}
	}
	return inp

}

func countStones(inp [][]rune) int {
	num := len(inp)
	result := 0

	for i, line := range inp {
		for _, c := range line {
			if c == 'O' {
				result += num - i
			}
		}
	}
	return result
}

func moveStonesDir(inp [][]rune, dir direction) [][]rune {
	clear(points)
	// count the stones, find where they should go
	for i := 0; i < len(inp); i++ {
		for j := 0; j < len(inp[0]); j++ {
			if inp[i][j] == 'O' {
				inp[i][j] = '.'
				// found stone, find whats above
				switch dir {
				case UP:
					for f := i; ; f-- {
						if f < 0 || inp[f][j] == '#' {
							num, ok := points[[2]int{f, j}]
							if ok {
								points[[2]int{f, j}] = num + 1
							} else {
								points[[2]int{f, j}] = 1
							}
							break
						}
					}
				case DOWN:
					for f := i; ; f++ {
						if f >= len(inp) || inp[f][j] == '#' {
							num, ok := points[[2]int{f, j}]
							if ok {
								points[[2]int{f, j}] = num + 1
							} else {
								points[[2]int{f, j}] = 1
							}
							break
						}
					}
				case LEFT:
					for f := j; ; f-- {
						if f < 0 || inp[i][f] == '#' {
							num, ok := points[[2]int{i, f}]
							if ok {
								points[[2]int{i, f}] = num + 1
							} else {
								points[[2]int{i, f}] = 1
							}
							break
						}
					}
				case RIGHT:
					for f := j; ; f++ {
						if f >= len(inp[0]) || inp[i][f] == '#' {
							num, ok := points[[2]int{i, f}]
							if ok {
								points[[2]int{i, f}] = num + 1
							} else {
								points[[2]int{i, f}] = 1
							}
							break
						}
					}
				}
			}
		}
	}
	// redraw the stones
	switch dir {
	case UP:
		for pos, num := range points {
			for i := 0; i < num; i++ {
				inp[pos[0]+i+1][pos[1]] = 'O'
			}
		}
	case DOWN:
		for pos, num := range points {
			for i := 0; i < num; i++ {
				inp[pos[0]-i-1][pos[1]] = 'O'
			}
		}
	case LEFT:
		for pos, num := range points {
			for i := 0; i < num; i++ {
				inp[pos[0]][pos[1]+i+1] = 'O'
			}
		}
	case RIGHT:
		for pos, num := range points {
			for i := 0; i < num; i++ {
				inp[pos[0]][pos[1]-i-1] = 'O'
			}
		}
	}
	return inp
}

func spinny(inp [][]rune, num int) [][]rune {
	for i := 0; i < num; i++ {
		inp = moveStonesDir(inp, UP)
		inp = moveStonesDir(inp, LEFT)
		inp = moveStonesDir(inp, DOWN)
		inp = moveStonesDir(inp, RIGHT)
	}
	return inp
}
