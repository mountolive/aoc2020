package day1

import (
	"fmt"
	"strconv"
	"strings"
)

const currPath = "./day1/day1_problem"

func SolutionPartOne() (int, error) {
	nums, err := ReadNumbers(currPath)
	if err != nil {
		return 0, fmt.Errorf("Solution's (Part one) read error: %w", err)
	}
	return solutionPartOne(nums), nil
}

func SolutionPartTwo() (int, error) {
	nums, err := ReadNumbers(currPath)
	if err != nil {
		return 0, fmt.Errorf("Solution's (Part two) read error: %w", err)
	}
	return solutionPartTwo(nums)
}

func solutionPartOne(nums []int) int {
	year := 2020
	dict := make(map[int]byte)
	for _, num := range nums {
		diff := year - num
		if dict[diff] == 1 {
			return num * diff
		}
		dict[num] = 1
	}
	for _, num := range nums {
		diff := year - num
		if dict[diff] == 1 {
			return num * diff
		}
	}
	return 0
}

func solutionPartTwo(nums []int) (int, error) {
	year := 2020
	dict := make(map[int]string)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			dict[nums[i]+nums[j]] = fmt.Sprintf("%d_%d", i, j)
		}
	}
	for k, num := range nums {
		diff := year - num
		indexes := dict[diff]
		if indexes != "" {
			strIndexes := strings.Split(indexes, "_")
			first, err := strconv.Atoi(strIndexes[0])
			if err != nil {
				return 0, fmt.Errorf("Unable to parse %s", strIndexes[0])
			}
			second, err := strconv.Atoi(strIndexes[1])
			if err != nil {
				return 0, fmt.Errorf("Unable to parse %s", strIndexes[0])
			}
			if k != first && k != second {
				return nums[first] * nums[second] * num, nil
			}
		}
	}
	return 0, fmt.Errorf("Not found")
}
