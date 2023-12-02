package commons

import (
	"testing"
)

func TestMovement(t *testing.T) {

	board := NewBoard(2, 2)
	piece := NewPiece(PAWN, WHITE)
	from := NewCoordinate(0, 0)
	to := NewCoordinate(0, 1)

	board, err := board.Place(piece, from)
	if err != nil {
		t.Errorf("error placing piece")
	}
	board, err = board.Move(from, to)
	if err != nil {
		t.Errorf("error mooving piece")
	}
	piece, err = board.PieceAt(to)
	if err != nil {
		t.Errorf("error getting piece")
	}
	if piece.Kind() != PAWN {
		t.Errorf("piece not moved")
	}

	res := board.IsEmptyAt(from)
	if !res {
		t.Errorf("piece not removed from origin")
	}

}

func TestMovementWithMissingPiece(t *testing.T) {

	board := NewBoard(2, 2)
	from := NewCoordinate(0, 0)
	to := NewCoordinate(0, 1)
	board, err := board.Move(from, to)
	if err == nil {
		t.Errorf("should have errored mooving piece")
	}

}

func TestPlaceShouldBeImmutable(t *testing.T) {

	board := NewBoard(2, 2)
	piece := NewPiece(PAWN, WHITE)
	coordinate := NewCoordinate(0, 1)
	_, err := board.Place(piece, coordinate)

	if err != nil {
		t.Errorf("error placing piece")
	}

	if _, ok := board.pieces[coordinate]; ok {
		t.Errorf("piece placed in original board")
	}

}

func TestMoveShouldBeImmutable(t *testing.T) {

	board := NewBoard(2, 2)
	piece := NewPiece(PAWN, WHITE)
	from := NewCoordinate(0, 0)
	to := NewCoordinate(0, 1)

	board, err := board.Place(piece, from)
	if err != nil {
		t.Errorf("error placing piece")
	}
	_, err = board.Move(from, to)
	if err != nil {
		t.Errorf("error mooving piece")
	}
	res := board.IsEmptyAt(from)
	if res {
		t.Errorf("piece removed from origin")
	}

}

func TestNegativeCoordinates(t *testing.T) {

	board := NewBoard(2, 2)
	piece := NewPiece(PAWN, WHITE)
	from := NewCoordinate(-1, 0)
	to := NewCoordinate(0, 1)

	_, err := board.Place(piece, from)
	if err == nil {
		t.Errorf("should have errored placing piece")
	}
	_, err = board.Move(from, to)
	if err == nil {
		t.Errorf("should have errored mooving piece")
	}

}
