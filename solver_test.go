package main

import (
	"testing"
)

func TestHasLeftUnfillableAreas(t *testing.T) {

	var grid = [][]uint8{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}

	if !hasLeftUnfillableAreas(grid, 2) {
		t.Errorf("Expected no fillable area.")
	}

	grid = [][]uint8{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}

	if !hasLeftUnfillableAreas(grid, 9) {
		t.Errorf("Expected no fillable area.")
	}

	grid = [][]uint8{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}

	if !hasLeftUnfillableAreas(grid, 3) {
		t.Errorf("Expected no fillable area.")
	}

	grid = [][]uint8{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1},
	}

	if !hasLeftUnfillableAreas(grid, 2) {
		t.Errorf("Expected no fillable area.")
	}

	grid = [][]uint8{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	if !hasLeftUnfillableAreas(grid, 26) {
		t.Errorf("Expected no fillable area.")
	}

	grid = [][]uint8{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}

	if !hasLeftUnfillableAreas(grid, 11) {
		t.Errorf("Expected no fillable area.")
	}

	grid = [][]uint8{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}

	if !hasLeftUnfillableAreas(grid, 100000) {
		t.Errorf("Expected no fillable area.")
	}

}

func TestComputesCorrectSolutions(t *testing.T) {

	var grid = [][]uint8{
		{1, 2, 2, 2},
		{1, 1, 4, 4},
		{1, 1, 4, 4},
		{3, 3, 3, 3},
	}

	pieces := GetPiecesFromGrid(grid)
	var solutions [][][]uint8

	var puzzle = Puzzle{
		pieces,
		copyGrid(grid[:]),
		grid[:],
		4,
		minPieceSize(pieces),
		&solutions,
		true,
		false,
		&WinInfo{},
	}

	Solver(&puzzle)

	if len(solutions) != 16 {
		t.Errorf("Solver must find 16 solutions")
	}
}
