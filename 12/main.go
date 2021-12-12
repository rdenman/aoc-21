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

	return search(data, "start", map[string]bool{}, false)
}

func solution2() int {
	data := readInputData()

	return search(data, "start", map[string]bool{}, true)
}

func search(adj map[string][]string, cave string, visited map[string]bool, canVisitSmall bool) int {
	if cave == "end" {
		return 1
	}

	if visited[cave] {
		if cave == "start" {
			return 0
		}

		if cave == strings.ToLower(cave) {
			if !canVisitSmall {
				return 0
			} else {
				canVisitSmall = false
			}
		}
	}

	sum := 0
	for _, a := range adj[cave] {
		copy := copyMap(visited)
		copy[cave] = true
		sum += search(adj, a, copy, canVisitSmall)
	}
	return sum
}

func readInputData() map[string][]string {
	input, err := util.ReadAndSplitInputData("./input.txt")
	util.CheckError(err)

	data := make(map[string][]string)
	for _, line := range input {
		paths := strings.Split(line, "-")
		data[paths[0]] = append(data[paths[0]], paths[1])
		data[paths[1]] = append(data[paths[1]], paths[0])
	}

	return data
}

func copyMap(m map[string]bool) map[string]bool {
	copy := make(map[string]bool)
	for k, v := range m {
		copy[k] = v
	}
	return copy
}
