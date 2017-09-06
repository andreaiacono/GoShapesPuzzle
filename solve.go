package main

import (
	"github.com/gotk3/gotk3/gtk"
	"time"
)

var found = false

// returns next solution
//func solve(puzzle Puzzle) {
//
//	remainingPieces := puzzle.Pieces
//	//grid := createEmptyGrid(puzzle.Grid)
//
	//solvePuzzle(puzzle.Grid, remainingPieces)
//}

func solvePuzzle(grid [][]int8, remainingPieces []Piece, puzzle *Puzzle, win *gtk.Window) {

	puzzle.Grid = grid
	win.QueueDraw()
	time.Sleep(100*time.Millisecond)
	if len(remainingPieces) == 0 {
		found = true
		return
	}

	for _, piece := range remainingPieces {

		result, updatedGrid := checkPiece(piece, grid)
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

func checkPiece(piece Piece, grid [][]int8) (bool, [][]int8) {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if pieceFits(piece, i, j, grid) {
				return true, placePiece(piece, i, j, grid)
			}
		}
	}

	return false, grid
}

func placePiece(piece Piece, dx, dy int, grid [][]int8) [][]int8 {

	updatedGrid := copyGrid(grid)
	for i := 0; i < len(piece.Shape); i++ {
		for j := 0; j < len(piece.Shape[0]); j++ {
			if piece.Shape[i][j] != EMPTY {
				updatedGrid[i][j] = piece.Number
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

func pieceFits(piece Piece, dx, dy int, grid [][]int8) bool {

	//log.Printf("grid: %v dx %d, dy %d", grid, dx, dy)
	for i := 0; i < len(piece.Shape); i++ {
		for j := 0; j < len(piece.Shape[0]); j++ {
			if dx+i < len(grid) && dy+j < len(grid[0]) && piece.Shape[i][j] != EMPTY && grid[i+dx][j+dy] != EMPTY {
				return false
			}
		}
	}

	return true
}
