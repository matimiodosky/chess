package validators

import (
	. "board_game/commons"
)

func IsKind(kind PieceKind) func(game Game, from Coordinate, to Coordinate) error {
	return func(game Game, from Coordinate, to Coordinate) error {
		piece, err := game.PieceAt(from)
		if err != nil {
			return err
		} else if piece.Kind() == kind {
			return nil
		} else {
			return InvalidMoveError("Not of specified kind")
		}
	}
}
