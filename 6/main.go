package main

import (
	"fmt"
	"math"
	"sync"
)

// hardcoded input, w/e
var (
	testRaces = [][2]int{
		{7, 9},
		{15, 40},
		{30, 200},
	}
	races = [][2]int{
		{57, 291},
		{72, 1172},
		{69, 1176},
		{92, 2026},
	}
	part2test = [2]int{
		71530,
		940200,
	}
	part2race = [2]int{
		57726992,
		291117211762026,
	}
)

func main() {
	fmt.Println(part1(races))
	fmt.Println(part2(part2race))

}

func part1(inp [][2]int) int {
	totalWins := []int{}

	for _, race := range inp {
		wins := 0
		for i := 1; i < race[0]; i++ {
			speed := i
			remainingTime := race[0] - i
			dist := speed * remainingTime
			if dist > race[1] {
				wins += 1
			}
		}
		totalWins = append(totalWins, wins)
	}
	result := totalWins[0]
	for i := 1; i < len(totalWins); i++ {
		result *= totalWins[i]
	}
	return result
}

// i threaded this because i thought this would be a brute force problem.
// ...it isnt
func part2(inp [2]int) int {
	wins := 0
	var wg sync.WaitGroup
	results := make(chan int, 5)

	for threads := 0; threads < 4; threads++ {
		wg.Add(1)
		go func(offset int) {
			localwins := 0
			winning := false
			for i := 0 + offset; i < inp[0]; i += 4 {
				speed := i
				remainingTime := inp[0] - i
				dist := remainingTime * speed
				if dist > inp[1] {
					winning = true
					localwins += 1
				} else if winning {
					break
				}
			}
			results <- localwins
			wg.Done()
		}(threads)
	}
	wg.Wait()
	close(results)
	for num := range results {
		wins += num
	}
	return wins
}
