package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rdenman/aoc-21/util"
)

type command struct {
	dir string
	val int
}

func main() {
	sol1 := solution1()
	sol2 := solution2()

	fmt.Printf("Solution #1: %d\n", sol1)
	fmt.Printf("Solution #2: %d\n", sol2)
}

func solution1() int {
	data := readInputData()

	hor, dep := 0, 0
	for _, c := range data {
		switch c.dir {
		case "forward":
			hor += c.val
		case "down":
			dep += c.val
		case "up":
			dep -= c.val
		}
	}

	return hor * dep
}

func solution2() int {
	data := readInputData()

	hor, dep, aim := 0, 0, 0
	for _, c := range data {
		switch c.dir {
		case "forward":
			hor += c.val
			dep += aim * c.val
		case "down":
			aim += c.val
		case "up":
			aim -= c.val
		}
	}

	return hor * dep
}

func readInputData() []command {
	rows, err := util.ReadAndSplitInputData("./input.txt")
	util.CheckError(err)

	data := make([]command, len(rows))
	for i, r := range rows {
		kv := strings.Split(r, " ")
		num, err := strconv.Atoi(kv[1])
		util.CheckError(err)

		data[i] = command{kv[0], num}
	}

	return data
}
