package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type hand struct {
	hand     string
	strength handType
	bid      int
}

type handType int

const (
	HIGH handType = iota
	PAIR
	TWO
	THREE
	FULL
	FOUR
	FIVE
)

func parseLine(inp string) hand {
	split := strings.Split(inp, " ")
	bid, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	result := hand{
		hand: split[0],
		bid:  bid,
	}
	cards := [13]int{}
	for _, c := range split[0] {
		switch c {
		case 'A':
			cards[0]++
		case 'K':
			cards[1]++
		case 'Q':
			cards[2]++
		case 'J':
			cards[3]++
		case 'T':
			cards[4]++
		case '9':
			cards[5]++
		case '8':
			cards[6]++
		case '7':
			cards[7]++
		case '6':
			cards[8]++
		case '5':
			cards[9]++
		case '4':
			cards[10]++
		case '3':
			cards[11]++
		case '2':
			cards[12]++
		}
	}
	// find max
	maxCount := 0
	pairs := 0
	for _, num := range cards {
		if num == 2 {
			pairs++
		}
		maxCount = max(maxCount, num)
	}
	// determine hand
	switch {
	case maxCount == 5:
		result.strength = FIVE
	case maxCount == 4:
		result.strength = FOUR
	case maxCount == 3 && pairs > 0:
		result.strength = FULL
	case maxCount == 3:
		result.strength = THREE
	case pairs == 2:
		result.strength = TWO
	case pairs == 1:
		result.strength = PAIR
	default:
		result.strength = HIGH
	}
	return result
}

func parseLine2(inp string) hand {
	split := strings.Split(inp, " ")
	bid, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	result := hand{
		hand: split[0],
		bid:  bid,
	}
	cards := [13]int{}
	for _, c := range split[0] {
		switch c {
		case 'A':
			cards[0]++
		case 'K':
			cards[1]++
		case 'Q':
			cards[2]++
		case 'J':
			cards[3]++
		case 'T':
			cards[4]++
		case '9':
			cards[5]++
		case '8':
			cards[6]++
		case '7':
			cards[7]++
		case '6':
			cards[8]++
		case '5':
			cards[9]++
		case '4':
			cards[10]++
		case '3':
			cards[11]++
		case '2':
			cards[12]++
		}
	}
	// find max
	maxCount := 0
	pairs := 0
	for i, num := range cards {
		if i != 3 {
			if num == 2 {
				pairs++
			}
			maxCount = max(maxCount, num)
		}
	}
	// determine hand
	// jokers are cards[3]
	switch {
	case maxCount+cards[3] == 5:
		result.strength = FIVE
	case maxCount+cards[3] == 4:
		result.strength = FOUR
	case maxCount == 3 && pairs == 1 || (pairs+cards[3] == 3):
		result.strength = FULL
	case maxCount == 3 || (pairs == 1 && cards[3] == 1) || (cards[3] == 2):
		result.strength = THREE
	case pairs+cards[3] == 2:
		result.strength = TWO
	case pairs+cards[3] == 1:
		result.strength = PAIR
	default:
		result.strength = HIGH
	}
	return result
}

func getValue(inp byte, part int) int {
	switch inp {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		switch part {
		case 1:
			return 11
		case 2:
			return 1
		default: // how
			panic("aaa")
		}
	case 'T':
		return 10
	default:
		num, err := strconv.Atoi(string(inp))
		if err != nil {
			panic(err)
		}
		return num
	}
}

// compare1 func. returns -1 for less, 0 for equal, 1 for greater
func compare1(first, second hand) int {
	if first.strength == second.strength {
		for i := 0; i < 5; i++ {
			if first.hand[i] == second.hand[i] {
				continue
			}
			f, s := getValue(first.hand[i], 1), getValue(second.hand[i], 1)
			if f < s {
				return -1
			} else {
				return 1
			}
		}
	}
	if first.strength < second.strength {
		return -1
	}
	return 0
}

func compare2(first, second hand) int {
	if first.strength == second.strength {
		for i := 0; i < 5; i++ {
			if first.hand[i] == second.hand[i] {
				continue
			}
			f, s := getValue(first.hand[i], 2), getValue(second.hand[i], 2)
			if f < s {
				return -1
			} else {
				return 1
			}
		}
	}
	if first.strength < second.strength {
		return -1
	}
	return 1
}

func solve(inp string, part int) int {

	hands := []hand{}
	split := strings.Split(inp, "\n")
	for _, line := range split {
		if len(line) < 1 {
			break
		}
		var hand hand
		switch part {
		case 1:
			hand = parseLine(line)
		case 2:
			hand = parseLine2(line)

		}
		hands = append(hands, hand)
	}

	switch part {
	case 1:
		slices.SortFunc(hands, compare1)
	case 2:
		slices.SortFunc(hands, compare2)
	}

	total := 0
	for i, h := range hands {
		total += (i + 1) * h.bid
	}
	return total
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part 1: ", solve(string(input), 1))
	fmt.Println("part 2: ", solve(string(input), 2))
}
