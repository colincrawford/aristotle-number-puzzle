package puzzle

func getAllPositions() [19]BoardPosition {
	positions := [19]BoardPosition{}

	inx := 0
	// 5 row from index 1 - 5
	for i := 1; i <= 5; i++ {
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

func getNextPosition(position *BoardPosition) (bool, BoardPosition) {
	rowLen := getRowLen(position.Row)
	incPos := position.Position + 1
	if incPos == rowLen {
		if position.Row == 5 {
			return false, position
		}
		return true, BoardPosition{position.Row + 1, 0}
	}
	return true, BoardPosition{position.Row, incPos}
}

func SolvePuzzle() Board {
	board := InitBoard()
	_, solvedBoard := solve(board, BoardPosition{1, 0})
	return solvedBoard
}

func solve(board *Board, position *BoardPosition) (bool, Board) {
	hasNextMove, nextPosition := getNextPosition(position)

	if !hasNextMove {
		return true, board
	}

	validMoves := GetValidMoves(board, position)
	for _, move := range validMoves {
		updatedBoard := SetPosition(board, position, move)
		solved, board := solve(updatedBoard, nextPosition)
		if solved {
			return true, board
		}
	}

	return false, board
}
