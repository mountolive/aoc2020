package day5

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
)

const day5Problem = "./day5/day5_problem"

var (
	leftRightReg = regexp.MustCompile(`[LR]`)
	frontBackReg = regexp.MustCompile(`[FB]`)

	InvalidRowDataError    = errors.New("Invalid boarding pass, Row")
	InvalidColumnDataError = errors.New("Invalid boarding pass, Column")
)

func FindHighestSeat() (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	lines := readLines(ctx, day5Problem)
	var max int
	lineNum := 0
	for line := range lines {
		lineNum += 1
		val, err := processTicket(line)
		if err != nil {
			return 0, fmt.Errorf("%w. Line %d", err, lineNum)
		}
		if val >= max {
			max = val
		}
	}
	return max, nil
}

func processTicket(line string) (int, error) {
	row, err := calculateRow(line[:7])
	if err != nil {
		return 0, err
	}
	col, err := calculateColumn(line[7:])
	if err != nil {
		return 0, err
	}
	return row*8 + col, nil
}

func calculateColumn(segment string) (int, error) {
	err := validateRunes(segment, leftRightReg)
	if err != nil {
		return 0, wrapError(InvalidColumnDataError, err)
	}
	column := 0
	start := 0
	bound := 7
	for _, step := range segment {
		half := halfFunc(bound, start)
		switch step {
		case 'L':
			bound = half
			column = start
			continue
		case 'R':
			start = half
			column = bound
			continue
		default:
			return 0, InvalidColumnDataError
		}
	}

	return column, nil
}

func calculateRow(segment string) (int, error) {
	err := validateRunes(segment, frontBackReg)
	if err != nil {
		return 0, wrapError(InvalidRowDataError, err)
	}
	row := 0
	start := 0
	bound := 127
	for _, step := range segment {
		half := halfFunc(bound, start)
		switch step {
		case 'F':
			bound = half
			row = start
			continue
		case 'B':
			start = half
			row = bound
			continue
		default:
			return 0, InvalidRowDataError
		}
	}
	return row, nil
}

func halfFunc(bound, start int) int {
	return (bound - start) / 2
}

func validateRunes(segment string, reg *regexp.Regexp) error {
	for _, letter := range segment {
		if !reg.Match([]byte(string(letter))) {
			return fmt.Errorf("Not matched: %s", string(letter))
		}
	}
	return nil
}

func wrapError(wrapper, err error) error {
	return fmt.Errorf("%w: %s", wrapper, err.Error())
}

func readLines(ctx context.Context, path string) <-chan string {
	lineStream := make(chan string)
	go func() {
		defer close(lineStream)
		f, err := os.Open(path)
		if err != nil {
			fmt.Printf("Error reading file %v\n", err)
			return
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			select {
			case <-ctx.Done():
				fmt.Println("Context canceled")
				return
			default:
				lineStream <- scanner.Text()
			}
		}
	}()
	return lineStream
}
