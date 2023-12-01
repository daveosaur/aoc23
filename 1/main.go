package main

import (
	"fmt"
	"os"
	"strings"
)

// lazy error check for file loading
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	testinput, err := os.ReadFile("testinput.txt")
	check(err)
	inp, err := os.ReadFile("input.txt")
	check(err)

	result := part1(string(testinput))
	fmt.Println(result)

	result = part1(string(inp))
	fmt.Println(result)

	result = part2(string(inp))
	fmt.Println(result)

}

func part1(inp string) int {
	result := 0
	lines := strings.Split(inp, "\n")
	for _, line := range lines {
		var first, last byte
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				first = line[i] - '0'
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				last = line[i] - '0'
				break
			}
		}
		// stitch together the digits
		val := int((first * 10) + last)
		result += val
	}

	// now add them
	return result
}

func part2(inp string) int {
	nums := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	result := 0
	lines := strings.Split(inp, "\n")

	for _, line := range lines {
		var first, last byte
	First:
		for i := 0; i < len(line); i++ {
			for index, num := range nums {
				if i < len(num) {
					continue
				}
				if strings.Contains(line[0:i], num) {
					first = byte(index)
					break First
				}
			}
			if line[i] >= '0' && line[i] <= '9' {
				first = line[i] - '0'
				break
			}
		}
	Last:
		for i := len(line) - 1; i >= 0; i-- {
			for index, num := range nums {
				if len(line)-1-i < len(num) {
					continue
				}
				if strings.Contains(line[i:], num) {
					last = byte(index)
					break Last
				}
			}
			if line[i] >= '0' && line[i] <= '9' {
				last = line[i] - '0'
				break
			}
		}
		val := int((first * 10) + last)
		result += val
	}
	return result
}
