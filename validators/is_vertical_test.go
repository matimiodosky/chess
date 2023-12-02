package validators

import "testing"

import . "board_game/commons"

func TestIsVertical(t *testing.T) {

	game := NewGame(2, 2)
	piece := NewPiece(ROOK, WHITE)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsVertical(game, NewCoordinate(0, 0), NewCoordinate(0, 1))
	if err != nil {
		t.Error(err)
	}

}

func TestIsNotVertical(t *testing.T) {

	game := NewGame(2, 2)
	piece := NewPiece(ROOK, WHITE)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsVertical(game, NewCoordinate(0, 0), NewCoordinate(1, 1))
	if err == nil {
		t.Error("Expected error")
	}
}
