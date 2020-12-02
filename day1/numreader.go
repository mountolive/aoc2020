package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadNumbers(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %w", err)
	}
	defer f.Close()
	result := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("An error occurred converting %s",
				scanner.Text())
		}
		result = append(result, val)
	}
	return result, nil
}
