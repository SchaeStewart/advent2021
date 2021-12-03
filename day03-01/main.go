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

func RuneMode(runes []rune) rune {
	counts := make(map[rune]int)
	for _, r := range runes {
		counts[r]++
	}

	max := 0
	var mode rune
	for r, count := range counts {
		if count > max {
			max = count
			mode = r
		}
	}
	return mode
}

func BinaryRuneMode(runes []rune) rune {
	counts := make(map[rune]int)
	for _, r := range runes {
		counts[r]++
	}
	if counts['1'] >= counts['0'] {
		return '1'
	} else {
		return '0'
	}
}

func InverseBinaryRuneMode(runes []rune) rune {
	counts := make(map[rune]int)
	for _, r := range runes {
		counts[r]++
	}
	if counts['1'] < counts['0'] {
		return '1'
	} else {
		return '0'
	}
}

func FlipBits(bits string) string {
	result := ""
	for _, r := range bits {
		if r == '0' {
			result += "1"
		} else if r == '1' {
			result += "0"
		} else {
			result += string(r)
		}
	}
	return result
}

func GetCounts(input []string) [][]rune {
	occurrences := make([][]rune, len(input[0]))
	for _, line := range input {
		for i, c := range line {
			if ok := occurrences[i]; ok == nil {
				occurrences[i] = make([]rune, 0)
			}
			occurrences[i] = append(occurrences[i], c)
		}
	}
	return occurrences
}

// Gamma, Epsilon
func FindRates(input []string) (gamma, epsilon, oxygen, c02 string) {
	occurrences := GetCounts(input)
	for _, counts := range occurrences {
		gamma += string(RuneMode(counts))
	}
	epsilon = FlipBits(gamma)

	oxygenRates := input
	for i := 0; i < len(occurrences); i++ {
		counts := GetCounts(oxygenRates)[i]
		criteria := BinaryRuneMode(counts)
		oxygenRates = FilterByCriteria(oxygenRates, string(criteria), i)
		if len(oxygenRates) == 1 {
			oxygen = oxygenRates[0]
			break
		}
	}

	c02Rates := input
	for i := 0; i < len(occurrences); i++ {
		counts := GetCounts(c02Rates)[i]
		critera := FlipBits(string(BinaryRuneMode(counts)))
		c02Rates = FilterByCriteria(c02Rates, critera, i)
		if len(c02Rates) == 1 {
			c02 = c02Rates[0]
			break
		}
	}
	return
}

func FilterByCriteria(input []string, criteria string, criteriaIndex int) []string {
	result := make([]string, 0)
	for _, line := range input {
		if string(line[criteriaIndex]) == criteria {
			result = append(result, line)
		}
	}
	return result
}

func BinaryToInt(str string) int64 {
	i, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		// fmt.Errorf(err.Error())
		panic(err)
	}
	return i
}

func main() {
	input, err := ReadInput()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	gammaStr, epsilonStr, oxygenStr, c02Str := FindRates(input)
	fmt.Println("Oxygen", oxygenStr)
	gamma := BinaryToInt(gammaStr)
	epsilon := BinaryToInt(epsilonStr)
	oxygen := BinaryToInt(oxygenStr)
	c02 := BinaryToInt(c02Str)
	fmt.Println("Part1: Gamma X Epsilon", gamma*epsilon)
	fmt.Println("Part2: Oxygen X C02", oxygen*c02)
}
