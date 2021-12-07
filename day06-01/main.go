package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func ReadInput() ([]string, error) {
	rawInput, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return nil, err
	}
	inputSplit := strings.Split(string(rawInput), "\n")
	return inputSplit, nil
}

func ParseInput(input []string) []*Fish{
	numbers := strings.Split(input[0], ",")
	result := make([]*Fish, len(numbers))
	for i, x := range numbers {
		n, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			panic(err)
		}
		result[i] = NewFish(int(n))
	}
	return result
}

type Fish struct {
	Timer int
}

func NewFish(timer int) *Fish {
	return &Fish{timer}
}

func (f *Fish) Day() (*Fish) {
	f.Timer -= 1
	if f.Timer < 0 {
		f.Timer = 6
		return NewFish(8)
	}
	return nil
}

func SimulateDaysAsync(days int, fishes []*Fish) int {
	if days == 0 {
		// fmt.Println("Early return")
		return len(fishes) 
	}
	counts := make(chan int)
	for i, fish := range fishes {
		// run async
		go func (fish *Fish, i int) {
			// fmt.Println("Fish #", i)
			// fmt.Println("simulating day", days)
			guppie := fish.Day()
			bowl := []*Fish{fish}
			if guppie != nil {
				bowl = append(bowl, guppie)
			}
			counts <- SimulateDaysAsync(days-1, bowl)
		}(fish,i)
	}
	sum := 0
	for i := 0; i < len(fishes); i++ {
		sum += <- counts
	}
	return sum
}

func SimulateDaysRescursive(days int, fishes []*Fish) int {
	if days == 0 {
		return len(fishes) 
	}
	counts := 0
	for _, fish := range fishes {
			guppie := fish.Day()
			bowl := []*Fish{fish}
			if guppie != nil {
				bowl = append(bowl, guppie)
			}
			counts += SimulateDaysRescursive(days-1, bowl)
		}
	return counts
}

 
func SimulateDays(days int, fishes []*Fish) []*Fish {
	for i := 0; i < days; i++ {
		fishBowl := []*Fish{}
		for _, fish := range fishes {
			fishBowl = append(fishBowl, fish.Day())
		}
		for _, fish := range fishBowl {
			if fish != nil {
				fishes = append(fishes, fish)
			}
		}
	}
	return fishes
}

func rotateLeft(arr []int) []int {
	result := make([]int, len(arr))
	first := arr[0]
	for i := 1; i < len(arr); i++ {
		result[i-1] = arr[i]
	}
	result[len(result)-1] = first
	return result 
}

func SimulateDaysInPlace(days int, fishes []int) int {
	for i := 0; i < days; i++ {
		fishes = rotateLeft(fishes)
		fishes[6] += fishes[8]
	}
	sum := 0
	for _, f := range fishes {
		sum += f
	}
	return sum	
}

func main() {
	input, err := ReadInput()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fish := ParseInput(input)

	start := time.Now()
	fishes := SimulateDays(80, fish)
	end := time.Now()
	totalTime := end.Sub(start)
	fmt.Println("Part 1 time", totalTime)
	fmt.Println("Part 1", len(fishes))

	// start = time.Now()
	// fish = ParseInput(input)
	// fishes = SimulateDays(256, fish)
	// // totalFish := SimulateDaysAsync(80, fish)
	// end = time.Now()
	// totalTime = end.Sub(start)
	// fmt.Println("Part 2 time", totalTime)

	fish = ParseInput(input)
	inPlaceFish := make([]int, 9)
	for _, f := range fish {
		inPlaceFish[f.Timer]++
	}

	totalFish := SimulateDaysInPlace(256, inPlaceFish)
	fmt.Println("Part 2", totalFish)
}
