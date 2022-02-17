package stringalgo

func rotatematrix(matrix [][]int) {

	n := len(matrix)

	for i, j := 0, 1; i < n && j < n; i, j = i+1, j+1 {

		temp := matrix[i][j]
		matrix[i][j] = matrix[j][i]
		matrix[j][i] = temp

	}

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[i][n-1-j]
			matrix[i][n-1-j] = temp
		}
	}

}
