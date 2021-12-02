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

type Instruction struct {
	Direction string
	Value     int
}

type Sub struct {
	Vert int
	Horz int
	Aim  int
}

func (s *Sub) MoveSimple(ins Instruction) {
	if ins.Direction == "forward" {
		s.Horz += ins.Value
	} else if ins.Direction == "up" {
		s.Vert -= ins.Value
	} else if ins.Direction == "down" {
		s.Vert += ins.Value
	} else {
		fmt.Errorf("Invalid direction %s", ins.Direction)
	}
}

func (s *Sub) Move(ins Instruction) {
	if ins.Direction == "forward" {
		s.Horz += ins.Value
		s.Vert += ins.Value * s.Aim
	} else if ins.Direction == "up" {
		s.Aim -= ins.Value
	} else if ins.Direction == "down" {
		s.Aim += ins.Value
	} else {
		fmt.Errorf("Invalid direction %s", ins.Direction)
	}
}

func InputToInstructions(input []string) ([]Instruction, error) {
	result := make([]Instruction, len(input))
	for i, l := range input {
		line := strings.Split(l, " ")
		ins := line[0]
		v := line[1]
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result[i] = Instruction{
			ins,
			n,
		}
	}
	return result, nil
}

func main() {
	input, err := ReadInput()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	instructions, err := InputToInstructions(input)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	sub := &Sub{}
	for _, ins := range instructions {
		sub.MoveSimple(ins)
	}
	fmt.Println("Part 1: vertical X horizontal", sub.Horz*sub.Vert)

	sub2 := &Sub{}
	for _, ins := range instructions {
		sub2.Move(ins)
	}
	fmt.Println("Part 2: vertical X horizontal", sub2.Horz*sub2.Vert)
}
