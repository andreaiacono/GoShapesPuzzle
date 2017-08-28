package main

import (
	"io/ioutil"
	"strings"
	"errors"
)

func ReadFile(filename string) (Puzzle, error) {

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return Puzzle{}, err
	}

	rows := strings.Split(strings.Trim(strings.ToUpper(string(dat[:])), " \t\n\r"), "\n")
	var maxLen int8 = 0
	var grid = make([][]int8, len(rows), len(rows[0]))
	for index, row := range rows {
		if len(row) == 0 {
			continue
		}
		var rowValues = []int8{}
		var counter int8 = 0
		for _, char := range row {
			val := int8(char)
			if val == 32 {
				continue
			}
			counter ++
			if val >= 48 && val <= 57 {
				rowValues = append(rowValues, val-48)
			} else if val >= 65 && val <= 90 {
				rowValues = append(rowValues, val-55)
			} else {
				return Puzzle{}, errors.New("Only numbers and characters allowed in model.")
			}
		}
		if maxLen < counter {
			maxLen = counter
		}
		grid[index] = rowValues
	}

	pieces := GetPiecesFromGrid(grid)
	return Puzzle { pieces, grid[:], max(int8(len(rows)), maxLen)}, nil
}

func max(a int8, b int8) int8 {
	if a >= b {
		return a
	}
	return b
}