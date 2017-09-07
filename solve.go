package main

import (
	"github.com/gotk3/gotk3/gtk"
	"time"
	"log"
)

var found = false

// returns next solution
func solve(puzzle *Puzzle, win *gtk.Window) {

	grid := createEmptyGrid(puzzle.Grid)
	puzzle.Grid = grid

	solvePuzzle(puzzle.Grid, puzzle.Pieces, puzzle, win)
}

func solvePuzzle(grid [][]int8, remainingPieces []Piece, puzzle *Puzzle, win *gtk.Window) {

	puzzle.Grid = grid
	win.QueueDraw()
	time.Sleep(200 * time.Millisecond)
	if len(remainingPieces) == 0 {
		found = true
		return
	}

	for _, piece := range remainingPieces {

		result, updatedGrid := placePiece(piece, grid)
		if result {
			remainingPieces = remainingPieces[1:]
			solvePuzzle(updatedGrid, remainingPieces, puzzle, win)
		} else {
			// if this piece cannot be placed into the grid
			// we can stop this branch
			return
		}
		//append(remainingPieces, piece)
	}

}

// placePiece checks if is there room for this piece (or one of its rotations)
// and if true add the piece to the grid, otherwise return false
func placePiece(piece Piece, grid [][]int8) (bool, [][]int8) {

	for i := 0; i < len(grid)-piece.MaxX; i++ {
		for j := 0; j < len(grid[0])-piece.MaxY; j++ {
			if pieceFits(piece, i, j, grid) {
				return true, addPieceToGrid(piece, i, j, grid)
			}
		}
	}

	return false, grid
}

func pieceFits(piece Piece, dx, dy int, grid [][]int8) bool {

	//log.Printf("grid: %v dx %d, dy %d", grid, dx, dy)
	for i := 0; i < len(piece.Shape); i++ {
		for j := 0; j < len(piece.Shape[0]); j++ {
			if piece.Shape[i][j] != EMPTY && grid[i+dx][j+dy] != EMPTY {
				return false
			}
		}
	}
	log.Printf("adding %v at %d, %d - %v", piece.Shape, dx, dy, grid)

	return true
}

// addPieceToGrid writes
func addPieceToGrid(piece Piece, dx, dy int, grid [][]int8) [][]int8 {

	updatedGrid := copyGrid(grid)
	for i := 0; i < len(piece.Shape); i++ {
		for j := 0; j < len(piece.Shape[0]); j++ {
			if piece.Shape[i][j] != EMPTY {
				updatedGrid[dx+i][dy+j] = piece.Number
			}
		}
	}

	return updatedGrid
}

func createEmptyGrid(grid [][]int8) [][]int8 {

	var copiedGrid = make([][]int8, len(grid))
	for i := 0; i < len(grid); i++ {
		copiedGrid[i] = make([]int8, len(grid[0]))
	}

	return copiedGrid
}

func copyGrid(grid [][]int8) [][]int8 {

	copiedGrid := createEmptyGrid(grid)

	var i, j int
	for i = 0; i < len(grid); i++ {
		for j = 0; j < len(grid[0]); j++ {
			copiedGrid[i][j] = grid[i][j]
		}
	}
	return copiedGrid
}
