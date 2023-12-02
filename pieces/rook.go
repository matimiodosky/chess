package pieces

import . "board_game/commons"
import . "board_game/validators"

var horizontal = And(
	IsKind(ROOK),
	IsHorizontal,
	//NoObstacles, // TODO
	Or(IsEmpty, IsOpponent),
)

var vertical = And(
	IsKind(ROOK),
	IsVertical,
	//NoObstacles,  // TODO
	Or(IsEmpty, IsOpponent),
)

var RookValidator = Or(horizontal, vertical)
