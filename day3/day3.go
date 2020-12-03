package day3

import (
	"bufio"
	"fmt"
	"os"
)

const day3Problem = "./day3/day3_problem"

type TreeCounter struct {
	treeLayout []string
}

func NewTreeCounter() (*TreeCounter, error) {
	f, err := os.Open(day3Problem)
	if err != nil {
		return nil, fmt.Errorf("Error opening the file: %w", err)
	}
	defer f.Close()

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return &TreeCounter{lines}, nil
}

func (t *TreeCounter) CountTrees(right, down int) int {
	trees := 0
	rightStep := right
	rowLen := 0

	for i := 0; i < len(t.treeLayout); i += down {
		row := t.treeLayout[i]
		if rowLen == 0 {
			rowLen = len(row)
			continue
		}
		if row[rightStep] == '#' {
			trees += 1
		}
		rightStep += right
		if rightStep >= rowLen {
			rightStep -= rowLen
		}
	}

	return trees
}

func (t *TreeCounter) MultiplyCount() int {
	rules := []struct{ right, down int }{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	multiplied := 1
	for _, pair := range rules {
		multiplied = multiplied * t.CountTrees(pair.right, pair.down)
	}
	return multiplied
}
