package pieces

import "testing"
import . "board_game/commons"

func TestWhiteForwardMovement(t *testing.T) {

	pawn := NewPiece(PAWN, WHITE)
	game := NewGame(3, 3)
	game, err := game.Place(pawn, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	game, err = game.Move(NewCoordinate(0, 0), NewCoordinate(0, 1))
	if err != nil {
		t.Error(err)
	}

}

func TestBlackForwardMovement(t *testing.T) {

	pawn := NewPiece(PAWN, BLACK)
	game := NewGame(3, 3)
	game, err := game.Place(pawn, NewCoordinate(0, 2))
	if err != nil {
		t.Error(err)
	}
	game, err = game.Move(NewCoordinate(0, 2), NewCoordinate(0, 1))
	if err != nil {
		t.Error(err)
	}
}

func TestWhiteDoubleForwardMoovement(t *testing.T) {

	pawn := NewPiece(PAWN, WHITE)
	game := NewGame(3, 3)
	game, err := game.Place(pawn, NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
	game, err = game.Move(NewCoordinate(0, 0), NewCoordinate(0, 2))
	if err != nil {
		t.Error(err)
	}

}

func TestBlackDoubleForwardMoovement(t *testing.T) {

	pawn := NewPiece(PAWN, BLACK)
	game := NewGame(3, 3)
	game, err := game.Place(pawn, NewCoordinate(0, 2))
	if err != nil {
		t.Error(err)
	}
	game, err = game.Move(NewCoordinate(0, 2), NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}

}
