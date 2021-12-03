package main

import (
	"fmt"
	"strconv"
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

	tot := len(data)
	gam, eps := "", ""
	for i := range data[0] {
		ones := 0
		for j := range data {
			if data[j][i] == "1" {
				ones++
			}
		}

		if ones > (tot / 2) {
			gam += "1"
			eps += "0"
		} else {
			gam += "0"
			eps += "1"
		}
	}


	g, err := strconv.ParseInt(gam, 2, 64)
	util.CheckError(err)
	e, err := strconv.ParseInt(eps, 2, 64)
	util.CheckError(err)

	return int(g) * int(e)
}

func solution2() int {
	data := readInputData()

	oxy := calculateRating(data, "1", "0")
	co2 := calculateRating(data, "0", "1")

	o, err := strconv.ParseInt(oxy, 2, 64)
	util.CheckError(err)
	c, err := strconv.ParseInt(co2, 2, 64)
	util.CheckError(err)

	return int(o) * int(c)
}

func readInputData() [][]string {
	input, err := util.ReadAndSplitInputData("./input.txt")
	util.CheckError(err)

	data := make([][]string, len(input))
	for i, b := range input {
		data[i] = strings.Split(b, "")
	}

	return data
}

func calculateRating(data [][]string, h string, l string) string {
	arr := make([][]string, len(data))
	copy(arr, data)
	for i := range arr[0] {
		ones := 0
		for _, row := range arr {
			if row[i] == "1" {
				ones++
			}
		}

		if ones >= (len(arr) - ones) {
			arr = onlyKeep(arr, h, i)
		} else {
			arr = onlyKeep(arr, l, i)
		}

		if len(arr) == 1 {
			break
		}
	}

	return strings.Join(arr[0], "")
}

func onlyKeep(s [][]string, v string, i int) [][]string {
	var new [][]string
	for _, row := range s {
		if row[i] == v {
			new = append(new, row)
		}
	}
	return new
}
