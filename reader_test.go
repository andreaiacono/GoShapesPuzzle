package main

import (
	"testing"
)

func TestCreatePuzzle(t *testing.T) {

	model := "1 2 2 2\n1 1 4 4\n1 1 4 4\n3 3 3 3"
	puzzle, err := createPuzzle(model)
	if err != nil {
		t.Errorf("There was an error creating the puzzle: %s", err)
	}

	if len(puzzle.Pieces) != 4 {
		t.Errorf("The number of pieces is 4")
	}

	if puzzle.MaxPieceSide != 4 {
		t.Errorf("The max side of a piece is 4")
	}

	if puzzle.MinPieceSize != 3 {
		t.Errorf("The min size of a piece is 3")
	}

	if len(puzzle.WorkingGrid) != 4 || len(puzzle.WorkingGrid[0]) != 4 {
		t.Errorf("The size of the grid is 4x4")
	}
}

func TestWrongModel(t *testing.T) {

	model := "1 1\n1 0"
	_, err := createPuzzle(model)
	if err == nil {
		t.Errorf("The '0' character is not allowed.")
	}

	model = "1 1\n1 0 1"
	_, err = createPuzzle(model)
	if err == nil {
		t.Errorf("The model must have a rectangular shape")
	}

	model = "1 1\n1 $ 1"
	_, err = createPuzzle(model)
	if err == nil {
		t.Errorf("The model can contain only latin characters and numbers [1-9]")
	}
}
