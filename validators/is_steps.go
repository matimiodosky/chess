package validators

import (
	. "board_game/commons"
)

func IsSteps(count int) func(Game, Coordinate, Coordinate) error {
	return func(_ Game, coordinate Coordinate, coordinate2 Coordinate) error {
		// check horizontal
		if coordinate.Row() == coordinate2.Row() {
			if coordinate.Column()-coordinate2.Column() == count || coordinate2.Column()-coordinate.Column() == count {
				return nil
			}
		}
		// check vertical
		if coordinate.Column() == coordinate2.Column() {
			if coordinate.Row()-coordinate2.Row() == count || coordinate2.Row()-coordinate.Row() == count {
				return nil
			}
		}
		// check diagonal
		if coordinate.Column()-coordinate2.Column() == count || coordinate2.Column()-coordinate.Column() == count {
			if coordinate.Row()-coordinate2.Row() == count || coordinate2.Row()-coordinate.Row() == count {
				return nil
			}
		}
		return InvalidMoveError("Not specified steps")
	}
}
