package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadInput() ([]string, error) {
	rawInput, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return nil, err
	}
	inputSplit := strings.Split(string(rawInput), "\n")
	return inputSplit, nil
}

func Abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func Min(ns ...int) int {
	min := ns[0]
	for _, n := range ns {
		if n < min {
			min = n
		}
	}
	return min
}

func Max(ns ...int) int {
	max := ns[0]
	for _, n := range ns {
		if n > max {
			max = n
		}
	}
	return max
}

func ParseInput(input []string) []int {
	numbers := strings.Split(input[0], ",")
	result := make([]int, len(numbers))
	for i, x := range numbers {
		n, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			panic(err)
		}
		result[i] = int(n)
	}
	return result
}

func CalculateCost(crabs []int, position int) int {
	cost := 0
	for _, crab := range crabs {
		cost += Abs(crab-position)
	}
	return cost
}

func realCost(a, b int) int {
	cost := 0
	start := Min(a, b)
	end := Max(a, b)
	for i := 1; i <= end-start; i++ {
		cost += i
	}
	return cost
}

func CalculateRealCost(crabs []int, postion int) int {
	cost := 0
	for _, crab := range crabs {
		cost += realCost(crab, postion)
	}
	return cost
}

type CostCalculator func ([]int, int) int

func FindLowestCost(crabs []int, calculator CostCalculator) int {
	costs := make([]int, 0)
	limit := Max(crabs...)
	for i := 0; i < limit; i++ {
		cost := calculator(crabs, i)
		costs = append(costs, cost)
	}
	lowest := Min(costs...) 
	return lowest
}

func main() {
	input, err := ReadInput()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	crabs := ParseInput(input)
	lowest := FindLowestCost(crabs, CalculateCost)
	fmt.Println("Part 1", lowest)

	crabs = ParseInput(input)
	lowest = FindLowestCost(crabs, CalculateRealCost)
	fmt.Println("Part 2", lowest)

}
