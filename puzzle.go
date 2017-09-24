package main

import (
	"math/rand"
	"github.com/gotk3/gotk3/gtk"
)

type Puzzle struct {
	Pieces       []Piece
	Grid         [][]uint8
	MaxPieceSide int8
	Speed        uint
	Computing    bool
	MinPieceSize int
	DrawNumbers  bool
	StatusBar    gtk.Statusbar
	Solutions	 *[][][]uint8
	UseGui		 bool
}

func (puzzle *Puzzle) ShufflePieces() {
	shuffledPieces := make([]Piece, len(puzzle.Pieces))
	perm := rand.Perm(len(puzzle.Pieces))
	for i, v := range perm {
		shuffledPieces[v] = puzzle.Pieces[i]
	}
	puzzle.Pieces = shuffledPieces
}
