package main

import (
	"github.com/gotk3/gotk3/gtk"
)

// a grid represents the frame in which to put
// the pieces
type Grid [][]uint8

// represents a puzzle
type Puzzle struct {
	Pieces       []Piece
	OriginalGrid Grid
	WorkingGrid  Grid
	MaxPieceSide int8
	MinPieceSize int
	Solutions    *[]Grid
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
