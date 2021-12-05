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

	fmt.Println(data)

	return 1
}

func solution2() int {
	data := readInputData()

	fmt.Println(data)

	return 1
}

func readInputData() []string {
	data, err := util.ReadAndSplitInputData("./input.txt")
	util.CheckError(err)

	return data
}
