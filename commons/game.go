package commons

type Game struct {
	board      Board
	validators []Validator
}

func NewGame(width int, height int, validators ...Validator) Game {
	bord := NewBoard(width, height)
	return Game{bord, validators}
}

func (game Game) Move(from Coordinate, to Coordinate) (Game, error) {
	newBoard, err := game.board.Move(from, to)
	if err != nil {
		return game, err
	}
	return Game{newBoard, game.validators}, nil
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
	return Game{newBoard, game.validators}, nil
}
