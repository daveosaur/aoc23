package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println(part1(string(input)))
	fmt.Println(part2(string(input)))

}

func part1(inp string) int {
	steps := 0
	split := strings.Split(inp, "\n")
	directions := split[0]

	nodes := parseNodes(split[2:])

	current := "AAA"
	index := 0
	for current != "ZZZ" {
		switch directions[index] {
		case 'L':
			current = nodes[current][0]
		case 'R':
			current = nodes[current][1]
		}
		steps++
		index = (index + 1) % len(directions)

	}

	return steps
}

func parseNodes(inp []string) map[string][2]string {
	result := make(map[string][2]string)
	for _, line := range inp {
		if len(line) < 1 {
			break
		}
		src := line[0:3]
		left := line[7:10]
		right := line[12:15]
		result[src] = [2]string{left, right}
	}

	return result
}

func part2(inp string) int {
	split := strings.Split(inp, "\n")
	directions := split[0]

	nodes := parseNodes(split[2:])

	starters := []string{}
	//find starters
	for _, line := range split[2:] {
		if len(line) < 1 {
			break
		}
		if line[2] == 'A' {
			starters = append(starters, line[0:3])
		}
	}

	cycles := []int{}

	for _, s := range starters {
		cycles = append(cycles, findCycles(s, directions, nodes))
	}

	return findLCM(cycles...)
}

func findCycles(inp, dir string, nodes map[string][2]string) int {
	cur := inp
	index := 0
	step := 0

	for {
		switch dir[index] {
		case 'L':
			cur = nodes[cur][0]
		case 'R':
			cur = nodes[cur][1]
		}
		step++
		index = (index + 1) % len(dir)
		if cur[2] == 'Z' {
			return step
		}
	}
}

func findGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func findLCM(inp ...int) int {
	result := 0
	switch len(inp) {
	case 1, 0:
		return 0
	case 2:
		result = inp[0] * inp[1] / findGCD(inp[0], inp[1])
	default:
		result = inp[0] * inp[1] / findGCD(inp[0], inp[1])
		for i := 2; i < len(inp); i++ {
			result = findLCM(result, inp[i])
		}
	}

	return result
}
