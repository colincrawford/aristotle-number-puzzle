package main

import (
	"github.com/colinlcrawford/aristotle-number-puzzle/puzzle"
)

func main() {
	board := puzzle.SolvePuzzle()
	puzzle.PrintBoard(board)
}
