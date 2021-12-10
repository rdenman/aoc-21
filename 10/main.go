package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rdenman/aoc-21/util"
)

var opened = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}
var closed = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var points1 = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}
var points2 = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func main() {
	sol1 := solution1()
	sol2 := solution2()

	fmt.Printf("Solution #1: %d\n", sol1)
	fmt.Printf("Solution #2: %d\n", sol2)
}

func solution1() int {
	data := readInputData()

	score := 0
	for _, line := range data {
		isCorrupt, corrupt := checkCorrupt(line)
		if isCorrupt {
			score += points1[corrupt]
		}
	}

	return score
}

func solution2() int {
	data := readInputData()

	incomplete := [][]string{}
	for _, line := range data {
		isCorrupt, _ := checkCorrupt(line)
		if !isCorrupt {
			incomplete = append(incomplete, line)
		}
	}

	scores := make([]int, len(incomplete))
	for i, line := range incomplete {
		scores[i] = getIncompleteScore(line)
	}

	sort.Ints(scores)

	return scores[(len(scores)-1)/2]
}

func readInputData() [][]string {
	input, err := util.ReadAndSplitInputData("./input.txt")
	util.CheckError(err)

	data := make([][]string, len(input))
	for i, row := range input {
		data[i] = strings.Split(row, "")
	}

	return data
}

func checkCorrupt(line []string) (bool, string) {
	isCorrupt, corrupt := false, ""

	open := []string{}
	for _, c := range line {
		if _, has := closed[c]; has {
			open = append(open, c)
		} else if open[len(open)-1] == opened[c] {
			open = open[:len(open)-1]
		} else {
			isCorrupt = true
			corrupt = c
			break
		}
	}

	return isCorrupt, corrupt
}

func getIncompleteScore(line []string) int {
	open := []string{}
	for _, c := range line {
		if _, has := closed[c]; has {
			open = append(open, c)
		} else if open[len(open)-1] == opened[c] {
			open = open[:len(open)-1]
		}
	}

	close := make([]string, len(open))
	last := len(open) - 1
	for i, c := range open {
		close[last-i] = closed[c]
	}

	score := 0
	for _, c := range close {
		score *= 5
		score += points2[c]
	}

	return score
}
