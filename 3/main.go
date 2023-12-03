package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bounds struct {
	startx, starty, endx, endy int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var input []byte
	var err error
	// pass an argument to run test input, otherwise normal input
	switch len(os.Args) {
	case 2:
		input, err = os.ReadFile("testinput.txt")
	default:
		input, err = os.ReadFile("input.txt")
	}
	check(err)

	fmt.Println(part1(string(input)))
}

func part1(inp string) int {
	result := 0
	numEnded, inNumber := false, false
	var buff strings.Builder
	input := strings.Split(inp, "\n")
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] >= '0' && input[i][j] <= '9' {
				buff.WriteByte(input[i][j])
				inNumber = true
				if j == len(input[i])-1 {
					numEnded = true
				}
			} else if inNumber {
				inNumber = false
				numEnded = true
			}
			if numEnded {
				numEnded = false
				b := bounds{
					endx:   j,
					endy:   i + 1,
					startx: j - buff.Len() - 1,
					starty: i - 1,
				}
				if checkForPart(input, b) {
					num, err := strconv.Atoi(buff.String())
					check(err)
					result += num
				}
				buff.Reset()
			}
		}
	}
	return result
}

func checkForPart(input []string, b bounds) bool {
	//fix bad bounds
	if b.startx < 0 {
		b.startx = 0
	}
	if b.starty < 0 {
		b.starty = 0
	}
	if b.endx >= len(input[0]) {
		b.endx = len(input[0]) - 1
	}
	if b.endy >= len(input)-1 {
		b.endy = len(input) - 2
	}

	for i := b.starty; i <= b.endy; i++ {
		for j := b.startx; j <= b.endx; j++ {
			if (input[i][j] < '0' || input[i][j] > '9') && input[i][j] != '.' {
				return true
			}
		}
	}
	return false
}
