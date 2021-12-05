package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/rdenman/aoc-21/util"
)

type point struct {
	x int
	y int
}

type line struct {
	p1 point
	p2 point
}

var maxX int
var maxY int

func main() {
	sol1 := solution1()
	sol2 := solution2()

	fmt.Printf("Solution #1: %d\n", sol1)
	fmt.Printf("Solution #2: %d\n", sol2)
}

func solution1() int {
	maxX = 0
	maxY = 0
	lines := readInputData()
	board := makeBoard()

	for _, l := range lines {
		if l.p1.x == l.p2.x {
			start := math.Min(float64(l.p1.y), float64(l.p2.y))
			stop := math.Max(float64(l.p1.y), float64(l.p2.y))
			for i := int(start); i <= int(stop); i++ {
				board[i][l.p1.x]++
			}
		}

		if l.p1.y == l.p2.y {
			start := math.Min(float64(l.p1.x), float64(l.p2.x))
			stop := math.Max(float64(l.p1.x), float64(l.p2.x))
			for i := int(start); i <= int(stop); i++ {
				board[l.p1.y][i]++
			}
		}
	}

	c := 0
	for _, row := range board {
		for _, v := range row {
			if v > 1 {
				c++
			}
		}
	}

	return c
}

func solution2() int {
	maxX = 0
	maxY = 0
	lines := readInputData()
	board := makeBoard()

	for _, l := range lines {
		startX := int(math.Min(float64(l.p1.x), float64(l.p2.x)))
		stopX := int(math.Max(float64(l.p1.x), float64(l.p2.x)))
		startY := int(math.Min(float64(l.p1.y), float64(l.p2.y)))
		stopY := int(math.Max(float64(l.p1.y), float64(l.p2.y)))

		if startX == stopX {
			for i := int(startY); i <= int(stopY); i++ {
				board[i][l.p1.x]++
			}
		} else if startY == stopY {
			for i := int(startX); i <= int(stopX); i++ {
				board[l.p1.y][i]++
			}
		} else {
			run := (l.p2.x - l.p1.x) / int(math.Abs(float64(l.p2.x - l.p1.x)))
			rise := (l.p2.y - l.p1.y) / int(math.Abs(float64(l.p2.y - l.p1.y)))
			for x, y := l.p1.x, l.p1.y; x != l.p2.x + run && y != l.p2.y + rise; x, y = x + run, y + rise {
				board[y][x]++
			}	
		}
	}

	c := 0
	for _, row := range board {
		for _, v := range row {
			if v > 1 {
				c++
			}
		}
	}

	return c
}

func readInputData() []line {
	input, err := util.ReadAndSplitInputData("./input.txt")
	util.CheckError(err)

	data := make([]line, len(input))
	for i, str := range input {
		data[i] = parseLine(str)
		updateMax(data[i])
	}

	return data
}

func parseLine(str string) line {
	pts := strings.Split(str, " -> ")
	ps1 := strings.Split(pts[0], ",")
	ps2 := strings.Split(pts[1], ",")

	p1x, err := strconv.Atoi(ps1[0])
	util.CheckError(err)
	p1y, err := strconv.Atoi(ps1[1])
	util.CheckError(err)
	p2x, err := strconv.Atoi(ps2[0])
	util.CheckError(err)
	p2y, err := strconv.Atoi(ps2[1])
	util.CheckError(err)

	p1 := point{p1x, p1y}
	p2 := point{p2x, p2y}

	return line{p1, p2}
}

func updateMax(l line) {
	if l.p1.x > maxX {
		maxX = l.p1.x
	}
	if l.p2.x > maxX {
		maxX = l.p2.x
	}
	if l.p1.y > maxY {
		maxY = l.p1.y
	}
	if l.p2.y > maxY {
		maxY = l.p2.y
	}
}

func makeBoard() [][]int {
	board := make([][]int, maxY + 1)

	for i := range board {
		board[i] = make([]int, maxX + 1)
	}

	return board
}
