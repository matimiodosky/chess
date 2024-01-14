package pieces

import (
	. "board_game/commons"
	. "board_game/validators"
)

var forwardMove = And(
	IsKind(PAWN),
	IsForward,
	IsSteps(1),
	IsEmpty,
)

var diagonalMove = And(
	IsKind(PAWN),
	IsForward,
	IsSteps(1),
	IsOpponent,
	IsDiagonal,
)

var doubleForwardMove = And(
	IsKind(PAWN),
	IsForward,
	IsEmpty,
	IsSteps(2),
	IsFirstMove,
)

var PawnValidator = Or(forwardMove, diagonalMove, doubleForwardMove)
