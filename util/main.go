package util

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadAndSplitInputData(filename string, split_on ...string) ([]string, error) {
	input, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	sep := "\n"
	if len(split_on) > 0 {
		sep = split_on[0]
	}

	data := strings.Split(string(input), sep)
	return data, nil
}

func StringArrayToIntArray(data []string) ([]int, error) {
	ints := make([]int, len(data))
	for i, s := range data {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		ints[i] = num
	}

	return ints, nil
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
