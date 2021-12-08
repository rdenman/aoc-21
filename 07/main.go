package main

import (
	"fmt"
	"sort"

	"github.com/rdenman/aoc-21/util"
)

func main() {
	sol1 := solution1()
	sol2 := solution2()

	fmt.Printf("Solution #1: %d\n", sol1)
	fmt.Printf("Solution #2: %d\n", sol2)
}

func solution1() int {
	data := readInputData()
	sort.Ints(data)

	med := data[len(data)/2]

	return calcCost(data, med)
}

func solution2() int {
	data := readInputData()

	sum := 0
	for _, d := range data {
		sum += d
	}

	mean := sum / len(data)
	cost := calcCostMult(data, mean)
	for i := 1; ; i++ {
		aCost := calcCostMult(data, mean+i)
		bCost := calcCostMult(data, mean-i)

		changed := false
		if aCost < cost {
			cost = aCost
			changed = true
		}
		if bCost < cost {
			cost = bCost
			changed = true
		}

		if !changed {
			break
		}
	}

	return cost
}

func readInputData() []int {
	input, err := util.ReadAndSplitInputData("./input.txt", ",")
	util.CheckError(err)

	data, err := util.StringArrayToIntArray(input)
	util.CheckError(err)

	return data
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calcCost(data []int, x int) int {
	cost := 0
	for _, d := range data {
		cost += abs(d - x)
	}
	return cost
}

func calcCostMult(data []int, x int) int {
	cost := 0
	for _, d := range data {
		change := abs(d - x)
		cost += (change * (change + 1)) / 2
	}
	return cost
}
