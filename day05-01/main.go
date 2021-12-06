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

type Point struct {
	X, Y int
}

func ParsePoint(xY string) Point {
	split := strings.Split(xY, ",")
	x, err := strconv.ParseInt(split[0], 10, 64)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		panic(err)
	}
	return Point{int(x), int(y)}
}

type Line struct {
	A,
	B Point
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

// func Bresenham(a, b Point) []Point {
// 	x1, y1 := a.X, a.Y
// 	x2, y2 := b.X, b.Y
// 	points := make([]Point, 0)
// 	m := 2 * (y2 - y1)
// 	slopeError := m - (x2 - x1)
// 	for x, y := x1, y1; x <= x2; x++ {
// 		fmt.Println(x, y)
// 		points = append(points, Point{x,y})
// 		fmt.Println(slopeError, m)
// 		slopeError += m
// 		if slopeError >= 0 {
// 			y++
// 			slopeError -= 2 * (x2 - x1)
// 		}
// 	}
// 	return points
// }

func (p Point) SharesAxisWith(b Point) bool {
	return p.X == b.X || p.Y == b.Y
}

func LineWalkStep(a, b int) int {
	if a > b {
		return -1
	} else if a < b {
		return 1
	} else {
		return 0
	}
}

func Points(a, b Point) []Point {
	xDelta := LineWalkStep(a.X, b.X)
	yDelta := LineWalkStep(a.Y, b.Y)
	steps := Max(Abs((a.X - b.X)), Abs(a.Y-b.Y))

	points := []Point{a}
	for i := 1; i <= steps; i++ {
		last := points[i-1]
		points = append(points, Point{last.X + xDelta, last.Y + yDelta})
	}
	return points
}

func ParseInput(input []string) []Line {
	lines := make([]Line, 0)
	for _, row := range input {
		split := strings.Split(row, " -> ")
		a := ParsePoint(split[0])
		b := ParsePoint(split[1])
		lines = append(lines, Line{a, b})
	}
	return lines
}

func FindBoundaries(lines []Line) (maxX, maxY int) {
	xs := []int{}
	ys := []int{}
	for _, line := range lines {
		xs = append(xs, line.A.X)
		xs = append(xs, line.B.X)
		ys = append(ys, line.A.Y)
		ys = append(ys, line.B.Y)
	}
	maxX = Max(xs...)+1
	maxY = Max(ys...)+1
	return
}

type Plot [][]int

func (p *Plot) AddLines(lines []Line, includeDiagonals bool) {
	for _, l := range lines {
		if !includeDiagonals && !l.A.SharesAxisWith(l.B) {
			continue
		}
		points := Points(l.A, l.B)
		for _, point := range points {
			(*p)[point.Y][point.X]+=1
		}
	}
}

func (p *Plot) CountOverlap(threshold int) int {
	sum := 0
	for _, row := range *p {
		for _, point := range row {
			if point >= threshold {
				sum++
			}
		}
	}
	return sum
}

func NewPlot(x, y int) Plot {
	p := make(Plot, y)
	for iy := 0; iy < y; iy++ {
		p[iy] = make([]int, x)
	}
	return p
}

func main() {
	input, err := ReadInput()
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	lines := ParseInput(input)
	maxX, maxY := FindBoundaries(lines)
	plot := NewPlot(maxX, maxY)
	plot.AddLines(lines, false)
	overlap := plot.CountOverlap(2)
	fmt.Println("Part 1:", overlap)

	diagonalPlot := NewPlot(maxX, maxY)
	diagonalPlot.AddLines(lines, true)
	diagonalOverlap := diagonalPlot.CountOverlap(2)
	fmt.Println("Part 2:", diagonalOverlap)
}
