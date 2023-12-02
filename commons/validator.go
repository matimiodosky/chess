package commons

type Validator func(Game, Coordinate, Coordinate) error

func Or(validators ...Validator) Validator {
	return func(game Game, from Coordinate, to Coordinate) error {
		for _, validator := range validators {
			if err := validator(game, from, to); err == nil {
				return nil
			}
		}
		return InvalidMoveError("Invalid move")
	}
}

func And(validators ...Validator) Validator {
	return func(game Game, from Coordinate, to Coordinate) error {
		for _, validator := range validators {
			if err := validator(game, from, to); err != nil {
				return err
			}
		}
		return nil
	}
}
