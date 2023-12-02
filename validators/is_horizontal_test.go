package validators

import "testing"

import . "board_game/commons"

func TestIsHorizontal(t *testing.T) {

	game := NewGame(2, 2)
	piece := NewPiece(ROOK, WHITE)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsHorizontal(game, NewCoordinate(0, 0), NewCoordinate(1, 0))
	if err != nil {
		t.Error(err)
	}

}

func TestIsNotHorizontal(t *testing.T) {

	game := NewGame(2, 2)
	piece := NewPiece(ROOK, WHITE)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsHorizontal(game, NewCoordinate(0, 0), NewCoordinate(1, 1))
	if err == nil {
		t.Error("Expected error")
	}
}
