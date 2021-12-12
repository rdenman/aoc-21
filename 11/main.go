package main

import (
	"fmt"
	"strings"

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

	count := 0
	for i := 0; i < 100; i++ {
		count += step(data)
	}

	return count
}

func solution2() int {
	data := readInputData()

	flashes := []int{}
	for {
		newFlashes := step(data)
		flashes = append(flashes, newFlashes)
		if newFlashes == len(data)*len(data[0]) {
			break
		}
	}

	return len(flashes)
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

func step(data [][]int) int {
	for i := range data {
		for j := range data[i] {
			data[i][j]++
		}
	}

	flashes := 0
	flashers := getFlashers(data)
	for len(flashers) > 0 {
		flashes += len(flashers)
		flash(data, flashers)
		flashers = getFlashers(data)
	}

	return flashes
}

func getFlashers(data [][]int) [][]int {
	flashers := [][]int{}
	for i := range data {
		for j, val := range data[i] {
			if val > 9 {
				flashers = append(flashers, []int{i, j, val})
			}
		}
	}
	return flashers
}

func flash(data [][]int, flashers [][]int) {
	for _, fl := range flashers {
		coords := getCoords(len(data), fl[0], fl[1])
		for _, c := range coords {
			if data[c[0]][c[1]] != 0 {
				data[c[0]][c[1]]++
			}
		}

		data[fl[0]][fl[1]] = 0
	}
}

func getCoords(size, i, j int) [][]int {
	coords := [][]int{}
	if i-1 >= 0 {
		coords = append(coords, []int{i - 1, j})
		if j-1 >= 0 {
			coords = append(coords, []int{i - 1, j - 1})
		}
		if j+1 < size {
			coords = append(coords, []int{i - 1, j + 1})
		}
	}

	if i+1 < size {
		coords = append(coords, []int{i + 1, j})
		if j-1 >= 0 {
			coords = append(coords, []int{i + 1, j - 1})
		}
		if j+1 < size {
			coords = append(coords, []int{i + 1, j + 1})
		}
	}

	if j-1 >= 0 {
		coords = append(coords, []int{i, j - 1})
	}

	if j+1 < size {
		coords = append(coords, []int{i, j + 1})
	}

	return coords
}
