package commons

import "testing"

var True = func(game Game, from Coordinate, to Coordinate) error {
	return nil
}

var False = func(game Game, from Coordinate, to Coordinate) error {
	return InvalidMoveError("False")
}

var cases = []struct {
	validator1 Validator
	function   func(...Validator) Validator
	validator2 Validator
	result     bool
}{
	{True, And, True, true},
	{True, And, False, false},
	{False, And, True, false},
	{False, And, False, false},
	{True, Or, True, true},
	{True, Or, False, true},
	{False, Or, True, true},
	{False, Or, False, false},
}

func TestAndOr(t *testing.T) {
	for _, c := range cases {
		validator := c.function(c.validator1, c.validator2)
		if validator(Game{}, Coordinate{0, 0}, Coordinate{0, 0}) != nil && c.result {
			t.Errorf("Expected %v, got %v", c.result, !c.result)
		} else if validator(Game{}, Coordinate{0, 0}, Coordinate{0, 0}) == nil && !c.result {
			t.Errorf("Expected %v, got %v", c.result, !c.result)
		}
	}
}
