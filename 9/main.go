package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println(solve(string(input), 1))
	fmt.Println(solve(string(input), 2))

}

func solve(inp string, part int) int {
	result := 0
	lines := strings.Split(inp, "\n")
	var histories [][]int

	for _, line := range lines {
		if line == "" {
			break
		}
		history := parseLine(line)
		histories = append(histories, history)
	}
	for _, h := range histories {
		switch part {
		case 1:
			result += processHistory(h)
		case 2:
			result += processHistoryP2(h)
		}
	}
	return result
}

func processHistory(inp []int) int {
	subHistories := [][]int{inp}
	for {
		allZero := true
		sub := subHistories[len(subHistories)-1]
		newSub := []int{}
		for i := 1; i < len(sub); i++ {
			num := sub[i] - sub[i-1]
			newSub = append(newSub, num)
			if num != 0 {
				allZero = false
			}
		}
		subHistories = append(subHistories, newSub)
		if allZero {
			break
		}
	}
	for i := len(subHistories) - 1; i > 0; i-- {
		sub := subHistories[i]
		prevSub := subHistories[i-1]
		newLast := prevSub[len(prevSub)-1] + sub[len(sub)-1]
		subHistories[i-1] = append(subHistories[i-1], newLast)
	}
	return subHistories[0][len(subHistories[0])-1]
}

func parseLine(inp string) []int {
	nums := []int{}
	var buff strings.Builder

	for i, c := range inp {
		switch {
		case i == len(inp)-1:
			buff.WriteRune(c)
			fallthrough
		case c == ' ':
			num, err := strconv.Atoi(buff.String())
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
			buff.Reset()
		default:
			buff.WriteRune(c)
		}
	}

	return nums
}

func processHistoryP2(inp []int) int {
	subHistories := [][]int{inp}
	for {
		allZero := true
		sub := subHistories[len(subHistories)-1]
		newSub := []int{}
		for i := 1; i < len(sub); i++ {
			num := sub[i] - sub[i-1]
			newSub = append(newSub, num)
			if num != 0 {
				allZero = false
			}
		}
		subHistories = append(subHistories, newSub)
		if allZero {
			break
		}
	}
	for i := len(subHistories) - 1; i > 0; i-- {
		sub := subHistories[i]
		prevSub := subHistories[i-1]
		newStart := prevSub[0] - sub[0]
		subHistories[i-1] = append([]int{newStart}, subHistories[i-1]...)
	}
	return subHistories[0][0]
}
