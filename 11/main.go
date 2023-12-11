package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type position struct {
	x, y int
}

func main() {
	input, _ := os.ReadFile("input.txt")

	// p1 := solve(string(input))
	// fmt.Println(p1)
	p2 := solveP2(string(input))
	fmt.Println(p2)

}

func solve(inp string) int {
	result := 0
	lines := strings.Split(inp, "\n")
	lines = lines[0 : len(lines)-1]
	var emptyLine strings.Builder
	for range lines[0] {
		emptyLine.WriteRune('.')
	}
	emptyRows := []int{}
	emptyCols := []int{}
	for i, line := range lines {
		if line == "" {
			break
		}
		if !strings.Contains(line, "#") {
			emptyRows = append(emptyRows, i)
		}
	}
	//checking columns harder
	for i := 0; i < len(lines[0]); i++ {
		empty := true
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == '#' {
				empty = false
			}
		}
		if empty {
			emptyCols = append(emptyCols, i)
		}
	}
	// expand the stars
	for i, row := range emptyRows {
		lines = slices.Insert(lines, row+i, emptyLine.String())
	}
	for j, col := range emptyCols {
		for i := range lines {
			lines[i] = lines[i][0:col+j] + "." + lines[i][col+j:]
		}
	}

	// now actually count the stars and find distances

	stars := make([]position, 0, 100)
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i] == "" {
				continue
			}
			if lines[i][j] == '#' {
				stars = append(stars, position{x: j, y: i})
			}
		}
	}
	for i := 0; i < len(stars); i++ {
		for j := i + 1; j < len(stars); j++ {
			dist := math.Abs(float64(stars[i].x-stars[j].x)) + math.Abs(float64(stars[i].y-stars[j].y))
			result += int(dist)
			// fmt.Println("star ", i, " star ", j, ": ", dist)
		}
	}
	// for _, line := range lines {
	// 	fmt.Println(line)
	// }
	fmt.Println(emptyRows, emptyCols)

	return result
}

func solveP2(inp string) int {
	result := 0
	lines := strings.Split(inp, "\n")
	lines = lines[0 : len(lines)-1]
	emptyRows := []int{}
	emptyCols := []int{}
	for i, line := range lines {
		if line == "" {
			break
		}
		if !strings.Contains(line, "#") {
			emptyRows = append(emptyRows, i)
		}
	}
	//checking columns harder
	for i := 0; i < len(lines[0]); i++ {
		empty := true
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == '#' {
				empty = false
			}
		}
		if empty {
			emptyCols = append(emptyCols, i)
		}
	}

	// now actually count the stars and find distances

	stars := make([]position, 0, 100)
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i] == "" {
				continue
			}
			if lines[i][j] == '#' {
				stars = append(stars, position{x: j, y: i})
			}
		}
	}
	for i := 0; i < len(stars); i++ {
		for j := i + 1; j < len(stars); j++ {
			dist := math.Abs(float64(stars[i].x-stars[j].x)) + math.Abs(float64(stars[i].y-stars[j].y))
			for _, col := range emptyCols {
				lowx := min(stars[i].x, stars[j].x)
				hix := max(stars[i].x, stars[j].x)
				if col > lowx && col < hix {
					dist += 999999
				}
			}
			for _, row := range emptyRows {
				lowy := min(stars[i].y, stars[j].y)
				hiy := max(stars[i].y, stars[j].y)
				if row > lowy && row < hiy {
					dist += 999999
				}

			}
			result += int(dist)
			// fmt.Println("star ", i, " star ", j, ": ", dist)
		}
	}
	// for _, line := range lines {
	// 	fmt.Println(line)
	// }
	fmt.Println(emptyRows, emptyCols)

	return result
}
