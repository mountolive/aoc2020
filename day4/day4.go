package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const day4Problem = "./day4/day4_problem"

type PassportCounter struct {
	passports    []string
	hairColorReg *regexp.Regexp
	eyeColorArr  [7]string
}

func NewPassportCounter() (*PassportCounter, error) {
	f, err := os.Open(day4Problem)
	if err != nil {
		return nil, fmt.Errorf("Error opening the file: %w", err)
	}
	defer f.Close()

	lines := []string{}
	scanner := bufio.NewScanner(f)
	passport := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			lines = append(lines, strings.Join(passport, " "))
			passport = []string{}
			continue
		}
		passport = append(passport, text)
	}
	if len(passport) != 0 {
		lines = append(lines, strings.Join(passport, " "))
	}
	hairColorReg, err := regexp.Compile("^[a-f0-9]*$")
	if err != nil {
		return nil, fmt.Errorf("Error compiling hair color regexp")
	}
	return &PassportCounter{
		passports:    lines,
		hairColorReg: hairColorReg,
		eyeColorArr: [7]string{
			"amb", "blu", "brn",
			"gry", "grn", "hzl",
			"oth"},
	}, nil
}

func (p *PassportCounter) CountValidPassportsPart1() int {
	validPassports := 0
	for _, passport := range p.passports {
		data := strings.Split(passport, " ")
		if len(data) < 7 {
			continue
		}
		if len(data) == 8 {
			validPassports += 1
		}
		if isCidIncluded(data) {
			continue
		}
		validPassports++
	}
	return validPassports
}

func (p *PassportCounter) CountValidPassportsPart2() int {
	validPassports := 0
	for _, passport := range p.passports {
		data := strings.Split(passport, " ")
		if len(data) < 7 {
			continue
		}
		if len(data) == 7 && isCidIncluded(data) {
			continue
		}

		if p.validPassport(data) {
			validPassports++
		}
	}
	return validPassports
}

func (p *PassportCounter) validPassport(data []string) bool {
	for _, param := range data {
		pair := strings.Split(param, ":")
		if pair[0] == "cid" {
			continue
		}
		if !p.validField(pair) {
			return false
		}
	}
	return true
}

func (p *PassportCounter) validField(pair []string) bool {
	switch pair[0] {
	case "pid":
		return validPid(pair[1])
	case "ecl":
		return p.validEclLinear(pair[1])
	case "byr":
		return correctByr(pair[1])
	case "iyr":
		return correctIyr(pair[1])
	case "eyr":
		return correctEyr(pair[1])
	case "hgt":
		return correctHgt(pair[1])
	case "hcl":
		if pair[1][0] != '#' {
			return false
		}
		return p.hairColorReg.MatchString(pair[1][1:])
	}
	return true
}

func isCidIncluded(data []string) bool {
	for _, param := range data {
		pair := strings.Split(param, ":")
		if pair[0] == "cid" {
			return true
		}
	}
	return false
}

func validPid(pid string) bool {
	n := len(pid)
	if n != 9 {
		return false
	}
	_, err := strconv.Atoi(pid)
	if err != nil {
		return false
	}
	return true
}

func (p *PassportCounter) validEclLinear(color string) bool {
	for _, c := range p.eyeColorArr {
		if c == color {
			return true
		}
	}
	return false
}

func correctByr(value string) bool {
	return correctRange(value, 1920, 2002)
}

func correctIyr(value string) bool {
	return correctRange(value, 2010, 2020)
}

func correctEyr(value string) bool {
	return correctRange(value, 2020, 2030)
}

func correctHgt(value string) bool {
	n := len(value)
	ruler := value[n-2 : n]
	if ruler == "in" {
		return correctRange(value[:n-2], 59, 76)
	}
	if ruler == "cm" {
		return correctRange(value[:n-2], 150, 193)
	}
	return false
}

func correctRange(val string, min, max int) bool {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return intVal >= min && intVal <= max
}
