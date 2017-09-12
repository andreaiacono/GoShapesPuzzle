package main

type Puzzle struct {
	Pieces       []Piece
	Grid         [][]uint8
	MaxPieceSide int8
	Speed        uint
	Computing    bool
	MinPieceSize int
	DrawNumbers	 bool
}
