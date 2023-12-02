package validators

import (
	. "board_game/commons"
)

func IsOpponent(game Game, coordinate Coordinate, coordinate2 Coordinate) error {
	piece, err := game.PieceAt(coordinate)
	if err != nil {
		return err
	}
	piece2, err := game.PieceAt(coordinate2)
	if err != nil {
		return err
	}
	if piece.Color() != piece2.Color() {
		return nil
	} else {
		return InvalidMoveError("Not Opponent")
	}
}
