// maybe if you want to make it less bad
// solve it in ranges instead of individual points.

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type conversions struct {
	toSoil, toFert   [][3]int
	toWater, toLight [][3]int
	toTemp, toHumid  [][3]int
	toLocation       [][3]int
}

func splitItems(inp []string) [3]int {
	var arr [3]int
	for i, item := range inp {
		num, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		arr[i] = num
	}
	return arr
}

func loadInput(path string) ([]int, *conversions) {
	var seeds []int
	conv := &conversions{}
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	section := 0
	for _, line := range lines {
		if line == "" {
			section += 1
			continue
		}
		// skip lines without useful data
		if strings.Contains(line, "-") {
			continue
		}
		var start int
		if section == 0 {
			start = strings.Index(line, ":") + 2
		}
		split := strings.Split(line[start:], " ")

		switch section {
		case 0:
			// load seed values
			for _, str := range split {
				num, _ := strconv.Atoi(str)
				seeds = append(seeds, num)
			}
		// load conversion tables
		case 1:
			conv.toSoil = append(conv.toSoil, splitItems(split))
		case 2:
			conv.toFert = append(conv.toFert, splitItems(split))
		case 3:
			conv.toWater = append(conv.toWater, splitItems(split))
		case 4:
			conv.toLight = append(conv.toLight, splitItems(split))
		case 5:
			conv.toTemp = append(conv.toTemp, splitItems(split))
		case 6:
			conv.toHumid = append(conv.toHumid, splitItems(split))
		case 7:
			conv.toLocation = append(conv.toLocation, splitItems(split))
		}

	}
	return seeds, conv
}

func convertChain(seed int, conv conversions) int {
	result := convertItem(seed, conv.toSoil)
	result = convertItem(result, conv.toFert)
	result = convertItem(result, conv.toWater)
	result = convertItem(result, conv.toLight)
	result = convertItem(result, conv.toTemp)
	result = convertItem(result, conv.toHumid)
	result = convertItem(result, conv.toLocation)

	return result
}

func convertItem(inp int, conv [][3]int) int {
	for _, arr := range conv {
		if inp >= arr[1] && inp <= arr[1]+arr[2]-1 {
			inp = inp - arr[1] + arr[0]
			break
		}
	}
	//didnt change
	return inp
}

func main() {
	data, conversions := loadInput("input.txt")
	fmt.Println(data)
	count := 0
	for i := 0; i < len(data); i++ {
		if i%2 == 1 {
			count += data[i]
		}
	}
	fmt.Println(count)

	fmt.Println(part1(data, conversions))

	part2(data, conversions)
}

func part1(seeds []int, conv *conversions) int {
	minResult := math.MaxInt
	for _, seed := range seeds {
		minResult = min(convertChain(seed, *conv), minResult)
	}

	return minResult
}

// lmao this took 15 minutes (on a pi4b) and didnt even finish.
// one of the results was small so i tried that one :)
func part2(seeds []int, conv *conversions) {
	if len(seeds)%2 != 0 {
		return
	}

	var wg sync.WaitGroup

	for i := 0; i < len(seeds); i += 2 {
		wg.Add(1)
		go func(i int) {
			minSeed := math.MaxInt
			for s := seeds[i]; s < (seeds[i] + seeds[i+1] - 1); s++ {
				minSeed = min(convertChain(s, *conv), minSeed)
			}
			fmt.Println(minSeed)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
