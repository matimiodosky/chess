package validators

import "testing"
import . "board_game/commons"

func TestForwardRight(t *testing.T) {
	game := NewGame(2, 2)
	err := IsDiagonal(game, NewCoordinate(0, 0), NewCoordinate(1, 1))
	if err != nil {
		t.Error(err)
	}
}

func TestBackwardLeft(t *testing.T) {

	game := NewGame(2, 2)
	err := IsDiagonal(game, NewCoordinate(1, 1), NewCoordinate(0, 0))
	if err != nil {
		t.Error(err)
	}
}

func TestForwardLeft(t *testing.T) {

	game := NewGame(2, 2)
	err := IsDiagonal(game, NewCoordinate(0, 1), NewCoordinate(1, 0))
	if err != nil {
		t.Error(err)
	}
}

func TestBackwardRight(t *testing.T) {

	game := NewGame(2, 2)
	err := IsDiagonal(game, NewCoordinate(1, 0), NewCoordinate(0, 1))
	if err != nil {
		t.Error(err)
	}
}

func TestNotDiagonal(t *testing.T) {

	game := NewGame(2, 2)
	err := IsDiagonal(game, NewCoordinate(0, 0), NewCoordinate(0, 1))
	if err == nil {
		t.Error("Expected error")
	}
}
