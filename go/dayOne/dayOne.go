package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	elves, err := readInput("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var calsPerElf, mostCals = calcCalories(elves)
	var elfNumber = findElfNumber(calsPerElf, mostCals)

	fmt.Printf("most calories: %d\n", mostCals)
	fmt.Printf("elf number: %d\n", elfNumber)
}

func readInput(filename string) ([][]int, error) {
	input, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)
	var elves [][]int // outerlist
	var elf []int     // innerlist

	// scan and create nested list
	for scanner.Scan() {
		var foodCal int
		fmt.Sscanf(scanner.Text(), "%d", &foodCal)
		if foodCal == 0 {
			elves = append(elves, elf)
			elf = nil
		}
		elf = append(elf, foodCal)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return elves, nil
}

func calcCalories(elves [][]int) ([]int, int) {
	var calsPerElf []int
	for _, elf := range elves {
		var totalElfCals int
		for _, cal := range elf {
			totalElfCals += cal
		}
		calsPerElf = append(calsPerElf, totalElfCals)
	}

	// make a copy of the calsPerElf so that we have the original list
	var sortedCalsPerElf = make([]int, len(calsPerElf))
	copy(sortedCalsPerElf, calsPerElf) // copy the elements
	// sort the calsPerElf
	sort.Ints(sortedCalsPerElf)
	var mostCals int = sortedCalsPerElf[len(sortedCalsPerElf)-1]
	return calsPerElf, mostCals
}

func findElfNumber(calsPerElf []int, mostCals int) int {
	var elfNumber int
	for i, cals := range calsPerElf {
		if cals == mostCals {
			elfNumber = i + 1 // arrays count from 0, we want to count from 1 thus adding 1
			break
		}
	}
	return elfNumber
}
