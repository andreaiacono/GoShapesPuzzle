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
