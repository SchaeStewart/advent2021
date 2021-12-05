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

func ParseInput(input []string) (callouts []int64, boards []BingoBoard) {
	boards = make([]BingoBoard, 0)

	calloutsStr := strings.Split(input[0], ",")
	for _, c := range calloutsStr {
		n, err := strconv.ParseInt(c, 10, 64)
		if err != nil {
			panic(err)
		}
		callouts = append(callouts, n)
	}

	currentBoard := []string{}
	for i := 2; i < len(input); i++ {
		if input[i] == "" {
			boards = append(boards, *NewBingoBoard(currentBoard))
			currentBoard = []string{}
		} else {
			currentBoard = append(currentBoard, input[i])
		}
	}
	boards = append(boards, *NewBingoBoard(currentBoard))

	return
}

type BingoCell struct {
	Number int64
	Called bool
}

type BingoBoard [][]BingoCell

func (b *BingoBoard) Callout(number int64) {
	for i, row := range *b {
		for j, item := range row {
			if item.Number == number {
				(*b)[i][j].Called = true
			}
		}
	}
}

func (b *BingoBoard) Score() int {
	sum := 0
	for _, row := range *b {
		for _, item := range row {
			if !item.Called {
				sum += int(item.Number)
			}
		}
	}
	return sum
}

func (bb *BingoBoard) Bingo() bool {
	b := *bb
	for i := 0; i < 5; i++ {
		if b[0][i].Called &&
			b[1][i].Called &&
			b[2][i].Called &&
			b[3][i].Called &&
			b[4][i].Called {
			return true
		}

		if b[i][0].Called &&
			b[i][1].Called &&
			b[i][2].Called &&
			b[i][3].Called &&
			b[i][4].Called {
			return true
		}
	}
	return false
}

func NewBingoBoard(input []string) *BingoBoard {
	board := make(BingoBoard, len(input))
	for i, line := range input {
		ns := strings.Split(line, " ")
		if board[i] == nil {
			board[i] = make([]BingoCell, 0)
		}
		for _, n := range ns {
			if n == "" || n == " " {
				continue
			}
			x, err := strconv.ParseInt(n, 10, 64)
			if err != nil {
				panic(err)
			}
			board[i] = append(board[i], BingoCell{
				Number: x,
				Called: false,
			})
		}
	}
	return &board
}

func RunGame(callouts []int64, boards []BingoBoard) (winningBoard *BingoBoard, lastCallout int64) {
	for _, callout := range callouts {
		for _, board := range boards {
			board.Callout(callout)
			if board.Bingo() {
				winningBoard = &board
				lastCallout = callout
				return
			}
		}
	}
	return nil, 0
}

func RunAllGames(callouts []int64, boards []BingoBoard) (winningBoards []*BingoBoard, lastCallout int64) {
	winningBoards = make([]*BingoBoard, 0, len(boards))
	for _, callout := range callouts {
		for _, board := range boards {
			if board.Bingo() { // Expensive
				continue
			}
			board.Callout(callout)
			if board.Bingo() {
				winningBoards = append(winningBoards, &board)
				lastCallout = callout
			}
			if len(winningBoards) == len(boards) {
				return
			}
		}
	}
	return nil, 0
}

func main() {
	input, err := ReadInput()
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	callouts, boards := ParseInput(input)
	winningBoard, lastCallout := RunGame(callouts, boards)
	for _, b := range boards {
		fmt.Println()
		for j, r := range b {
			fmt.Println(j, r)
		}
	}
	fmt.Println("Part1:", int64(winningBoard.Score())*lastCallout)

	callouts, boards = ParseInput(input)
	winningBoards, lastCallout := RunAllGames(callouts, boards)
	lastWinningBoard := winningBoards[len(winningBoards)-1]
	fmt.Println("Part2:", int64(lastWinningBoard.Score())*lastCallout)

}
