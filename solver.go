package main

import (
	"github.com/gotk3/gotk3/gtk"
	"time"
)

var found = false

// returns next solution
func solver(puzzle *Puzzle, win *gtk.Window) {

	grid := createEmptyGrid(puzzle.Grid)
	puzzle.Grid = grid

	solvePuzzle(puzzle.Grid, puzzle.Pieces, puzzle, win)
}

func solvePuzzle(grid [][]int8, remainingPieces []Piece, puzzle *Puzzle, win *gtk.Window) bool {

	puzzle.Grid = grid
	//log.Printf("solve with grid: %v, remain: %v", grid, remainingPieces)
	win.QueueDraw()
	time.Sleep(time.Duration(puzzle.Speed) * time.Millisecond)
	if len(remainingPieces) == 0 {
		found = true
		return false
	}

	for i := 1; i < len(remainingPieces); i++ {
		var tmp = remainingPieces[i]
		remainingPieces[i] = remainingPieces[0]
		remainingPieces[0] = tmp
		for _, piece := range remainingPieces {
			result, updatedGrid := placePiece(piece, grid)
			if result {
				remainingPieces = remainingPieces[1:]
				return solvePuzzle(updatedGrid, remainingPieces, puzzle, win)
			} else {
				return false
			}
			//append(remainingPieces, piece)
		}
	}
	return false
}

// placePiece checks if is there room for this piece (or one of its rotations)
// and if true add the piece to the grid, otherwise return false
func placePiece(piece Piece, grid [][]int8) (bool, [][]int8) {

	for i := 0; i < len(grid)-piece.MaxX; i++ {
		for j := 0; j < len(grid[0])-piece.MaxY; j++ {
			result, index := piecesFit(piece, i, j, grid)
			if result {
				//log.Printf("Piece %v could be placed in %v at %d,%d", piece.Rotations[index], grid, i, j)
				return true, addShapeToGrid(piece.Rotations[index], i, j, grid, piece.Number)
			}
		}
	}

	//log.Printf("Piece %v could NOT be placed in %v", piece.Shape, grid)
	return false, grid
}

func piecesFit(piece Piece, dx, dy int, grid [][]int8) (bool, int) {

	for index, rot := range piece.Rotations {
		if pieceFits(rot, dx, dy, grid) {
			return true, index
		}
	}
	return false, -1
}

func pieceFits(shape Shape, dx, dy int, grid [][]int8) bool {

	//log.Printf("grid: %v dx %d, dy %d", grid, dx, dy)
	for i := 0; i < len(shape); i++ {
		for j := 0; j < len(shape[0]); j++ {
			if shape[i][j] != EMPTY && grid[i+dx][j+dy] != EMPTY {
				return false
			}
		}
	}
	//log.Printf("adding %v at %d, %d - %v", shape, dx, dy, grid)
	return true
}

// addShapeToGrid writes
func addShapeToGrid(shape Shape, dx, dy int, grid [][]int8, number int8) [][]int8 {

	updatedGrid := copyGrid(grid)
	for i := 0; i < len(shape); i++ {
		for j := 0; j < len(shape[0]); j++ {
			if shape[i][j] != EMPTY {
				updatedGrid[dx+i][dy+j] = number
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
