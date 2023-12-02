package validators

import (
	. "board_game/commons"
)

func IsEmpty(game Game, _ Coordinate, coordinate2 Coordinate) error {
	if game.IsEmptyAt(coordinate2) {
		return nil
	} else {
		return InvalidMoveError("Not Empty")
	}
}
