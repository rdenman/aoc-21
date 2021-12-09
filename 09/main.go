package main

import (
	"fmt"
	"strings"

	"github.com/rdenman/aoc-21/util"
)

type point struct {
	I, J int
}

func main() {
	sol1 := solution1()
	sol2 := solution2()

	fmt.Printf("Solution #1: %d\n", sol1)
	fmt.Printf("Solution #2: %d\n", sol2)
}

func solution1() int {
	data := readInputData()

	sum := 0
	for i, row := range data {
		for j, v := range row {
			if isLowPoint(data, v, i, j) {
				sum += (v + 1)
			}
		}
	}

	return sum
}

func solution2() int {
	data := readInputData()

	lows := make(map[point]int)
	for i, row := range data {
		for j, v := range row {
			if isLowPoint(data, v, i, j) {
				p := point{i, j}
				lows[p] = getBasinSize(data, p, []point{p})
			}
		}
	}

	fmt.Println(lows)

	return 1
}

func readInputData() [][]int {
	input, err := util.ReadAndSplitInputData("./input.txt")
	util.CheckError(err)

	data := make([][]int, len(input))
	for i, row := range input {
		data[i], err = util.StringArrayToIntArray(strings.Split(row, ""))
		util.CheckError(err)
	}

	return data
}

func isLowPoint(data [][]int, value, i, j int) bool {
	up, down, left, right := -1, -1, -1, -1
	if i > 0 {
		up = i - 1
	}
	if i < len(data) - 1 {
		down = i + 1
	}
	if j > 0 {
		left = j - 1
	}
	if j < len(data[i]) - 1 {
		right = j + 1
	}

	return (up == -1 || value < data[up][j]) &&
		(down == -1 || value < data[down][j]) &&
		(left == -1 || value < data[i][left]) && 
		(right == -1 || value < data[i][right])
}

func getBasinSize(data [][]int, p point, visited []point) int {
	size := 0

	if p.I > 0 && data[p.I - 1][p.J] < 9 && !hasVisited(visited, point{p.I - 1, p.J}) {
		visited = append(visited, point{p.I - 1, p.J})
		size += 1 + getBasinSize(data, point{p.I - 1, p.J}, visited)
	}

	if p.I < len(data) - 1 && data[p.I + 1][p.J] < 9 && !hasVisited(visited, point{p.I + 1, p.J}) {
		visited = append(visited, point{p.I + 1, p.J})
		size += 1 + getBasinSize(data, point{p.I + 1, p.J}, visited)
	}

	if p.J > 0 && data[p.I][p.J - 1] < 9 && !hasVisited(visited, point{p.I, p.J - 1}) {
		visited = append(visited, point{p.I, p.J - 1})
		size += 1 + getBasinSize(data, point{p.I, p.J - 1}, visited)
	}

	if p.J < len(data[p.I]) - 1 && data[p.I][p.J + 1] < 9 && !hasVisited(visited, point{p.I, p.J + 1}) {
		visited = append(visited, point{p.I, p.J + 1})
		size += 1 + getBasinSize(data, point{p.I, p.J + 1}, visited)
	}

	return size
}

func hasVisited(li []point, p point) bool {
	has := false
	for _, v := range li {
		if v.I == p.I && v.J == p.J {
			has = true
			break
		}
	}
	return has
}
