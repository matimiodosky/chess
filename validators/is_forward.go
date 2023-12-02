package validators

import (
	. "board_game/commons"
)

func IsForward(game Game, from Coordinate, to Coordinate) error {
	piece, err := game.PieceAt(from)
	if err != nil {
		return err
	} else if piece.Color() == WHITE && from.Row() < to.Row() {
		return nil
	} else if piece.Color() == BLACK && from.Row() > to.Row() {
		return nil
	} else {
		return InvalidMoveError("Not Forward")
	}
}
