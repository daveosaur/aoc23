package main

import (
	"fmt"
	"os"
	"strings"
)

type beam struct {
	_x, _y int
	dir    direction
	stop   bool
}
type direction int

type tile struct {
	energized             bool
	up, down, left, right bool
}

const (
	UP direction = iota
	LEFT
	DOWN
	RIGHT
)

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println(solve(string(input[:len(input)-1]), -1, 0, RIGHT))
	fmt.Println(part2(string(input[:len(input)-1])))
}

func solve(inp string, startX, startY int, dir direction) int {
	result := 0
	var beams []*beam
	split := strings.Split(inp, "\n")

	// energized grid
	grid := make([][]tile, len(split))
	for i := range grid {
		grid[i] = make([]tile, len(split[0]))
	}

	//starter beam
	beams = append(beams, &beam{
		_x:  startX,
		_y:  startY,
		dir: dir,
	})

	for steps := 0; steps < 800; steps++ {
		for _, b := range beams {
			if b.stop {
				continue
			}

			switch b.dir {
			case UP:
				b._y--
			case LEFT:
				b._x--
			case DOWN:
				b._y++
			case RIGHT:
				b._x++
			}
			// oob check
			if b._x < 0 || b._y < 0 || b._x >= len(split[0]) || b._y >= len(split) {
				b.stop = true
				continue
			}
			grid[b._y][b._x].energized = true
			switch b.dir {
			case UP:
				if grid[b._y][b._x].up {
					b.stop = true
					continue
				}
				grid[b._y][b._x].up = true
			case LEFT:
				if grid[b._y][b._x].left {
					b.stop = true
					continue
				}
				grid[b._y][b._x].left = true
			case DOWN:
				if grid[b._y][b._x].down {
					b.stop = true
					continue
				}
				grid[b._y][b._x].down = true
			case RIGHT:
				if grid[b._y][b._x].right {
					b.stop = true
					continue
				}
				grid[b._y][b._x].right = true
			}

			//interact with mirrors/splitters
			switch split[b._y][b._x] {
			case '\\':
				switch b.dir {
				case UP:
					b.dir = LEFT
				case DOWN:
					b.dir = RIGHT
				case LEFT:
					b.dir = UP
				case RIGHT:
					b.dir = DOWN
				}
			case '/':
				switch b.dir {
				case UP:
					b.dir = RIGHT
				case DOWN:
					b.dir = LEFT
				case LEFT:
					b.dir = DOWN
				case RIGHT:
					b.dir = UP
				}
			case '-':
				if b.dir == UP || b.dir == DOWN {
					b.dir = LEFT
					newBeam := &beam{
						_x:   b._x,
						_y:   b._y,
						dir:  RIGHT,
						stop: false,
					}
					beams = append(beams, newBeam)

				}
			case '|':
				if b.dir == LEFT || b.dir == RIGHT {
					b.dir = UP
					newBeam := &beam{
						_x:   b._x,
						_y:   b._y,
						dir:  DOWN,
						stop: false,
					}
					beams = append(beams, newBeam)
				}
			}

		}
	}
	for _, line := range grid {
		for _, c := range line {
			if c.energized {
				result++
			}
		}
	}

	return result
}

func part2(inp string) int {
	result := 0
	// hardcoded 100 for input. didnt feel like testing
	for i := 0; i < 100; i++ {
		result = max(solve(inp, -1, i, RIGHT), result)
	}
	for i := 0; i < 100; i++ {
		result = max(solve(inp, 100, i, LEFT), result)
	}
	for i := 0; i < 100; i++ {
		result = max(solve(inp, i, -1, DOWN), result)
	}
	for i := 0; i < 100; i++ {
		result = max(solve(inp, i, 100, UP), result)
	}

	return result
}
