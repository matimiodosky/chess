package commons

type Game struct {
	board      Board
	validators []Validator
	history    []Move
}

func NewGame(width int, height int, validators ...Validator) Game {
	bord := NewBoard(width, height)
	return Game{bord, validators, []Move{}}
}

func (game Game) Move(from Coordinate, to Coordinate) (Game, error) {
	piece, _ := game.PieceAt(from)
	move := Move{from, to, piece.id}
	newBoard, err := game.board.Move(from, to)
	newHistory := append(game.history, move)
	err = runValidators(game.validators, game, from, to)
	if err != nil {
		return game, err
	}
	return Game{newBoard, game.validators, newHistory}, nil
}

func runValidators(validators []Validator, game Game, from Coordinate, to Coordinate) error {
	for _, validator := range validators {
		err := validator(game, from, to)
		if err == nil {
			return nil
		}
	}
	return InvalidMoveError("Invalid move")
}

func (game Game) PieceAt(coordinate Coordinate) (Piece, error) {
	piece, err := game.board.PieceAt(coordinate)
	return piece, err
}

func (game Game) IsEmptyAt(coordinate Coordinate) bool {
	return game.board.IsEmptyAt(coordinate)
}

func (game Game) Place(piece Piece, coordinate Coordinate) (Game, error) {
	newBoard, err := game.board.Place(piece, coordinate)
	if err != nil {
		return Game{}, err
	}
	return Game{newBoard, game.validators, game.history}, nil
}

func (game Game) PieceAtHasMoved(coordinate Coordinate) (bool, error) {

	piece, err := game.board.PieceAt(coordinate)
	if err != nil {
		return false, err
	}

	for _, move := range game.history {
		if move.pieceId == piece.id {
			return true, nil
		}
	}

	return false, nil
}
