package validators

import "testing"

import . "board_game/commons"

func TestForwardWhite(t *testing.T) {
	game := NewGame(2, 2)
	piece := NewPiece(PAWN, WHITE)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsForward(game, NewCoordinate(0, 0), NewCoordinate(0, 1))
	if err != nil {
		t.Error(err)
	}
}

func TestBackwardWhite(t *testing.T) {
	game := NewGame(2, 2)
	piece := NewPiece(PAWN, WHITE)
	game, err := game.Place(piece, NewCoordinate(0, 1))
	if err != nil {
		t.Error(err)
	}
	err = IsForward(game, NewCoordinate(0, 1), NewCoordinate(0, 0))
	if err == nil {
		t.Error("Expected error")
	}
}

func TestForwardBlack(t *testing.T) {
	game := NewGame(2, 2)
	piece := NewPiece(PAWN, BLACK)
	game, err := game.Place(piece, NewCoordinate(0, 1))
	if err != nil {
		t.Error(err)
	}
	err = IsForward(game, NewCoordinate(0, 1), NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
}

func TestBackwardBlack(t *testing.T) {
	game := NewGame(2, 2)
	piece := NewPiece(PAWN, BLACK)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsForward(game, NewCoordinate(0, 0), NewCoordinate(0, 1))
	if err == nil {
		t.Error("Expected error")
	}
}

func TestNotForward(t *testing.T) {
	game := NewGame(2, 2)
	piece := NewPiece(PAWN, WHITE)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsForward(game, NewCoordinate(0, 0), NewCoordinate(1, 0))
	if err == nil {
		t.Error("Expected error")
	}
}
