package main

import (
	"fmt"

	"github.com/rdenman/aoc-21/util"
)

func main() {
	sol1 := solution1()
	sol2 := solution2()

	fmt.Printf("Solution #1: %d\n", sol1)
	fmt.Printf("Solution #2: %d\n", sol2)
}

// naive
func solution1() int {
	data := readInputData()

	const days = 80
	for d := 0; d < days; d++ {
		for i := range data {
			if data[i] == 0 {
				data[i] = 6
				data = append(data, 8)
			} else {
				data[i]--
			}
		}
	}

	return len(data)
}

// optimized
func solution2() int {
	data := readInputData()

	const days = 256
	fish := make([]int, 9)
	for _, d := range data {
		fish[d]++
	}

	for d := 0; d < days; d++ {
		fish[7] += fish[0]
		last := fish[0]
		fish = fish[1:]
		fish = append(fish, last)
	}

	tot := 0
	for _, f := range fish {
		tot += f
	}

	return tot
}

func readInputData() []int {
	input, err := util.ReadAndSplitInputData("./input.txt", ",")
	util.CheckError(err)

	data, err := util.StringArrayToIntArray(input)
	util.CheckError(err)

	return data
}
