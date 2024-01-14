package validators

import (
	"testing"
)
import . "board_game/commons"

func TestHasNotMoved(t *testing.T) {
	game := NewGame(2, 2)
	piece := NewPiece(PAWN, WHITE)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsFirstMove(game, NewCoordinate(0, 0), NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
}

func TestHasMoved(t *testing.T) {
	game := NewGame(2, 2, IsForward)
	piece := NewPiece(PAWN, WHITE)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	game, err = game.Move(NewCoordinate(0, 0), NewCoordinate(0, 1))
	if err != nil {
		t.Error(err)
	}
	err = IsFirstMove(game, NewCoordinate(0, 1), NewCoordinate(0, 2))
	if err == nil {
		t.Error("Expected error")
	}
}

// test for when there is not piece in the origin
func TestHasNotMovedNoPiece(t *testing.T) {
	game := NewGame(2, 2)
	err := IsFirstMove(game, NewCoordinate(0, 0), NewCoordinate(0, 0))
	if err == nil {
		t.Error("Expected error")
	}
}
