package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	totalRed   = 12
	totalGreen = 13
	totalBlue  = 14
)

type maxColors struct {
	red, green, blue int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var input []byte
	var err error
	//pass any argument to run test input. because i am lazy
	switch len(os.Args) {
	case 1:
		input, err = os.ReadFile("input.txt")
		check(err)
	case 2:
		input, err = os.ReadFile("testinput.txt")

	}

	fmt.Println(part1(string(input)))

	fmt.Println(part2(string(input)))

}

func part1(inp string) int {
	result := 0
	lines := strings.Split(inp, "\n")
Loop:
	for i, line := range lines {
		if len(line) < 1 {
			break
		}
		start := strings.Index(line, ":")
		split := strings.Split(line[start+1:], ";")
		for _, s := range split {
			var a, b, c string
			var x, y, z int
			reader := strings.NewReader(s)
			fmt.Fscanf(reader, " %d %s %d %s %d %s", &x, &a, &y, &b, &z, &c)
			if !(verify(a, x) && verify(b, y) && verify(c, z)) {
				continue Loop
			}
		}
		result += i + 1

	}
	return result
}

func part2(inp string) int {
	result := 0
	lines := strings.Split(inp, "\n")
	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		var maxC maxColors
		start := strings.Index(line, ":")
		split := strings.Split(line[start+1:], ";")
		for _, s := range split {
			var cube1, cube2, cube3 string
			var cube1Val, cube2Val, cube3Val int
			reader := strings.NewReader(s)
			fmt.Fscanf(reader, " %d %s %d %s %d %s",
				&cube1Val, &cube1, &cube2Val, &cube2, &cube3Val, &cube3)
			maxC.parse(cube1, cube1Val)
			maxC.parse(cube2, cube2Val)
			maxC.parse(cube3, cube3Val)
		}
		result += maxC.red * maxC.green * maxC.blue

	}
	return result
}

// updates the new maximum for the parsed color
func (m *maxColors) parse(tok string, count int) {
	switch tok {
	case "red", "red,":
		m.red = max(m.red, count)
	case "green", "green,":
		m.green = max(m.green, count)
	case "blue", "blue,":
		m.blue = max(m.blue, count)

	}
}

// makes sure the parsed color isnt above global max
func verify(tok string, count int) bool {
	switch tok {
	case "red", "red,":
		if count > totalRed {
			return false
		}
	case "green", "green,":
		if count > totalGreen {
			return false
		}
	case "blue", "blue,":
		if count > totalBlue {
			return false
		}
	}
	return true
}
