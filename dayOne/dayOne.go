package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// read file in ./input.txt
	input, err := os.Open("./input.txt")
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
		fmt.Println(err)
	}

	var calsPerElf []int
	for _, elf := range elves {
		var totalElfCals int
		for _, cal := range elf {
			totalElfCals += cal
		}
		calsPerElf = append(calsPerElf, totalElfCals)
	}

	// sort the calsPerElf
	// make a copy of the calsPerElf so that we have the original list
	var sortedCalsPerElf = make([]int, len(calsPerElf))
	copy(sortedCalsPerElf, calsPerElf) // copy the elements
	// sort the calsPerElf
	sort.Ints(sortedCalsPerElf)

	var mostCals int = sortedCalsPerElf[len(sortedCalsPerElf)-1]

	// find the elf number
	var elfNumber int
	for i, cals := range calsPerElf {
		if cals == mostCals {
			elfNumber = i + 1 // arrays count from 0, we want to count from 1 thus adding 1
			break
		}
	}
	fmt.Printf("most calories: %d\n", mostCals)
	fmt.Printf("elf number: %d\n", elfNumber)
}
