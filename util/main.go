package util

import "strconv"

func StringArrayToIntArray(data []string) ([]int, error) {
	var ints []int

	for _, s := range data {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		ints = append(ints, i)
	}

	return ints, nil
}
