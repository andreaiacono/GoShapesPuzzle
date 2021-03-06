package main

import (
	"fmt"
)

type Shape [][]uint8

type Piece struct {
	Shape     Shape
	Number    uint8
	Rotations []Shape
	Size 	  int
}

func copyPiece(piece Piece) Piece {
	return Piece{piece.Shape, piece.Number, piece.Rotations, piece.Size}
}

func (piece Piece) Flip() Piece {
	var flipped = copyPiece(piece)
	flipped.Shape = flip(piece.Shape)
	return flipped
}

func flip(shape Shape) Shape {
	var flipped Shape = copyShape(shape)
	var cols = len(shape[0]) - 1
	var i, j int
	for i = 0; i < len(shape); i++ {
		for j = 0; j < len(shape[0]); j++ {
			flipped[i][j] = shape[i][cols-j]
		}
	}
	return flipped
}

func (piece Piece) Rotate() Piece {
	var rotated = copyPiece(piece)
	rotated.Shape = rotate(piece.Shape)
	return rotated
}

func rotate(shape Shape) Shape {

	var n = len(shape)
	var m = len(shape[0])
	var rotatedShape Shape = make(Shape, m)
	for ii := 0; ii < m; ii++ {
		rotatedShape[ii] = make([]uint8, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			rotatedShape[m-j-1][i] = shape[i][j]
		}
	}
	return rotatedShape
}

func copyShape(shape Shape) Shape {
	var copiedShape = make(Shape, len(shape))
	for i := 0; i < len(shape); i++ {
		copiedShape[i] = make([]uint8, len(shape[0]))
	}

	var i, j int
	for i = 0; i < len(shape); i++ {
		for j = 0; j < len(shape[0]); j++ {
			copiedShape[i][j] = shape[i][j]
		}
	}
	return copiedShape
}

func GetPiecesFromGrid(grid Grid) []Piece {
	pieces := []Piece{}
	values := getValuesFromGrid(grid)

	var i int
	for i = 0; i < len(values); i++ {
		pieces = append(pieces, getPiece(grid, uint8(values[i])))
	}

	return pieces
}

func getPiece(grid Grid, pieceNumber uint8) Piece {

	var minX = 1000
	var minY = 1000
	var maxX = 0
	var maxY = 0
	var i, j int
	var size int = 0

	for i = 0; i < len(grid); i++ {
		for j = 0; j < len(grid[i]); j++ {
			if grid[i][j] == pieceNumber {
				size ++
				if i < minX {
					minX = i
				}
				if i > maxX {
					maxX = i
				}
				if j < minY {
					minY = j
				}
				if j > maxY {
					maxY = j
				}
			}
		}
	}

	pieceGrid := make(Shape, maxX-minX+1)
	for i := range pieceGrid {
		pieceGrid[i] = make([]uint8, maxY-minY+1)
	}

	for i = minX; i <= maxX; i++ {
		for j = minY; j <= maxY; j++ {
			if grid[i][j] == pieceNumber {
				pieceGrid[i-minX][j-minY] = pieceNumber
			}
		}
	}
	return Piece{pieceGrid, pieceNumber, getRotations(pieceGrid), size}
}

// there's the piece itself, the three 90-degrees rotation of the piece plus the flipped
// piece and its three 90-degrees rotation: 8 in total. Equal rotations are excluded in
// case of some kind of symmetry
func getRotations(piece Shape) []Shape {

	var rotations [8]Shape
	rotations[0] = copyShape(piece)
	var count = 1

	var i int
	for i = 0; i < 3; i++ {
		piece = rotate(piece)
		if ! containsShape(rotations[:], piece) {
			rotations[count] = copyShape(piece)
			count ++
		}
	}
	piece = flip(piece)
	if !containsShape(rotations[:], piece) {
		rotations[count] = piece
		count ++
	}
	for i = 0; i < 3; i++ {
		piece = rotate(piece)
		if ! containsShape(rotations[:], piece) {
			rotations[count] = copyShape(piece)
			count ++
		}
	}

	return rotations[0:count]
}

func containsShape(shapes []Shape, newShape Shape) bool {
	for _, shape := range shapes {
		if areEqualPieces(shape, newShape) {
			return true
		}
	}
	return false
}

func areEqualPieces(shape1 [][]uint8, shape2 [][]uint8) bool {
	if len(shape1) != len(shape2) || len(shape1[0]) != len(shape2[0]) {
		return false
	}

	for i := 0; i < len(shape1); i++ {
		for j := 0; j < len(shape1[0]); j++ {
			if shape1[i][j] != shape2[i][j] {
				return false
			}
		}

	}
	return true
}

func getValuesFromGrid(grid Grid) []uint8 {

	var i, j int
	var values = []uint8{}

	for i = 0; i < len(grid); i++ {
		for j = 0; j < len(grid[i]); j++ {
			if !contains(values, grid[i][j]) {
				values = append(values, grid[i][j])
			}
		}
	}
	return values
}

func contains(values []uint8, value uint8) bool {
	var i int
	for i = 0; i < len(values); i++ {
		if values[i] == value {
			return true
		}
	}
	return false
}

func (piece Piece) String() string {
	return fmt.Sprintf("%d", piece.Number)
}

