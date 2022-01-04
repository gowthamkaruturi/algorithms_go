package arrays

import "fmt"

func Permutate(str string, l int, r int) {

	if l == r {
		fmt.Println(str)
	} else {

		for i := l; i < r; i++ {

			str = swap(str, l, i)
			Permutate(str, l+1, r)
			str = swap(str, l, i)
		}

	}

}

func swap(str string, i int, j int) string {

	chars := []rune(str)

	temp := chars[i]
	chars[i] = chars[j]
	chars[j] = temp

	return string(chars)

}
