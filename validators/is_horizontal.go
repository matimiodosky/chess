package validators

import . "board_game/commons"

func IsHorizontal(game Game, from Coordinate, to Coordinate) error {
	if from.Row() == to.Row() {
		return nil
	} else {
		return InvalidMoveError("Not Horizontal")
	}
}
