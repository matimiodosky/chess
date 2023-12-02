package validators

import . "board_game/commons"

func IsVertical(game Game, from Coordinate, to Coordinate) error {
	if from.Column() == to.Column() {
		return nil
	} else {
		return InvalidMoveError("Not Vertical")
	}
}
