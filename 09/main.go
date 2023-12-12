package main

import (
	"fmt"
	"os"
	"slices"
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
	histories := make([][]int, 0, 32)

	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		history := parseLine(line)
		histories = append(histories, history)
	}
	// lineBegin := 0
	// for i := 0; i < len(inp); i++ {
	// 	switch {
	// 	case inp[i] == '\n':
	// 		history := parseLine(inp[lineBegin:i])
	// 		histories = append(histories, history)
	// 		lineBegin = i + 1
	// 	}
	// }

	for _, history := range histories {
		switch part {
		case 1:
			result += processHistory(history)
		case 2:
			result += processHistoryP2(history)
		}
	}
	return result
}

func processHistory(inp []int) int {
	subHistories := make([][]int, 0, 12)
	subHistories = append(subHistories, inp)
	for {
		allZero := true
		sub := subHistories[len(subHistories)-1]
		newSub := make([]int, 0, 8)
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
	nums := make([]int, 0, 16)

	// split := strings.Split(inp, " ")
	// for _, s := range split {
	// 	num, _ := strconv.Atoi(s)
	// 	nums = append(nums, num)
	// }
	last := 0
	for i, c := range inp {
		switch {
		case c == ' ':
			num, _ := strconv.Atoi(inp[last:i])
			nums = append(nums, num)
			last = i + 1
		}
	}
	num, _ := strconv.Atoi(inp[last:])
	nums = append(nums, num)

	return nums
}

func processHistoryP2(inp []int) int {
	subHistories := make([][]int, 0, 12)
	subHistories = append(subHistories, inp)
	for {
		allZero := true
		sub := subHistories[len(subHistories)-1]
		newSub := make([]int, 0, 8)
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
		subHistories[i-1] = slices.Insert(subHistories[i-1], 0, newStart)
	}
	return subHistories[0][0]
}
