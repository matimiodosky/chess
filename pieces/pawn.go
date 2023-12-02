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
	IsSteps(2),
	IsEmpty,
	//IsFirstMove, // TODO
)

var doubleForwardMove = And(
	IsKind(PAWN),
	IsForward,
	IsSteps(1),
	IsDiagonal,
	IsOpponent,
)

var PawnValidator = Or(forwardMove, diagonalMove, doubleForwardMove)
