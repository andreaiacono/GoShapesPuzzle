package main

type Puzzle struct {
	Pieces       []Piece
	Grid         [][]uint8
	MaxPieceSide int8
	Speed        uint8
	Computing    bool
	MinPieceSize int
}
