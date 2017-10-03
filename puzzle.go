package main

import (
	"github.com/gotk3/gotk3/gtk"
)

// represents a puzzle
type Puzzle struct {
	Pieces       []Piece
	OriginalGrid [][]uint8
	Grid		 [][]uint8
	MaxPieceSide int8
	MinPieceSize int
	Solutions    *[][][]uint8
	IsRunning    bool
	HasGui       bool
	WinInfo      *WinInfo
}

// represent the info needed for displaying the puzzle in a GUI
type WinInfo struct {
	MainWindow  *gtk.Window
	StatusBar   gtk.Statusbar
	SolveButton gtk.Button
	ProgressBar gtk.ProgressBar
	Speed       uint
	DrawNumbers bool
}
