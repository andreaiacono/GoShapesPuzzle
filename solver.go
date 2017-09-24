package main

import (
	"github.com/gotk3/gotk3/gtk"
	"time"
	"fmt"
	"log"
)

// solves the puzzle
func Solver(puzzle *Puzzle, win *gtk.Window) {

	grid := createEmptyGrid(puzzle.Grid)
	puzzle.Grid = grid

	solvePuzzle(puzzle.Grid, puzzle.Pieces, puzzle, win)
	if len(*puzzle.Solutions) > 0 {
		puzzle.Grid = (*puzzle.Solutions)[0]
	}
	message := fmt.Sprintf("Finished solving. Found %d solutions.", len(*puzzle.Solutions))
	if puzzle.UseGui {
		win.QueueDraw()
		puzzle.StatusBar.Push(1, message)
	} else {
		log.Println(message)
	}
}

// the recursive solving function
func solvePuzzle(grid [][]uint8, remainingPieces []Piece, puzzle *Puzzle, win *gtk.Window) {

	if ! puzzle.Computing {
		return
	}

	puzzle.Grid = grid
	if puzzle.UseGui {
		time.Sleep(time.Duration(puzzle.Speed) * time.Millisecond)
		win.QueueDraw()
	}
	if len(remainingPieces) == 0 {
		addSolution(puzzle.Solutions, grid, puzzle.StatusBar, puzzle.UseGui)
		return
	}
	minPieceSize := minPieceSize(remainingPieces)

	// loops over the remaining pieces
	for _, piece := range remainingPieces {

		// considers all possible rotations of this piece
		for _, rot := range piece.Rotations {

			// tries every cell of the grid
			for j := 0; j <= len(grid[0])-len(rot[0]); j++ {
				for i := 0; i <= len(grid)-len(rot); i++ {

					// if the cell is empty anf the piece doesn't overlap with other pieces
					if grid[i][j] == EMPTY && pieceFits(rot, i, j, grid) {

						// adds the piece to the grid
						updatedGrid := addShapeToGrid(rot, i, j, grid, piece.Number)

						// if the piece doesn't leave any unfillable cell
						if !hasLeftUnfillableAreas(updatedGrid, minPieceSize) {

							// updates the remaining pieces
							remainingPieces, index = removePieceFromRemaining(remainingPieces, piece)

							// recursively calls this function
							solvePuzzle(updatedGrid, remainingPieces, puzzle, win)

							// after having tried, remove this piece and goes on
							updatedGrid = removeShapeFromGrid(updatedGrid, piece.Number)
							//remainingPieces = append(remainingPieces, piece)
							remainingPieces = append(remainingPieces[:index], append([]Piece{piece}, remainingPieces[index:]...)...)
						}
					}
				}
			}
		}
	}
}

// TODO
// create a map with the placed pieces and positions, and check if the chosen configuraiton
// has been already computed

func addSolution(solutions *[][][]uint8, solution [][]uint8, statusBar gtk.Statusbar, useGui bool) [][][]uint8 {

	for sol := range *solutions {
		if areEqualPieces((*solutions)[sol], solution) {
			return *solutions
		}
	}
	*solutions = append(*solutions, solution)

	if useGui {
		statusBar.Push(1, fmt.Sprintf("Found %d solutions", len(*solutions)))
	} else {
		log.Printf("Solution #%d: %v",  len(*solutions), solution)
	}
	return *solutions
}

func removePieceFromRemaining(pieces []Piece, piece Piece) ([]Piece, int) {

	for i, v := range pieces {
		if v.Number == piece.Number {
			return append(pieces[:i], pieces[i+1:]...), i
		}
	}
	// FIXME maybe a log.panic?
	return pieces, -1
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
