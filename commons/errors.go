package commons

import "errors"

var MissingPieceError = errors.New("no piece found at coordinate")

var CoordinateOutsideOfBoundaries = errors.New("coordinate outside of boundaries")

func InvalidMoveError(message string) error {
	return errors.New(message)
}
