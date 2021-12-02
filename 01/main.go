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

func solution1() int {
	data := readInputData()
	ans, prev := 0, 999999
	for _, d := range data {
		if d > prev {
			ans++
		}
		prev = d
	}
	return ans
}

func solution2() int {
	data := readInputData()
	ans, prev := 0, 999999
	for i := range data {
		if i < len(data)-2 {
			sum := data[i] + data[i+1] + data[i+2]
			if sum > prev {
				ans++
			}
			prev = sum
		}
	}
	return ans
}

func readInputData() []int {
	data, err := util.ReadAndSplitInputData("./input.txt")
	util.CheckError(err)

	ints, err := util.StringArrayToIntArray(data)
	util.CheckError(err)
	return ints
}
