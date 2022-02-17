package stringalgo

import "reflect"

func FindRotation(mat [][]int, target [][]int) bool {

	for i := 0; i < 4; i++ {
		if reflect.DeepEqual(mat, target) {
			return true
		}
		rotatematrix(mat)
	}
	return false
}
