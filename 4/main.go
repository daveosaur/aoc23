package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var input []byte
	var err error

	switch len(os.Args) {
	case 2:
		input, err = os.ReadFile("testinput.txt")
	default:
		input, err = os.ReadFile("input.txt")
	}
	check(err)

	fmt.Println(part1(string(input)))
	fmt.Println(part2(string(input)))

}

func part1(inp string) int {
	result := 0
	input := strings.Split(inp, "\n")
	for _, line := range input {
		if len(line) < 1 {
			break
		}
		points := 0

		start := strings.Index(line, ":") + 1
		data := strings.Split(line[start:], "|")
		winning := parseCard(data[0])
		nums := parseCard(data[1])

		for _, num := range nums {
			for _, win := range winning {
				if num == win {
					if points > 0 {
						points *= 2
					} else {
						points = 1
					}
				}
			}
		}
		result += points
	}

	return result
}

func part2(inp string) int {
	result := 0
	input := strings.Split(inp, "\n")
	cardCount := make([]int, len(input)+1)
	for index, line := range input {
		if len(line) < 1 {
			break
		}
		start := strings.Index(line, ":") + 1
		data := strings.Split(line[start:], "|")
		winning := parseCard(data[0])
		nums := parseCard(data[1])

		// one point per match, instead of 1 and then doubling
		points := 0
		for _, num := range nums {
			for _, win := range winning {
				if num == win {
					points += 1
				}
			}
		}
		// append the results to each subsequent cardcount
		for i := 0; i <= cardCount[index]; i++ {
			for j := index + 1; j < index+1+points; j++ {
				cardCount[j] += 1
			}
			result += 1
		}
	}

	return result
}

// parses numbers into the nums slice
func parseCard(inp string) []int {
	var buff strings.Builder
	nums := make([]int, 0, 20)
	for _, c := range inp {
		switch c {
		case ' ':
			num, err := strconv.Atoi(buff.String())
			if err != nil {
				break
			}
			nums = append(nums, num)
			buff.Reset()
		default:
			buff.WriteRune(c)
		}
	}
	num, _ := strconv.Atoi(buff.String())
	nums = append(nums, num)
	return nums
}
