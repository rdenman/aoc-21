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
	template, rules := readInputData()

	for i := 0; i < 10; i++ {
		built := ""
		for j, ru := range template {
			ch := string(ru)
			if j == len(template) - 1 {
				built += ch
			} else {
				pair := ch + string(template[j + 1])
				built += ch + rules[pair]
			}
		}
		template = built
	}

	counts := map[string]int{}
	for _, ru := range template {
		ch := string(ru)
		counts[ch]++
	}

	min, max := 999999, -1
	for _, val := range counts {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return max - min
}

func solution2() int {
	const steps = 40
	template, rules := readInputData()
	initialPairs := getTemplatePairs(template)

	pairs := []string{}
	for k := range rules {
		pairs = append(pairs, k)
	}

	pairCount := map[string][]int{}
	for _, p := range pairs {
		pairCount[p] = make([]int, steps + 1)
	}

	counts := make(map[string]int)
	for _, ru := range template {
		ch := string(ru)
		counts[ch]++
	}
	for _, pair := range initialPairs {
		pairCount[pair][0]++
	}

	for step := 1; step <= steps; step++ {
		lastStep := step - 1
		for _, pair := range pairs {
			lastCount := pairCount[pair][lastStep]
			if lastCount > 0 {
				ps := strings.Split(pair, "")
				rule := rules[pair]
				pairCount[ps[0] + rule][step] += lastCount
				pairCount[rule + ps[1]][step] += lastCount
				counts[rule] += lastCount
			}
		}
	}

	min, max := 999999999999999999, -1
	for _, val := range counts {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return max - min
}

func readInputData() (string, map[string]string) {
	input, err := util.ReadAndSplitInputData("./input.txt")
	util.CheckError(err)

	rules := make(map[string]string)
	for _, r := range input[2:] {
		kv := strings.Split(r, " -> ")
		rules[kv[0]] = kv[1]
	}

	return input[0], rules
}

func getTemplatePairs(template string) []string {
	pairs := make([]string, len(template) - 1)
	for i := range pairs {
		pairs[i] = string(template[i]) + string(template[i + 1])
	}
	return pairs
}