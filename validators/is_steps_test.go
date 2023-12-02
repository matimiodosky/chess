package validators

import "testing"

import . "board_game/commons"

func TestVerticalIsSteps(t *testing.T) {
	game := NewGame(2, 2)
	piece := NewPiece(PAWN, WHITE)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsSteps(1)(game, NewCoordinate(0, 0), NewCoordinate(0, 1))
	if err != nil {
		t.Error(err)
	}
	err = IsSteps(2)(game, NewCoordinate(0, 0), NewCoordinate(0, 1))
	if err == nil {
		t.Error("Expected error")
	}
}

func TestHorizontalIsSteps(t *testing.T) {
	game := NewGame(2, 2)
	piece := NewPiece(PAWN, WHITE)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsSteps(1)(game, NewCoordinate(0, 0), NewCoordinate(1, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsSteps(2)(game, NewCoordinate(0, 0), NewCoordinate(1, 0))
	if err == nil {
		t.Error("Expected error")
	}
}

func TestDiagonalIsSteps(t *testing.T) {
	game := NewGame(2, 2)
	piece := NewPiece(PAWN, WHITE)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsSteps(1)(game, NewCoordinate(0, 0), NewCoordinate(1, 1))
	if err != nil {
		t.Error(err)
	}
	err = IsSteps(2)(game, NewCoordinate(0, 0), NewCoordinate(1, 1))
	if err == nil {
		t.Error("Expected error")
	}
}
