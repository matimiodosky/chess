package commons

type Piece struct {
	kind  PieceKind
	color PieceColor
}

func (p Piece) Kind() PieceKind {
	return p.kind
}

func (p Piece) Color() PieceColor {
	return p.color
}

func NewPiece(kind PieceKind, color PieceColor) Piece {
	return Piece{kind, color}
}

type PieceKind int

type PieceColor int

const (
	PAWN PieceKind = iota
	KING PieceKind = iota
)

const (
	WHITE PieceColor = iota
	BLACK PieceColor = iota
)
