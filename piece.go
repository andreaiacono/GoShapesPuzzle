package main

type Shape [][]int8

type Piece struct {
	Shape     [][]int8
	Number    int8
	Rotations []Shape
}

func copyPiece(piece Piece) Piece {
	return Piece{piece.Shape, piece.Number, piece.Rotations}
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
		rotatedShape[ii] = make([]int8, n)
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
		copiedShape[i] = make([]int8, len(shape[0]))
	}

	var i, j int
	for i = 0; i < len(shape); i++ {
		for j = 0; j < len(shape[0]); j++ {
			copiedShape[i][j] = shape[i][j]
		}
	}
	return copiedShape
}

func GetPiecesFromGrid(grid [][]int8) []Piece {
	pieces := []Piece{}
	values := getValuesFromGrid(grid)

	var i int
	for i = 0; i < len(values); i++ {
		pieces = append(pieces, getPiece(grid, int8(values[i])))
	}

	return pieces
}

func getPiece(grid [][] int8, pieceNumber int8) Piece {

	var minX = 1000
	var minY = 1000
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
	//log.Printf("Piece n. %d, maxX:%d maxY:%d, minX:%d, minY:%d", pieceNumber, maxX, maxY, minX, minY)

	pieceGrid := make([][]int8, maxX-minX+1)
	for i := range pieceGrid {
		pieceGrid[i] = make([]int8, maxY-minY+1)
	}

	for i = minX; i <= maxX; i++ {
		for j = minY; j <= maxY; j++ {
			if grid[i][j] == pieceNumber {
				pieceGrid[i-minX][j-minY] = pieceNumber
			}
		}
	}

	//log.Printf("Piece %d: %v", pieceNumber, pieceGrid)
	return Piece{pieceGrid, pieceNumber, getRotations(pieceGrid)}
}

func getRotations(piece Shape) []Shape {

	// there's the piece itself, the 3 90-degrees rotation of the piece plus the flipped
	// piece and its 3 90-degrees rotation: 8 in total
	var rotations [8]Shape
	rotations[0] = copyShape(piece)

	var i int
	for i = 0; i < 3; i++ {
		piece = rotate(piece)
		rotations[i+1] = copyShape(piece)
	}
	piece = flip(piece)
	rotations[4] = piece
	for i = 0; i < 3; i++ {
		piece = rotate(piece)
		rotations[i+5] = copyShape(piece)
	}

	//FIXME remove symmetries!
	return rotations[:]
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
