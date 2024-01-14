package validators

import (
	. "board_game/commons"
)

func IsFirstMove(game Game, from Coordinate, _ Coordinate) error {
	hasMoved, err := game.PieceAtHasMoved(from)
	if err != nil {
		return err
	}
	if hasMoved {
		return InvalidMoveError("Piece has already moved")
	} else {
		return nil
	}
}
