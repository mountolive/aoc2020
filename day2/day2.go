package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const day2Problem = "./day2/day2_problem"

type rule struct {
	firstNum,
	secondNum int
	letter rune
}

func ValidPasswordsPart2() (int, error) {
	return ruleProcessor(validPasswordPart2)
}

func ValidPasswordsPart1() (int, error) {
	return ruleProcessor(validPasswordPart1)
}

func ruleProcessor(process func(*int, string) error) (int, error) {
	f, err := os.Open(day2Problem)
	if err != nil {
		return 0, fmt.Errorf("Error reading file: %w", err)
	}
	defer f.Close()

	var counter int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		err := process(&counter, scanner.Text())
		if err != nil {
			return counter, err
		}
	}
	return counter, nil
}

func validPasswordPart1(counter *int, line string) error {
	ruler, password, err := getLineRule(line)
	if err != nil {
		return err
	}

	letter := ruler.letter
	min := ruler.firstNum
	max := ruler.secondNum

	passMap := make(map[rune]int)

	for _, l := range password {
		passMap[l] += 1
	}

	if passMap[letter] >= min && passMap[letter] <= max {
		*counter += 1
	}
	return nil
}

func validPasswordPart2(counter *int, line string) error {
	ruler, password, err := getLineRule(line)
	if err != nil {
		return err
	}

	letter := ruler.letter
	first := ruler.firstNum
	second := ruler.secondNum

	matches := 0
	for i, l := range password {
		idx := i + 1
		if (idx == first || idx == second) && l == letter {
			matches += 1
		}
	}

	if matches == 1 {
		*counter += 1
	}

	return nil
}

func getLineRule(line string) (*rule, string, error) {
	params := strings.Split(line, " ")
	splittedNums := strings.Split(params[0], "-")

	first, err := strconv.Atoi(splittedNums[0])
	if err != nil {
		return nil, "", fmt.Errorf("An error occurred converting %s",
			splittedNums[0])
	}
	second, err := strconv.Atoi(splittedNums[1])
	if err != nil {
		return nil, "", fmt.Errorf("An error occurred converting %s",
			splittedNums[1])
	}

	letter := rune(params[1][0])

	return &rule{first, second, letter}, params[2], nil
}
