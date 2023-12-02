package commons

import "maps"

type Board struct {
	width  int
	height int
	pieces map[Coordinate]Piece
}

func NewBoard(width int, height int) Board {
	return Board{width, height, make(map[Coordinate]Piece)}
}

func (b Board) Move(from Coordinate, to Coordinate) (Board, error) {

	err := validateCoordinates(b, from, to)
	if err != nil {
		return Board{}, err
	}

	if piece, ok := b.pieces[from]; !ok {
		return Board{}, MissingPieceError
	} else {
		newPieces := maps.Clone(b.pieces)
		delete(newPieces, from)
		newPieces[to] = piece
		return Board{b.width, b.height, newPieces}, nil
	}
}

func (b Board) Place(p Piece, c Coordinate) (Board, error) {
	err := validateCoordinate(c, b)
	if err != nil {
		return Board{}, err
	}
	newPieces := maps.Clone(b.pieces)
	newPieces[c] = p
	return Board{b.width, b.height, newPieces}, nil
}

func (b Board) PieceAt(coordinate Coordinate) (Piece, error) {
	if piece, ok := b.pieces[coordinate]; ok {
		return piece, nil
	}
	return Piece{}, nil
}

func (b Board) IsEmptyAt(coordinate Coordinate) bool {
	if _, ok := b.pieces[coordinate]; ok {
		return false
	}
	return true
}

func validateCoordinate(c Coordinate, b Board) error {
	if c.y > b.height || c.x > b.width || c.x < 0 || c.y < 0 {
		return CoordinateOutsideOfBoundaries
	}
	return nil
}

func validateCoordinates(b Board, coordinates ...Coordinate) error {
	for _, coordinate := range coordinates {
		err := validateCoordinate(coordinate, b)
		if err != nil {
			return err
		}
	}
	return nil
}
