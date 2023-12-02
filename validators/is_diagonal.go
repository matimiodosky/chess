package validators

import (
	. "board_game/commons"
	"math"
)

func IsDiagonal(_ Game, coordinate Coordinate, coordinate2 Coordinate) error {

	rowDifference := coordinate.Row() - coordinate2.Row()
	columnDifference := coordinate.Column() - coordinate2.Column()

	if math.Abs(float64(rowDifference)) == math.Abs(float64(columnDifference)) {
		return nil
	} else {
		return InvalidMoveError("Not Diagonal")
	}

}
