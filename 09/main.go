package main

import (
	"fmt"
	"sort"
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
				lows[p] = getBasinSize(data, p, map[point]bool{p: true})
			}
		}
	}

	set := []int{}
	for _, c := range lows {
		set = append(set, c)
	}

	sort.Ints(set)

	count := len(set)
	return set[count-1] * set[count-2] * set[count-3]
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
	up, down, left, right := getAdjacentIndicies(len(data), len(data[i]), i, j)

	return (up == -1 || value < data[up][j]) &&
		(down == -1 || value < data[down][j]) &&
		(left == -1 || value < data[i][left]) &&
		(right == -1 || value < data[i][right])
}

func getBasinSize(data [][]int, p point, visited map[point]bool) int {
	up, down, left, right := getAdjacentIndicies(len(data), len(data[p.I]), p.I, p.J)
	adj := []point{}

	if up != -1 && data[up][p.J] != 9 && !visited[point{up, p.J}] {
		adj = append(adj, point{up, p.J})
		visited[point{up, p.J}] = true
	}

	if down != -1 && data[down][p.J] != 9 && !visited[point{down, p.J}] {
		adj = append(adj, point{down, p.J})
		visited[point{down, p.J}] = true
	}

	if left != -1 && data[p.I][left] != 9 && !visited[point{p.I, left}] {
		adj = append(adj, point{p.I, left})
		visited[point{p.I, left}] = true
	}

	if right != -1 && data[p.I][right] != 9 && !visited[point{p.I, right}] {
		adj = append(adj, point{p.I, right})
		visited[point{p.I, right}] = true
	}

	for _, a := range adj {
		getBasinSize(data, a, visited)
	}

	return len(visited)
}

func getAdjacentIndicies(iLen, jLen, i, j int) (int, int, int, int) {
	up, down, left, right := -1, -1, -1, -1
	if i > 0 {
		up = i - 1
	}
	if i < iLen-1 {
		down = i + 1
	}
	if j > 0 {
		left = j - 1
	}
	if j < jLen-1 {
		right = j + 1
	}
	return up, down, left, right
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
