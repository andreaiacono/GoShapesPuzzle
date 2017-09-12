package main

import (
	"github.com/gotk3/gotk3/gtk"
	"time"
	"log"
)

// returns next solution
func solver(puzzle *Puzzle, win *gtk.Window) {

	grid := createEmptyGrid(puzzle.Grid)
	puzzle.Grid = grid

	solvePuzzle(puzzle.Grid, puzzle.Pieces, puzzle, win, 0)
	win.QueueDraw()
	log.Printf("Finished solving.")
}

func solvePuzzle(grid [][]uint8, remainingPieces []Piece, puzzle *Puzzle, win *gtk.Window, calls int) bool {

	if ! puzzle.Computing {
		return false
	}

	puzzle.Grid = grid
	//log.Printf("solve with grid: %v, remain: %v, calls: %d", grid, remainingPieces, calls)
	win.QueueDraw()
	time.Sleep(time.Duration(puzzle.Speed) * time.Millisecond)
	if len(remainingPieces) == 0 {
		return true
	}

	// this loop is for starting placing an always different piece
	for i := 1; i < len(remainingPieces); i++ {
		for _, piece := range remainingPieces {
			minPieceSize := minPieceSize(remainingPieces)
			result, updatedGrid := placePiece(piece, grid, minPieceSize)
			if result {
				remainingPieces = removePieceFromRemaining(remainingPieces, piece)
				result := solvePuzzle(updatedGrid, remainingPieces, puzzle, win, calls+1)
				if result {
					return true
				} else {
					updatedGrid = removeShapeFromGrid(updatedGrid, piece.Number)
					remainingPieces = append(remainingPieces, piece) // append([]Piece{piece}, remainingPieces...)
				}
			}
		}
	}
	return false
}

func removePieceFromRemaining(pieces []Piece, piece Piece) []Piece {
	for i, v := range pieces {
		if v.Number == piece.Number {
			return append(pieces[:i], pieces[i+1:]...)
		}
	}
	return pieces
}

// placePiece checks if is there room for this piece (or one of its rotations)
// and if true add the piece to the grid, otherwise return false
func placePiece(piece Piece, grid [][]uint8, minPieceSize int) (bool, [][]uint8) {

	// consider all possible rotations of this piece
	for _, rot := range piece.Rotations {

		// loops over all possible cells where to place this piece
		for i := 0; i <= len(grid)-len(rot); i++ {
			for j := 0; j <= len(grid[0])-len(rot[0]); j++ {
				if grid[i][j] == EMPTY && pieceFits(rot, i, j, grid) && !hasLeftUnfillableAreas(grid, minPieceSize) {
					//log.Printf("Piece %v could be placed in %v at %d,%d", rot, grid, i, j)
					return true, addShapeToGrid(rot, i, j, grid, piece.Number)
				}
			}
		}
	}
	return false, grid
}

func hasLeftUnfillableAreas(grid [][]uint8, minPieceSize int) bool {
	var gridCopy = copyGrid(grid)
	var min = 10000
	for i := 0; i < len(gridCopy); i++ {
		for j := 0; j < len(gridCopy[0]); j++ {
			if gridCopy[i][j] == EMPTY {
				var area = getAreaSize(&gridCopy, i, j)
				if min > area {
					min = area
				}
			}
		}
	}
	return min < minPieceSize
}

func getAreaSize(grid *[][]uint8, x, y int) int {
	(*grid)[x][y] = FLOOD_FILL_VALUE
	size := 1

	if y > 0 && (*grid)[x][y-1] == EMPTY {
		size += getAreaSize(grid, x, y-1)
	}
	if x > 0 && (*grid)[x-1][y] == EMPTY {
		size += getAreaSize(grid, x-1, y)
	}
	if x < len((*grid))-1 && (*grid)[x+1][y] == EMPTY {
		size += getAreaSize(grid, x+1, y)
	}
	if y < len((*grid)[0])-1 && (*grid)[x][y+1] == EMPTY {
		size += getAreaSize(grid, x, y+1)
	}

	return size
}

func pieceFits(shape Shape, dx, dy int, grid [][]uint8) bool {

	for i := 0; i < len(shape); i++ {
		for j := 0; j < len(shape[0]); j++ {
			if shape[i][j] != EMPTY && grid[i+dx][j+dy] != EMPTY {
				return false
			}
		}
	}
	return true
}

// addShapeToGrid writes
func addShapeToGrid(shape Shape, dx, dy int, grid [][]uint8, number uint8) [][]uint8 {

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

func removeShapeFromGrid(grid [][]uint8, number uint8) [][]uint8 {

	updatedGrid := copyGrid(grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == number {
				updatedGrid[i][j] = EMPTY
			}
		}
	}
	return updatedGrid
}

func createEmptyGrid(grid [][]uint8) [][]uint8 {

	var copiedGrid = make([][]uint8, len(grid))
	for i := 0; i < len(grid); i++ {
		copiedGrid[i] = make([]uint8, len(grid[0]))
	}
	return copiedGrid
}

func copyGrid(grid [][]uint8) [][]uint8 {

	copiedGrid := createEmptyGrid(grid)

	var i, j int
	for i = 0; i < len(grid); i++ {
		for j = 0; j < len(grid[0]); j++ {
			copiedGrid[i][j] = grid[i][j]
		}
	}
	return copiedGrid
}
