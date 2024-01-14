package commons

import (
	"math/rand"
	"time"
)

type Piece struct {
	id    int
	kind  PieceKind
	color PieceColor
}

func (p Piece) Kind() PieceKind {
	return p.kind
}

func (p Piece) Color() PieceColor {
	return p.color
}

func randomId() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Int()
}

func NewPiece(kind PieceKind, color PieceColor) Piece {
	return Piece{randomId(), kind, color}
}

type PieceKind int

type PieceColor int

const (
	PAWN PieceKind = iota
	ROOK PieceKind = iota
	KING PieceKind = iota
)

const (
	WHITE PieceColor = iota
	BLACK PieceColor = iota
)
