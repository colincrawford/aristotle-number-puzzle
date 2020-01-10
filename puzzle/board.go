package puzzle

import "fmt"

// A single game piece
// values are 1 - 19
type GamePiece struct {
	Filled bool
	Value  int
}

// The game board is a hexagonal type shape
// with 19 total pieces.
// Ex:
//     1 2 3
//    4 5 6 7
//  8 9 10 11 12
//   12 13 14 15
//    16 17 18
type Board struct {
	Row1 [3]GamePiece
	Row2 [4]GamePiece
	Row3 [5]GamePiece
	Row4 [4]GamePiece
	Row5 [3]GamePiece
}

const (
	MinPieceValue = 1
	MaxPieceValue = 19
	TargetRowSum  = 38
)

func rowToString(row []GamePiece) string {
	rowStr := ""
	for i, piece := range row {
		rowStr += fmt.Sprintf("%v", piece.Value)
		if i != len(row) {
			rowStr += " "
		}
	}
	return rowStr
}

func BoardToString(board *Board) string {
	str := ""
	str += fmt.Sprintf("  %s\n", rowToString(board.Row1[:]))
	str += fmt.Sprintf(" %s\n", rowToString(board.Row2[:]))
	str += fmt.Sprintf("%s\n", rowToString(board.Row3[:]))
	str += fmt.Sprintf(" %s\n", rowToString(board.Row4[:]))
	str += fmt.Sprintf("  %s\n", rowToString(board.Row5[:]))
	return str
}

func PrintBoard(board *Board) {
	fmt.Println(BoardToString(board))
}

// A specific spot in the board
type BoardPosition struct {
	Row      int
	Position int
}

func InitBoard() Board {
	board := Board{}
	board.Row1 = [3]GamePiece{}
	board.Row2 = [4]GamePiece{}
	board.Row3 = [5]GamePiece{}
	board.Row4 = [4]GamePiece{}
	board.Row5 = [3]GamePiece{}
	return board
}

func getPosition(board *Board, position *BoardPosition) GamePiece {
	row := getRow(board, position.Row)
	return row[position.Position]
}

func getRowLen(row int) int {
	switch row {
	case 1, 5:
		return 3
	case 2, 4:
		return 4
	case 3:
		return 5
	default:
		panic(fmt.Sprintf("%v is not a valid row", row))
	}
}

func getRow(board *Board, n int) []GamePiece {
	switch n {
	case 1:
		return board.Row1[:]
	case 2:
		return board.Row2[:]
	case 3:
		return board.Row3[:]
	case 4:
		return board.Row4[:]
	case 5:
		return board.Row5[:]
	default:
		panic(fmt.Sprintf("%v is not a valid row", n))
	}
}

func removeEmptyPieces(pieces *[]GamePiece) []GamePiece {
	validPieces := []GamePiece{}
	for _, piece := range *pieces {
		if piece.Filled {
			validPieces = append(validPieces, piece)
		}
	}
	return validPieces
}

// The horizontal row for any given board position
func getPrevHorizontalRowPieces(board *Board, position *BoardPosition) []GamePiece {
	row := getRow(board, position.Row)
	return row[0:(position.Position + 1)]
}

// The diagonal row starting from the top and going down / left
// for a given board position
func getPrevLeftDiagRowPieces(board *Board, position *BoardPosition) []GamePiece {
	if position.Row == 1 {
		return []GamePiece{}
	}

	previousPieces := []GamePiece{}
	previousRow := position.Row - 1
	previousRowLength := len(getRow(board, previousRow))
	previousPosition := position.Position + 1
	for (previousRow > 0) && (previousPosition < previousRowLength) {
		previousPieces = append(previousPieces, getPosition(board, BoardPosition{previousRow, previousPosition}))
		previousRow = position.Row - 1
		previousRowLength = len(getRow(board, previousRow))
		previousPosition = position.Position + 1
	}
	return previousPieces
}

// The diagonal row starting from the top and going down / right
// for a given board position
func getPrevRightDiagRowPieces(board *Board, position *BoardPosition) []GamePiece {
	if position.Row == 1 {
		return []GamePiece{}
	}

	previousPieces := []GamePiece{}
	previousRow := position.Row - 1
	previousPosition := position.Position - 1
	for (previousRow > 0) && (previousPosition >= 0) {
		previousPieces = append(previousPieces, getPosition(board, BoardPosition{previousRow, previousPosition}))
		previousRow = position.Row - 1
		previousPosition = position.Position - 1
	}
	return previousPieces
}

func allPieces(board *Board) []GamePiece {
	rows := [][]GamePiece{board.Row1[:], board.Row2[:], board.Row3[:], board.Row4[:], board.Row5[:]}
	pieces := []GamePiece{}
	for _, row := range rows {
		pieces = append(pieces, removeEmptyPieces(&row)...)
	}
	return pieces
}

func rowSum(pieces *[]GamePiece) int {
	total := 0
	for _, piece := range *pieces {
		total += piece.Value
	}
	return total
}

// Get the valid pieces for a board position
// given the pieces currently in a board
func GetValidMoves(board *Board, position *BoardPosition) []GamePiece {
	previousPieces := allPieces(board)
	usedPieces := make(map[int]bool)
	for _, piece := range previousPieces {
		usedPieces[piece.Value] = true
	}
	horizontalRowSum := rowSum(getPrevHorizontalRowPieces(board, position))
	leftDiagRowSum := rowSum(getPrevLeftDiagRowPieces(board, position))
	rightDiagRowSum := rowSum(getPrevRightDiagRowPieces(board, position))
	validMoves := []GamePiece{}
	for i := MinPieceValue; i <= MaxPieceValue; i++ {
		if usedPieces[i] {
			continue
		}
		if (horizontalRowSum + i) > TargetRowSum {
			continue
		}
		if (leftDiagRowSum + i) > TargetRowSum {
			continue
		}
		if (rightDiagRowSum + i) > TargetRowSum {
			continue
		}
		validMoves = append(validMoves, GamePiece{true, i})
	}
	return validMoves
}

func SetPosition(board *Board, position *BoardPosition, move *GamePiece) Board {
	switch position.Row {
	case 1:
		board.Row1[position.Position] = *move
	case 2:
		board.Row2[position.Position] = *move
	case 3:
		board.Row3[position.Position] = *move
	case 4:
		board.Row4[position.Position] = *move
	case 5:
		board.Row5[position.Position] = *move
	}
	return board
}
