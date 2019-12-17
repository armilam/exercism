// Package connect containts functions to determine the winner of a game of connect.
package connect

import (
	"strings"
)

// A Position represents a single position within a connect game and contains a reference to
// each of that Position's neighboring Positions.
type Position struct {
	Player      string
	Left        *Position
	TopLeft     *Position
	TopRight    *Position
	Right       *Position
	BottomRight *Position
	BottomLeft  *Position
}

// Board is an array of arrays of positions, representing a game of connect.
type Board struct {
	Positions [][]*Position
}

// ResultOf determines the winner of the connect game, given a board represented by an array
// of strings, each of which contains an O, X, or a . for each position.
func ResultOf(boardRows []string) (string, error) {
	board := ParseBoard(boardRows)

	if Xwins(board) {
		return "X", nil
	} else if Owins(board) {
		return "O", nil
	} else {
		return "", nil
	}
}

// Xwins determines whether X wins the game, given a Board.
func Xwins(board Board) bool {
	startingPosition := &(*board.Positions[0][0])

	// Find a position on the left side of the board held by X
	// and find out if that starts a path to the right side of
	// the board. If not, keep going down the left side of the
	// board to see if there is a different winning path.
	for startingPosition != nil {
		if XwinsFromPath(startingPosition, []*Position{}) {
			return true
		}

		startingPosition = (*startingPosition).BottomRight
	}

	return false
}

// XwinsFromPath determines whether X wins the connect game by recursively following paths of
// Positions held by X.
func XwinsFromPath(position *Position, path []*Position) bool {
	// If there is nothing here, we are off the board. Return false.
	if position == nil {
		return false
	}

	// Detect a circular path. If we've already been here, return false.
	if IsCircularPath(position, path) {
		return false
	}

	// If this position is not occupied by an X, the path has ended. Return false.
	if (*position).Player != "X" {
		return false
	}

	// If we are on the right side of the board with an X, then X won. Return true.
	if (*position).Right == nil {
		return true
	}

	// Recursively continue finding the path.
	if XwinsFromPath((*position).Left, append(path, position)) {
		return true
	} else if XwinsFromPath((*position).TopLeft, append(path, position)) {
		return true
	} else if XwinsFromPath((*position).TopRight, append(path, position)) {
		return true
	} else if XwinsFromPath((*position).Right, append(path, position)) {
		return true
	} else if XwinsFromPath((*position).BottomRight, append(path, position)) {
		return true
	} else if XwinsFromPath((*position).BottomLeft, append(path, position)) {
		return true
	}

	return false
}

// Owins determines whether O wins the game, given a Board.
func Owins(board Board) bool {
	startingPosition := &(*board.Positions[0][0])

	// Find a position on the top side of the board held by O
	// and find out if that starts a path to the bottom side of
	// the board. If not, keep going across the top side of the
	// board to see if there is a different winning path.
	for startingPosition != nil {
		if OwinsFromPath(startingPosition, []*Position{}) {
			return true
		}

		startingPosition = (*startingPosition).Right
	}

	return false
}

// OwinsFromPath determines whether O wins the connect game by recursively following paths of
// Positions held by O.
func OwinsFromPath(position *Position, path []*Position) bool {
	// If there is nothing here, we are off the board. Return false.
	if position == nil {
		return false
	}

	// Detect a circular path. If we've already been here, return false.
	if IsCircularPath(position, path) {
		return false
	}

	// If this position is not occupied by an O, the path has ended. Return false.
	if (*position).Player != "O" {
		return false
	}

	// If we are on the bottom side of the board with an O, then O won. Return true.
	if (*position).BottomLeft == nil && (*position).BottomRight == nil {
		return true
	}

	// Recursively continue finding the path.
	if OwinsFromPath((*position).Left, append(path, position)) {
		return true
	} else if OwinsFromPath((*position).TopLeft, append(path, position)) {
		return true
	} else if OwinsFromPath((*position).TopRight, append(path, position)) {
		return true
	} else if OwinsFromPath((*position).Right, append(path, position)) {
		return true
	} else if OwinsFromPath((*position).BottomRight, append(path, position)) {
		return true
	} else if OwinsFromPath((*position).BottomLeft, append(path, position)) {
		return true
	}

	return false
}

// IsCircularPath detects whether position is contained in path. If so, then it
// is a circular path and returns true. Otherwise, returns false.
func IsCircularPath(position *Position, path []*Position) bool {
	for i := 0; i < len(path); i++ {
		if position == path[i] {
			return true
		}
	}

	return false
}

// ParseBoard takes an array of strings representing a board and turns it into
// a two-dimensional array of Position pointers, each of which has pointers to
// its neighbors.
func ParseBoard(boardRows []string) Board {
	board := Board{}

	// Parse the board one row at a time.
	for rowIndex := 0; rowIndex < len(boardRows); rowIndex++ {
		row := strings.Split(boardRows[rowIndex], "")
		board.Positions = append(board.Positions, []*Position{})

		// Parse the row one column at a time.
		for columnIndex := 0; columnIndex < len(row); columnIndex++ {
			position := Position{Player: row[columnIndex]}

			// If this is not the leftmost column, connect to the left neighbor.
			if columnIndex > 0 {
				position.Left = board.Positions[rowIndex][columnIndex-1]
				board.Positions[rowIndex][columnIndex-1].Right = &position
			}

			// If this is not the topmost row, connect to the top left neighbor.
			if rowIndex > 0 {
				position.TopLeft = board.Positions[rowIndex-1][columnIndex]
				board.Positions[rowIndex-1][columnIndex].BottomRight = &position

				// If this is not the rightmost column, connect to the top right neighbor.
				if columnIndex < len(row)-1 {
					position.TopRight = board.Positions[rowIndex-1][columnIndex+1]
					board.Positions[rowIndex-1][columnIndex+1].BottomLeft = &position
				}
			}

			// Store the position in the board.
			board.Positions[rowIndex] = append(board.Positions[rowIndex], &position)
		}
	}

	return board
}
