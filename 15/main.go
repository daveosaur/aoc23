package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type lens struct {
	label string
	pow   int
}

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println(solve(string(input)))
	fmt.Println(solveP2(string(input)))

}

func solve(inp string) int {
	result := 0
	split := strings.Split(inp, ",")
	for _, token := range split {
		if len(token) < 2 {
			break
		}
		num := 0
		for _, c := range token {
			if c == '\n' {
				break
			}
			num = ((num + int(c)) * 17) % 256
		}
		result += num
	}

	return result
}

func solveP2(inp string) int {
	result := 0
	split := strings.Split(inp, ",")
	boxes := make([][]lens, 256)

	for _, token := range split {
	loop:
		for i, c := range token {
			switch c {
			case '=':
				label := token[0:i]
				boxnum := solve(label)
				num, _ := strconv.Atoi(token[i+1:])
				for j, box := range boxes[boxnum] {
					if box.label == label {
						boxes[boxnum][j].pow = num
						continue loop
					}
				}
				l := lens{
					label: label,
					pow:   num,
				}
				boxes[boxnum] = append(boxes[boxnum], l)
			case '-':
				label := token[0:i]
				boxnum := solve(label)
				for j, box := range boxes[boxnum] {
					if box.label == label {
						boxes[boxnum] = slices.Delete(boxes[boxnum], j, j+1)
					}
				}
			}
		}
	}
	//count it up
	for boxcount, box := range boxes {
		for lenscount, lens := range box {
			result += (boxcount + 1) * (lenscount + 1) * lens.pow
		}
	}

	return result
}
