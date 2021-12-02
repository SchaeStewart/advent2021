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

func InputToInt(input []string) ([]int, error) {
	result := make([]int, len(input))
	for i, line := range input {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		result[i] = n
	}
	return result, nil
}

func FindIncrements(numbers []int) int {
	increased := 0
	previous := numbers[0]
	for _, val := range numbers {
		if val > previous {
			increased++
		}
		previous = val
	}
	return increased
}

func getSumOfWindow(numbers []int, window, offset int) int {
	sum := 0
	for i := offset; i < offset+window && i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func FindSlidingIncrements(numbers []int, window int) int {
	increased := 0
	previous := getSumOfWindow(numbers, window, 0)
	for i := 1; i < len(numbers); i++ {
		current := getSumOfWindow(numbers, window, i)
		if current > previous {
			increased++
		}
		previous = current
	}
	return increased
}

func main() {
	input, err := ReadInput()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	numbers, err := InputToInt(input)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	increments := FindIncrements(numbers)
	fmt.Println("Single increments:", increments)

	increments = FindSlidingIncrements(numbers, 3)
	fmt.Println("Sliding increments:", increments)

	fmt.Println("a", getSumOfWindow(numbers, 3, 0))
	fmt.Println("b", getSumOfWindow(numbers, 3, 1))
}
