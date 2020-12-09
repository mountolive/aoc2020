package main

import (
	"fmt"
	"log"

	"github.com/mountolive/advent2020/day1"
	"github.com/mountolive/advent2020/day2"
	"github.com/mountolive/advent2020/day3"
	"github.com/mountolive/advent2020/day4"
	"github.com/mountolive/advent2020/day5"
)

func main() {
	fmt.Println("-----------DAY 1-------------")
	fmt.Println(day1.SolutionPartOne())
	fmt.Println(day1.SolutionPartTwo())
	fmt.Println("-----------DAY 2-------------")
	fmt.Println(day2.ValidPasswordsPart1())
	fmt.Println(day2.ValidPasswordsPart2())
	fmt.Println("-----------DAY 3-------------")
	treeCounter, err := day3.NewTreeCounter()
	if err != nil {
		log.Fatalf("Error creating the tree counter: %w", err)
	}
	fmt.Println(treeCounter.CountTrees(3, 1))
	fmt.Println(treeCounter.MultiplyCount())
	fmt.Println("------------DAY 4-------------")
	passportCounter, err := day4.NewPassportCounter()
	fmt.Println(passportCounter.CountValidPassportsPart1())
	fmt.Println(passportCounter.CountValidPassportsPart2())
	fmt.Println("-----------DAY 5--------------")
	fmt.Println(day5.FindHighestSeat())
}
