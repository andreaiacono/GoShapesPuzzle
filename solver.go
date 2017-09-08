package main

import (
	"github.com/gotk3/gotk3/gtk"
	"time"
	//"log"
	"log"
)

var found = false

// returns next solution
func solver(puzzle *Puzzle, win *gtk.Window) {

	grid := createEmptyGrid(puzzle.Grid)
	puzzle.Grid = grid

	solvePuzzle(puzzle.Grid, puzzle.Pieces, puzzle, win, 0)
	log.Printf("Finished solving.")
}

func solvePuzzle(grid [][]int8, remainingPieces []Piece, puzzle *Puzzle, win *gtk.Window, calls int) bool {

	puzzle.Grid = grid
	//log.Printf("solve with grid: %v, remain: %v, calls: %d", grid, remainingPieces, calls)
	win.QueueDraw()
	time.Sleep(time.Duration(puzzle.Speed) * time.Millisecond)
	if len(remainingPieces) == 0 {
		found = true
		return true
	}

	// this loop is for starting placing an always different piece
	for i := 1; i < len(remainingPieces); i++ {
		for _, piece := range remainingPieces {
			//log.Printf("%s Trying piece %v", tabs(calls), piece)
			result, updatedGrid := placePiece(piece, grid)
			if result {
				remainingPieces = removePieceFromRemaining(remainingPieces, piece)
				result := solvePuzzle(updatedGrid, remainingPieces, puzzle, win, calls+1)
				if result {
					return true
				} else {
					updatedGrid = removeShapeFromGrid(updatedGrid, piece.Number)
					remainingPieces = append(remainingPieces, piece )// append([]Piece{piece}, remainingPieces...)
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

//func tabs(count int) string {
//	s := ""
//	for i := 0; i < count; i++ {
//		s += "\t"
//	}
//	return s
//}

// placePiece checks if is there room for this piece (or one of its rotations)
// and if true add the piece to the grid, otherwise return false
func placePiece(piece Piece, grid [][]int8) (bool, [][]int8) {

	// consider all possibile rotation of this piece
	for index, rot := range piece.Rotations {

		// loops over all possibile cells where to place this piece
		for i := 0; i <= len(grid)-len(rot); i++ {
			for j := 0; j <= len(grid[0])-len(rot[0]); j++ {
				if grid[i][j] == EMPTY && pieceFits(rot, i, j, grid) {
					//log.Printf("Piece %v could be placed in %v at %d,%d", piece.Rotations[index], grid, i, j)
					return true, addShapeToGrid(piece.Rotations[index], i, j, grid, piece.Number)
				}
			}
		}
	}
	//log.Printf("Piece %v could NOT be placed in %v", piece.Shape, grid)
	return false, grid
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

func removeShapeFromGrid(grid [][]int8, number int8) [][]int8 {

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
