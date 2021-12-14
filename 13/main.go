package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rdenman/aoc-21/util"
)

type coord struct {
	x, y int
}

type instruction struct {
	axis string
	val  int
}

func main() {
	sol1 := solution1()
	sol2 := solution2()

	fmt.Printf("Solution #1: %d\n", sol1)
	fmt.Printf("Solution #2: %d\n", sol2)
}

func solution1() int {
	coords, instructions, maxX, maxY := readInputData()
	sheet := createSheet(maxX, maxY, instructions)

	count := 0
	for _, c := range coords {
		x, y := fold(instructions[0], c.x, c.y)
		if sheet[x][y] == 0 {
			sheet[x][y] = 1
			count++
		}
	}

	return count
}

func solution2() int {
	coords, instructions, maxX, maxY := readInputData()
	sheet := createSheet(maxX, maxY, instructions)

	for _, c := range coords {
		x, y := c.x, c.y
		for _, ins := range instructions {
			x, y = fold(ins, x, y)
		}
		if sheet[x][y] == 0 {
			sheet[x][y] = 1
		}
	}

	printSheet(sheet)

	return 1
}

func readInputData() ([]coord, []instruction, int, int) {
	input, err := util.ReadAndSplitInputData("./input.txt", "\n\n")
	util.CheckError(err)

	maxX, maxY := 0, 0
	coords := make([]coord, len(strings.Split(input[0], "\n")))
	for i, line := range strings.Split(input[0], "\n") {
		c, err := util.StringArrayToIntArray(strings.Split(line, ","))
		util.CheckError(err)
		coords[i] = coord{c[0], c[1]}
		maxX = max(maxX, coords[i].x)
		maxY = max(maxY, coords[i].y)
	}

	instructions := make([]instruction, len(strings.Split(input[1], "\n")))
	for i, line := range strings.Split(input[1], "\n") {
		ins := strings.Split(strings.Fields(line)[2], "=")
		val, err := strconv.Atoi(ins[1])
		util.CheckError(err)
		instructions[i] = instruction{ins[0], val}
	}

	return coords, instructions, maxX + 1, maxY + 1
}

func createSheet(sizeX, sizeY int, instructions []instruction) [][]int {
	sheet := make([][]int, sizeX)
	for i := range sheet {
		sheet[i] = make([]int, sizeY)
	}
	return sheet
}

func printSheet(sheet [][]int) {
	for j := 0; j < len(sheet[0]); j++ {
		for i := 0; i < len(sheet); i++ {
			if sheet[i][j] == 1 {
				fmt.Printf(" # ")
			} else {
				fmt.Printf(" . ")
			}
		}
		fmt.Println("")
	}
}

func fold(ins instruction, x, y int) (int, int) {
	fx, fy := 0, 0
	if ins.axis == "x" {
		fx = ins.val
	} else {
		fy = ins.val
	}

	return abs(fx - abs(fx-x)), abs(fy - abs(fy-y))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
