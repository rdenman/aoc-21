package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/rdenman/aoc-21/util"
)

type entry struct {
	input []string
	output  []string
}

func main() {
	sol1 := solution1()
	sol2 := solution2()

	fmt.Printf("Solution #1: %d\n", sol1)
	fmt.Printf("Solution #2: %d\n", sol2)
}

func solution1() int {
	data := readInputData()

	count := 0
	for _, ent := range data {
		for _, out := range ent.output {
			l := len(out)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count++
			}
		}
	}

	return count
}

func solution2() int {
	data := readInputData()

	count := 0
	for _, ent := range data {
		count += sumOutput(ent.input, ent.output)
	}

	return count
}

func readInputData() []entry {
	input, err := util.ReadAndSplitInputData("./input.txt")
	util.CheckError(err)

	data := make([]entry, len(input))
	for i, line := range input {
		sp := strings.Split(line, "|")
		data[i] = entry{
			strings.Fields(sp[0]),
			strings.Fields(sp[1]),
		}
	}

	return data
}

func sumOutput(input, output []string) int {
	lenMap := make(map[int][]string)
	for _, in := range input {
		lenMap[len(in)] = append(lenMap[len(in)], in)
	}

	solution := map[string]int{
		lenMap[2][0]: 1,
		lenMap[3][0]: 7,
		lenMap[4][0]: 4,
		lenMap[7][0]: 8,
	}

	one := lenMap[2][0]
	four := lenMap[4][0]	

	var nine string
	for _, sixLettered := range lenMap[6] {
		if len(removeFromString(sixLettered, one)) == 5 {
			solution[sixLettered] = 6
		} else if len(removeFromString(sixLettered, four)) == 3 {
			solution[sixLettered] = 0
		} else {
			nine = sixLettered
			solution[sixLettered] = 9
		}
	}

	for _, fiveLettered := range lenMap[5] {
		if len(removeFromString(fiveLettered, one)) == 3 { 
			solution[fiveLettered] = 3
		} else if len(removeFromString(fiveLettered, nine)) == 0 {
			solution[fiveLettered] = 5
		} else {
			solution[fiveLettered] = 2
		}
	}

	count := 0
	for i, out := range output {
		if _, ok := solution[out]; ok {
			count += solution[out] * int(math.Pow(10, float64(3 - i)))
			continue
		}

		L:
		for _, poss := range lenMap[len(out)] {
			for _, char := range poss {
				if !strings.ContainsRune(out, char) {
					continue L
				}
			}

			count += solution[poss] * int(math.Pow(10, float64(3 - i)))
			break
		}
	}

	return count
}

func removeFromString(str, rem string) string {
	for _, char := range rem {
		str = strings.Replace(str, string(char), "", 1)
	}
	return str
}
