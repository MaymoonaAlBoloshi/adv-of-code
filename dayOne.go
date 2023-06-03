// not doing things in cool ways, just trying to get it done while understainding different parts of go
package main

import (
  "fmt"
  "bufio"
  "os"
  "sort"
)

func main() {
  // read file in ./input.txt
  input, err :=  os.Open("./input.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer input.Close()


  scanner := bufio.NewScanner(input)

  // create an empty array
  var elves []int
  
  // make an array of all the numbers, empty lines are 0
  for scanner.Scan() {
    var elf int
    fmt.Sscanf(scanner.Text(), "%d", &elf)
    elves = append(elves, elf)
  }

  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }

  // calculate the calories for each elf , we know 0 starts a new elf
    
  var everyElfCalories []int
  var elfTotal int

  for i := 0; i < len(elves); i++ {
    if elves[i] !=0 {
      elfTotal += elves[i]
    } else {
      everyElfCalories = append(everyElfCalories, elfTotal)
      elfTotal = 0
    }
  }

  // now we have an array of all the calories for each elf
  // we need to sort it, where the greatest is first
  var sortedCalories []int = sort.IntSlice(everyElfCalories)

  // get the elf number from everyElfCalories
  var elfNumber int = sort.SearchInts(everyElfCalories, sortedCalories[0])
  fmt.Println("most calories : %d", sortedCalories[0])
  fmt.Println("elf number : %d", elfNumber)

}
