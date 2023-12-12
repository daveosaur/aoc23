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
	enclosed = 0
)

type position struct {
	_x, _y int
	dir    direction
}

func main() {
	input, _ := os.ReadFile("input.txt")

	// fmt.Println(solve(string(input)))
	solve(string(input))
}

func solve(inp string) int {
	// result := 0
	var pos position
	// get starting dimensions/positions
	grid := strings.Split(inp, "\n")
	pos._x = strings.Index(inp, "S") % (len(grid))
	pos._y = strings.Index(inp, "S") / (len(grid))
	// position is grid[y][x]
	part2 := make([][]rune, 140)
	for i, line := range grid {
		for _, c := range line {
			if line == "" {
				break
			}
			part2[i] = append(part2[i], c)
		}
	}

	// pos.dir = UP
	// pos._y--
	pos.dir = DOWN
	pos._y++
	// pos.dir = RIGHT
	// pos._x++

	steps := 1
	for grid[pos._y][pos._x] != 'S' {
		switch grid[pos._y][pos._x] {
		case 'F':
			switch pos.dir {
			case UP:
				pos.dir = RIGHT
			default:
				pos.dir = DOWN
			}
		case '7':
			switch pos.dir {
			case UP:
				pos.dir = LEFT
			default:
				pos.dir = DOWN
			}
		case 'L':
			switch pos.dir {
			case DOWN:
				pos.dir = RIGHT
			default:
				pos.dir = UP
			}
		case 'J':
			switch pos.dir {
			case DOWN:
				pos.dir = LEFT
			default:
				pos.dir = UP
			}
		}
		switch pos.dir {
		case UP:
			pos._y--
		case DOWN:
			pos._y++
		case LEFT:
			pos._x--
		case RIGHT:
			pos._x++
		}
		steps++
	}
	for _, line := range part2 {
		fmt.Println(string(line))
	}

	//shot in the fuckin dark i guess
	for i := 1; i < len(part2)-1; i++ {
		for j := 1; j < len(part2[0])-1; j++ {
			findEnclosed(part2, j, i)
		}
	}
	killEs(part2, 2, 2)
	for _, line := range part2 {
		fmt.Println(string(line))
		for _, c := range line {
			if c == 'E' {
				enclosed++
			}

		}
	}
	fmt.Println(enclosed)

	return steps / 2
}

func killEs(inp [][]rune, _x, _y int) {
	if _x < 0 || _x > 139 || _y < 0 || _y > 139 {
		return
	}
	for inp[_y][_x] == 'E' {
		inp[_y][_x] = '.'
		killEs(inp, _x-1, _y)
		killEs(inp, _x+1, _y)
		killEs(inp, _x, _y-1)
		killEs(inp, _x, _y+1)
	}
}

func findEnclosed(inp [][]rune, _x, _y int) bool {
	if _x < 0 || _x > 139 || _y < 0 || _y > 139 {
		return false
	}
	if inp[_y][_x] == 'O' {
		return true
	}
	if inp[_y][_x] == 'E' {
		return false
	} else {
		inp[_y][_x] = 'E'
		if findEnclosed(inp, _x-1, _y) &&
			findEnclosed(inp, _x+1, _y) &&
			findEnclosed(inp, _x, _y-1) &&
			findEnclosed(inp, _x, _y+1) {
			return true
		}
	}
	return false
}
