package validators

import (
	. "board_game/commons"
	"testing"
)

func TestIs(t *testing.T) {

	piece := NewPiece(PAWN, WHITE)
	game := NewGame(2, 2)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsKind(PAWN)(game, NewCoordinate(0, 0), NewCoordinate(0, 1))
	if err != nil {
		t.Error(err)
	}

}

func TestIsNot(t *testing.T) {

	piece := NewPiece(KING, WHITE)
	game := NewGame(2, 2)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsKind(PAWN)(game, NewCoordinate(0, 0), NewCoordinate(0, 1))
	if err == nil {
		t.Error("Expected error")
	}
}
