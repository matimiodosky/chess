package validators

import "testing"

import . "board_game/commons"

func TestIsOpponent(t *testing.T) {
	game := NewGame(2, 2)
	piece := NewPiece(PAWN, WHITE)
	opponentPiece := NewPiece(PAWN, BLACK)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	game, err = game.Place(opponentPiece, NewCoordinate(0, 1))

	err = IsOpponent(game, NewCoordinate(0, 0), NewCoordinate(0, 1))
	if err != nil {
		t.Error(err)
	}
}

func TestIsNotOpponent(t *testing.T) {
	game := NewGame(2, 2)
	piece := NewPiece(PAWN, WHITE)
	game, err := game.Place(piece, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	err = IsOpponent(game, NewCoordinate(0, 0), NewCoordinate(0, 1))
	if err == nil {
		t.Error("Expected error")
	}
}
