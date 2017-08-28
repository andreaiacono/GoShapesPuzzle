package main


type Piece struct {
	Shape  [][] int8
	Number int8
	MaxX   int
	MaxY   int
}

func GetPiecesFromGrid(grid [][] int8) []Piece {
	pieces := []Piece{}
	values := getValuesFromGrid(grid)

	var i int
	for i = 0; i < len(values); i++ {
		pieces = append(pieces, getPiece(grid, int8(values[i])))
	}

	return pieces
}

func getPiece(grid [][] int8, pieceNumber int8) Piece {

	var minX = 0
	var minY = 0
	var maxX = 0
	var maxY = 0
	var i, j int

	for i = 0; i < len(grid); i++ {
		for j = 0; j < len(grid[i]); j++ {
			if grid[i][j] == pieceNumber {
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

	pieceGrid := make([][]int8, maxX-minX)
	for i := range pieceGrid {
		pieceGrid[i] = make([]int8, maxY-minY)
	}

	for i = minX; i < maxX; i++ {
		for j = minY; j < maxY; j++ {
			pieceGrid[i-minX][j-minY] = pieceNumber
		}
	}
	return Piece{pieceGrid, pieceNumber, maxX, maxY}
}

func getValuesFromGrid(grid [][]int8) []int8 {

	var i, j int
	var values = []int8{}

	for i = 0; i < len(grid); i++ {
		for j = 0; j < len(grid[i]); j++ {
			if !contains(values, grid[i][j]) {
				values = append(values, grid[i][j])
			}
		}
	}
	return values
}

func contains(values []int8, value int8) bool {
	var i int
	for i = 0; i < len(values); i++ {
		if values[i] == value {
			return true
		}
	}
	return false
}
