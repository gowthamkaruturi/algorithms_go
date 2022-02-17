package arrays

// func IsValidSquare()
func countofPath(grids [][]bool, row int, col int) int {

	return countofPath(grids, row+1, col) + countofPath(grids, row, col+1)
}
