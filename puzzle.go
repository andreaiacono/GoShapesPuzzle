package main

import "math/rand"

type Puzzle struct {
	Pieces       []Piece
	Grid         [][]uint8
	MaxPieceSide int8
	Speed        uint
	Computing    bool
	MinPieceSize int
	DrawNumbers	 bool
}


func (puzzle *Puzzle) ShufflePieces() {
	shuffledPieces := make([]Piece, len(puzzle.Pieces))
	perm := rand.Perm(len(puzzle.Pieces))
	for i, v := range perm {
		shuffledPieces[v] = puzzle.Pieces[i]
	}
	puzzle.Pieces = shuffledPieces
}
