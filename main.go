package main

import (
	"fmt"
	"strings"
)

type Position struct {
	row, col int
}

const TARGET_NUM = 38

var DIAGS = [][]Position{
	// top right to bottom left
	{Position{0, 0}, Position{1, 0}, Position{2, 0}},
	{Position{0, 1}, Position{1, 1}, Position{2, 1}, Position{3, 0}},
	{Position{0, 2}, Position{1, 2}, Position{2, 2}, Position{3, 1}, Position{4, 0}},
	{Position{1, 3}, Position{2, 3}, Position{3, 2}, Position{4, 1}},
	{Position{2, 4}, Position{3, 3}, Position{4, 2}},

	// top left to bottom right
	{Position{2, 0}, Position{3, 0}, Position{4, 0}},
	{Position{1, 0}, Position{2, 1}, Position{3, 1}, Position{4, 1}},
	{Position{0, 0}, Position{1, 1}, Position{2, 2}, Position{3, 2}, Position{4, 2}},
	{Position{0, 1}, Position{1, 2}, Position{2, 3}, Position{3, 3}},
	{Position{0, 2}, Position{1, 3}, Position{2, 4}},
}

var rowLen = []int{3, 4, 5, 4, 3}

func previousPos(pos *Position) Position {
	if pos.row == 0 && pos.col == 0 {
		panic("Cannot move back from the first position")
	}
	row := pos.row
	col := pos.col
	if col == 0 {
		row -= 1
		col = rowLen[row] - 1
	} else {
		col -= 1
	}
	return Position{row, col}
}

func nextPos(pos *Position) Position {
	if pos.row == 4 && pos.col == 2 {
		panic("Cannot move forward from the last position")
	}
	row := pos.row
	col := pos.col
	if col == (rowLen[row] - 1) {
		row += 1
		col = 0
	} else {
		col += 1
	}
	return Position{row, col}
}

func main() {
	board := NewBoard()
	board.Solve()
	fmt.Print(board.ToString())
}

type Board struct {
	rows       [][]int
	usedNums   map[int]bool
	currentPos Position
	IsFull     bool
}

func NewBoard() Board {
	board := Board{}
	board.rows = [][]int{
		{0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0},
	}
	board.usedNums = make(map[int]bool)
	board.currentPos = Position{0, 0}
	board.IsFull = false

	return board
}

func (board *Board) Solve() {
	for !board.IsFull {
		couldMove := board.NextMove()
		if !couldMove {
			board.Backtrack()
		}
	}
}

func (board *Board) Backtrack() {
	if board.currentPos.col == 0 && board.currentPos.row == 0 {
		fmt.Println(board.ToString())
		panic("Cannot backtrack from the first spot!")
	}

	pos := board.currentPos
	currVal := board.rows[pos.row][pos.col]

	board.usedNums[currVal] = false
	board.IsFull = false
	board.rows[pos.row][pos.col] = 0

	board.currentPos = previousPos(&board.currentPos)
}

func (board *Board) NextMove() bool {
	if board.IsFull {
		panic("No next move at the last position")
	}

	pos := board.currentPos
	currNum := board.rows[pos.row][pos.col]

	for i := currNum + 1; i < 20; i++ {
		if board.usedNums[i] {
			continue
		}

		pos = board.currentPos
		currNum = board.rows[pos.row][pos.col]
		board.rows[pos.row][pos.col] = i
		board.usedNums[i] = true
		board.usedNums[currNum] = false

		if board.IsValid() {
			if board.currentPos.row == 4 && board.currentPos.col == 2 {
				board.IsFull = true
			} else {
				board.currentPos = nextPos(&pos)
			}
			return true
		}
	}

	return false
}

func (board *Board) IsValid() bool {
	// check the rows
	for _, row := range board.rows {
		sum := 0
		for i, ele := range row {
			sum += ele
			if sum > TARGET_NUM {
				return false
			}
			if i == (len(row)-1) && ele > 0 && sum != TARGET_NUM {
				return false
			}
		}
	}

	// check the diags
	for _, positions := range DIAGS {
		sum := 0
		for _, pos := range positions {
			sum += board.rows[pos.row][pos.col]
			if sum > TARGET_NUM {
				return false
			}
		}
	}

	return true
}

func (board *Board) ToString() string {
	var sb strings.Builder
	sb.WriteString("----------------------\n")
	for _, row := range board.rows {
		rowLen := len(row)
		for i := 0; i < (5 - rowLen); i++ {
			sb.WriteString(" ")
		}
		for inx, ele := range row {
			sb.WriteString(fmt.Sprintf("%d", ele))
			if inx != rowLen {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}
	sb.WriteString("----------------------\n")
	return sb.String()
}
