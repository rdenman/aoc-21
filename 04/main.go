package main

import (
	"fmt"
	"strings"

	"github.com/rdenman/aoc-21/util"
)

type board [][]int

func main() {
	sol1 := solution1()
	sol2 := solution2()

	fmt.Printf("Solution #1: %d\n", sol1)
	fmt.Printf("Solution #2: %d\n", sol2)
}

func solution1() int {
	nums, boards := readInputData()

	won := false
	var winner board
	var lastNum int
	for _, num := range nums {
		for _, b := range boards {
			won = checkBoard(b, num)
			if won {
				winner = b
				lastNum = num
				break
			}
		}
		if won {
			break
		}
	}

	sum := 0
	for _, row := range winner {
		for _, num := range row {
			if num != -1 {
				sum += num
			}
		}
	}

	return lastNum * sum
}

func solution2() int {
	nums, boards := readInputData()

	left := len(boards)
	var lastNum int
	var loser board
	for _, num := range nums {
		for i, b := range boards {
			if boards[i] != nil {
				won := checkBoard(b, num)
				if won {
					boards[i] = nil
					left--
				}

				if left == 0 {
					lastNum = num
					loser = b
					break
				}
			}
		}
		if left == 0 {
			break
		}
	}

	sum := 0
	for _, row := range loser {
		for _, num := range row {
			if num != -1 {
				sum += num
			}
		}
	}

	return lastNum * sum
}

func checkBoard(b board, n int) bool {
	done := false
	for i, row := range b {
		for j, col := range row {
			if col == n {
				b[i][j] = -1
				done = true
				break
			}
		}
		if done {
			break
		}
	}

	return isWinner(b)
}

// TODO could use a global var to optimize these calcs
func isWinner(b board) bool {
	rows := make([]int, 5)
	cols := make([]int, 5)
	for i, row := range b {
		for j, num := range row {
			if num == -1 {
				rows[i]++
				cols[j]++
			}
		}
	}

	won := false
	for i := 0; i < 5; i++ {
		if rows[i] == 5 || cols[i] == 5 {
			won = true
			break
		}
	}

	return won
}

func readInputData() ([]int, []board) {
	data, err := util.ReadAndSplitInputData("./input.txt")
	util.CheckError(err)

	nums, err := util.StringArrayToIntArray(strings.Split(data[0], ","))
	util.CheckError(err)

	var boards []board
	for i := 1; i < len(data); i++ {
		if data[i] != "" {
			b := make(board, 5)
			for j := 0; j < 5; j++ {
				b[j], err = util.StringArrayToIntArray(strings.Fields(data[i]))
				util.CheckError(err)
				i++
			}
			boards = append(boards, b)
		}
	}

	return nums, boards
}
