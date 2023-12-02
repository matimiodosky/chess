package pieces

import "testing"
import . "board_game/commons"

func TestHorizontalMovement(t *testing.T) {

	rook := NewPiece(ROOK, WHITE)
	game := NewGame(3, 3, RookValidator)
	game, err := game.Place(rook, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	game, err = game.Move(NewCoordinate(0, 0), NewCoordinate(2, 0))
	if err != nil {
		t.Error(err)
	}

}

func TestVerticalMovement(t *testing.T) {

	rook := NewPiece(ROOK, WHITE)
	game := NewGame(3, 3, RookValidator)
	game, err := game.Place(rook, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	game, err = game.Move(NewCoordinate(0, 0), NewCoordinate(0, 2))
	if err != nil {
		t.Error(err)
	}
}

func TestDiagonalMovement(t *testing.T) {

	rook := NewPiece(ROOK, WHITE)
	game := NewGame(3, 3, RookValidator)
	game, err := game.Place(rook, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	game, err = game.Move(NewCoordinate(0, 0), NewCoordinate(2, 2))
	if err == nil {
		t.Error("Expected error")
	}
}
