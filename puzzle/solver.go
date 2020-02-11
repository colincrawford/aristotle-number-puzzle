package puzzle

import (
	"fmt"
)

func getRowLen(rowInx int) int {
	if rowInx == 0 || rowInx == 4 {
		return 3
	}
	if rowInx == 1 || rowInx == 3 {
		return 4
	}
	if rowInx == 2 {
		return 5
	}
	panic(fmt.Sprintf("Invalid rowInx %d", rowInx))
}

func getAllPositions() [19]BoardPosition {
	positions := [19]BoardPosition{}

	inx := 0
	// 5 row from index 1 - 5
	for i := 0; i < 5; i++ {
		// the row length depends on it index
		rowLen := getRowLen(i)

		for j := 0; j < rowLen; j++ {
			positions[inx] = BoardPosition{i, j}
			inx++
		}
	}

	return positions
}

var positions = getAllPositions()

func getNextPosition(position *BoardPosition) (bool, *BoardPosition) {
	rowLen := getRowLen(position.Row)
	incPos := position.Column + 1
	if incPos == rowLen {
		if position.Row == 4 {
			return false, position
		}
		return true, &BoardPosition{position.Row + 1, 0}
	}
	return true, &BoardPosition{position.Row, incPos}
}

func SolvePuzzle() *Board {
	board := InitBoard()
	_, solvedBoard := solve(&board, &BoardPosition{0, 0})
	return solvedBoard
}

func solve(board *Board, position *BoardPosition) (bool, *Board) {
	hasNextMove, nextPosition := getNextPosition(position)
	println("============")
	fmt.Printf("Has next move %t\n", hasNextMove)
	fmt.Printf("At %d,%d -> next %d,%d\n", position.Row, position.Column, nextPosition.Row, nextPosition.Column)

	if !hasNextMove {
		return true, board
	}

	validMoves := GetValidMoves(board, position)
	fmt.Println("valid moves: ")
	fmt.Println(validMoves)
	fmt.Println("")

	for _, move := range validMoves {
		updatedBoard := SetPosition(board, position, &move)

		fmt.Printf("Next position: Row %d, Column %d\n", position.Row, position.Column)
		fmt.Printf("Trying %d\n", move.Value)
		PrintBoard(board)

		solved, board := solve(updatedBoard, nextPosition)
		if solved {
			return true, board
		}
	}

	return false, board
}
