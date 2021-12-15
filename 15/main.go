package main

import (
	"fmt"
	"strings"

	"github.com/rdenman/aoc-21/util"
)

type coord struct {
	x, y int
}

func main() {
	sol1 := solution1()
	sol2 := solution2()

	fmt.Printf("Solution #1: %d\n", sol1)
	fmt.Printf("Solution #2: %d\n", sol2)
}

func solution1() int {
	grid, size := readInputData()
	return findMinPathLength(grid, size)
}

func solution2() int {
	grid, size := readInputData()
	grid = increaseGridSize(grid, size, 5)
	return findMinPathLength(grid, size*5)
}

func readInputData() (map[coord]int, int) {
	data, err := util.ReadAndSplitInputData("./input.txt")
	util.CheckError(err)

	grid := make(map[coord]int)
	for y, row := range data {
		line, err := util.StringArrayToIntArray(strings.Split(row, ""))
		util.CheckError(err)
		for x, val := range line {
			grid[coord{x, y}] = val
		}
	}

	return grid, len(data)
}

func findMinPathLength(grid map[coord]int, size int) int {
	min := 0
	for x := 1; x < size; x++ {
		if v, ok := grid[coord{x, 0}]; ok {
			min += v
		}
	}
	for y := 1; y < size; y++ {
		if v, ok := grid[coord{size - 1, y}]; ok {
			min += v
		}
	}

	mins := make(map[coord]int)
	for k := range grid {
		mins[k] = -1
	}
	mins[coord{0, 0}] = 0

	end := coord{size - 1, size - 1}
	fringe := []coord{{0, 0}}
	n := 0
	for len(fringe) > 0 {
		nextFringe := []coord{}
		for _, tip := range fringe {
			t := mins[tip]
			for _, next := range []coord{{tip.x + 1, tip.y}, {tip.x - 1, tip.y}, {tip.x, tip.y + 1}, {tip.x, tip.y - 1}} {
				if val, ok := grid[next]; ok {
					nextVal := val + t
					prevVal := mins[next]

					if nextVal < prevVal || prevVal == -1 {
						mins[next] = nextVal
						if next == end {
							min = nextVal
						}
						nextFringe = append(nextFringe, next)
					}
				}
			}
		}
		n += 1
		fringe = nextFringe
	}

	return min
}

func increaseGridSize(original map[coord]int, originalSize, inc int) map[coord]int {
	next := make(map[coord]int)
	for i := 0; i < inc; i++ {
		for j := 0; j < inc; j++ {
			for c, val := range original {
				x, y := c.x, c.y
				n := (val + i + j)
				if n > 9 {
					n -= 9
				}
				next[coord{x + i*originalSize, y + j*originalSize}] = n
			}
		}
	}
	return next
}
