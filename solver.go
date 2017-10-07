package main

import (
	"github.com/gotk3/gotk3/gtk"
	"time"
	"fmt"
	"log"
	"math"
)

// this map contains a string representation of the
var visited = map[string]bool{}
var totalStates float64
var actualStates float64

// solves the puzzle
func Solver(puzzle *Puzzle) {

	puzzle.WorkingGrid = createEmptyGrid(puzzle.OriginalGrid)
	defer elapsed(*puzzle)()

	visited = map[string]bool{}

	solvePuzzle(puzzle, puzzle.Pieces)
	if len(*puzzle.Solutions) > 0 {
		puzzle.WorkingGrid = (*puzzle.Solutions)[0]
	}
	puzzle.IsRunning = false
	if puzzle.HasGui {
		puzzle.WinInfo.SolveButton.SetLabel("Find solutions")
	}
}

// the recursive solving function
func solvePuzzle(puzzle *Puzzle, remainingPieces []Piece) {

	if ! puzzle.IsRunning {
		return
	}

	if puzzle.HasGui {
		time.Sleep(time.Duration(puzzle.WinInfo.Speed) * time.Millisecond)
		puzzle.WinInfo.MainWindow.QueueDraw()
	}

	if len(remainingPieces) == 0 {
		addSolution(puzzle.Solutions, puzzle.WorkingGrid, puzzle.WinInfo.StatusBar, puzzle.HasGui)
		return
	}
	minPieceSize := minPieceSize(remainingPieces)

	// loops over the remaining pieces
	for _, piece := range remainingPieces {

		// considers all possible rotations of this piece
		for _, rot := range piece.Rotations {

			// tries every cell of the grid (limited to the positions where
			// the piece is not outside the boundaries of the frame)
			for j := 0; j <= len(puzzle.WorkingGrid[0])-len(rot[0]); j++ {
				for i := 0; i <= len(puzzle.WorkingGrid)-len(rot); i++ {

					actualStates ++

					// if the cell is empty and the piece doesn't overlap with other pieces
					if puzzle.WorkingGrid[i][j] == EMPTY && pieceFits(rot, i, j, puzzle.WorkingGrid) {

						// adds the piece to the grid
						updatedGrid := addShapeToGrid(rot, i, j, puzzle.WorkingGrid, piece.Number)

						// checks for already visited states
						if checkAndUpdateVisitedState(updatedGrid) {
							continue
						}

						// if the piece doesn't leave any unfillable cell
						if !hasLeftUnfillableAreas(updatedGrid, minPieceSize) {

							// updates the remaining pieces
							index, remainingPieces := removePieceFromRemaining(remainingPieces, piece)

							// recursively calls this function
							puzzle.WorkingGrid = updatedGrid
							solvePuzzle(puzzle, remainingPieces)

							// after having tried, remove this piece and goes on
							updatedGrid = removeShapeFromGrid(updatedGrid, piece.Number)
							puzzle.WorkingGrid = updatedGrid
							remainingPieces = append(remainingPieces[:index], append([]Piece{piece}, remainingPieces[index:]...)...)
						}
					}
				}
			}
		}
	}
}

func checkAndUpdateVisitedState(grid Grid) bool {

	gridString := fmt.Sprintf("%s", grid)
	_, isPresent := visited[gridString]

	if !isPresent {
		visited[gridString] = true
	}
	return isPresent
}

func addSolution(solutions *[]Grid, solution Grid, statusBar gtk.Statusbar, useGui bool) []Grid {

	for sol := range *solutions {
		if areEqualPieces((*solutions)[sol], solution) {
			return *solutions
		}
	}
	*solutions = append(*solutions, solution)

	if useGui {
		statusBar.Push(1, fmt.Sprintf("Found %d solutions", len(*solutions)))
	} else {
		log.Printf("Solution #%d: %v", len(*solutions), solution)
	}
	return *solutions
}

func removePieceFromRemaining(pieces []Piece, piece Piece) (int, []Piece) {

	for i, v := range pieces {
		if v.Number == piece.Number {
			return i, append(pieces[:i], pieces[i+1:]...)
		}
	}

	log.Fatal("Trying to remove a not found piece from remaining ones.")
	return -1, pieces
}

func hasLeftUnfillableAreas(grid Grid, minPieceSize int) bool {

	var gridCopy = copyGrid(grid)
	var min = math.MaxInt32
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

func getAreaSize(grid *Grid, x, y int) int {

	(*grid)[x][y] = FLOOD_FILL_VALUE
	size := 1

	if y > 0 && (*grid)[x][y-1] == EMPTY {
		size += getAreaSize(grid, x, y-1)
	}
	if x > 0 && (*grid)[x-1][y] == EMPTY {
		size += getAreaSize(grid, x-1, y)
	}
	if x < len(*grid)-1 && (*grid)[x+1][y] == EMPTY {
		size += getAreaSize(grid, x+1, y)
	}
	if y < len((*grid)[0])-1 && (*grid)[x][y+1] == EMPTY {
		size += getAreaSize(grid, x, y+1)
	}
	return size
}

func pieceFits(shape Shape, dx, dy int, grid Grid) bool {

	for i := 0; i < len(shape); i++ {
		for j := 0; j < len(shape[0]); j++ {
			if shape[i][j] != EMPTY && grid[i+dx][j+dy] != EMPTY {
				return false
			}
		}
	}
	return true
}

func addShapeToGrid(shape Shape, dx, dy int, grid Grid, number uint8) Grid {

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

func removeShapeFromGrid(grid Grid, number uint8) Grid {

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

func elapsed(puzzle Puzzle) func() {
	start := time.Now()

	return func() {
		message := fmt.Sprintf("Found %d solutions in %v.", len(*puzzle.Solutions), RoundedSince(start))
		if puzzle.HasGui {
			puzzle.WinInfo.StatusBar.Push(1, message)
		} else {
			log.Println(message)
		}
	}
}

func RoundedSince(value time.Time) time.Duration {
	return time.Duration(time.Since(value)/time.Millisecond) * time.Millisecond
}

func createEmptyGrid(grid Grid) Grid {

	var copiedGrid = make(Grid, len(grid))
	for i := 0; i < len(grid); i++ {
		copiedGrid[i] = make([]uint8, len(grid[0]))
	}
	return copiedGrid
}

func copyGrid(grid Grid) Grid {

	copiedGrid := createEmptyGrid(grid)

	var i, j int
	for i = 0; i < len(grid); i++ {
		for j = 0; j < len(grid[0]); j++ {
			copiedGrid[i][j] = grid[i][j]
		}
	}
	return copiedGrid
}

func Factorial(n uint64) uint64 {
	var result uint64 = 1
	var i uint64
	for i=2; i<=n; i++ {
		result *= i
	}
	return result
}