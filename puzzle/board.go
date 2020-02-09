package puzzle

import "fmt"

// A single game piece
// values are 1 - 19
type GamePiece struct {
	Filled bool
	Value  int
}

// A specific spot in the board
type BoardPosition struct {
	Row    int
	Column int
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
	Rows [5][]GamePiece
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
	spaces := []string{"  ", " ", "", " ", "  "}
	str := ""
	for i, row := range board.Rows {
		str += fmt.Sprintf("%s%s\n", spaces[i], rowToString(row))
	}
	return str
}

func PrintBoard(board *Board) {
	fmt.Println(BoardToString(board))
}

func InitBoard() Board {
	board := Board{}
	for i := 0; i < 5; i++ {
		if i == 0 || i == 4 {
			board.Rows[i] = make([]GamePiece, 3)
		} else if i == 1 || i == 3 {
			board.Rows[i] = make([]GamePiece, 4)
		} else {
			board.Rows[i] = make([]GamePiece, 5)
		}
	}
	return board
}

func getPosition(board *Board, position *BoardPosition) GamePiece {
	row := getRow(board, position.Row)
	return row[position.Column]
}

func getRow(board *Board, n int) []GamePiece {
	return board.Rows[n]
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
	return row[0:(position.Column + 1)]
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
	previousPosition := position.Column + 1
	for (previousRow > 0) && (previousPosition < previousRowLength) {
		previousPieces = append(previousPieces, getPosition(board, &BoardPosition{previousRow, previousPosition}))
		previousRow = position.Row - 1
		previousRowLength = len(getRow(board, previousRow))
		previousPosition = position.Column + 1
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
	previousPosition := position.Column - 1
	for (previousRow > 0) && (previousPosition >= 0) {
		previousPieces = append(previousPieces, getPosition(board, &BoardPosition{previousRow, previousPosition}))
		previousRow = position.Row - 1
		previousPosition = position.Column - 1
	}
	return previousPieces
}

func getUsedPieces(board *Board) map[int]bool {
	seenPieces := make(map[int]bool)
	for _, row := range board.Rows {
		for _, piece := range row {
			seenPieces[piece.Value] = true
		}
	}
	return seenPieces
}

func rowSum(pieces *[]GamePiece) int {
	total := 0
	for _, piece := range *pieces {
		total += piece.Value
	}
	return total
}

func getPiecesLte(min int) []GamePiece {
	pieces := []GamePiece{}
	for i := MinPieceValue; i <= min; i++ {
		pieces = append(pieces, GamePiece{true, i})
	}
	return pieces
}

func max(nums ...int) int {
	n := nums[0]
	for _, num := range nums[1:] {
		if num > n {
			n = num
		}
	}
	return n
}

// Get the valid pieces for a board position
// given the pieces currently in a board
func GetValidMoves(board *Board, position *BoardPosition) []GamePiece {
	usedPieces := getUsedPieces(board)

	// get the sum of each of the three rows for this position
	horizontalRowPieces := getPrevHorizontalRowPieces(board, position)
	horizontalRowSum := rowSum(&horizontalRowPieces)
	leftDiagRowPieces := getPrevLeftDiagRowPieces(board, position)
	leftDiagRowSum := rowSum(&leftDiagRowPieces)
	rightDiagRowPieces := getPrevRightDiagRowPieces(board, position)
	rightDiagRowSum := rowSum(&rightDiagRowPieces)

	// the highest of those constrains the possible pieces in this position
	highestRowSum := max(horizontalRowSum, rightDiagRowSum, leftDiagRowSum)
	possibleMoves := getPiecesLte(TargetRowSum - highestRowSum)

	// remove the moves we've already used
	validMoves := []GamePiece{}
	for _, possibleMove := range possibleMoves {
		if usedPieces[possibleMove.Value] {
			continue
		}
		validMoves = append(validMoves, possibleMove)
	}

	return validMoves
}

func SetPosition(board *Board, position *BoardPosition, move *GamePiece) *Board {
	board.Rows[position.Row][position.Column] = GamePiece{true, move.Value}
	return board
}
