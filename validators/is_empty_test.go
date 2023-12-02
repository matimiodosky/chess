package validators

import "testing"
import . "board_game/commons"

func TestIsEmpty(t *testing.T) {
	game := NewGame(2, 2)
	err := IsEmpty(game, NewCoordinate(0, 0), NewCoordinate(0, 1))
	if err != nil {
		t.Error(err)
	}
}

func TestNotEmpty(t *testing.T) {
	piece := NewPiece(KING, WHITE)
	game := NewGame(2, 2)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsEmpty(game, NewCoordinate(0, 1), NewCoordinate(0, 0))
	if err == nil {
		t.Error("Expected error")
	}
}
