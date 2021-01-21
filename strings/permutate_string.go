package main

import (
	"fmt"
	"strings"
)

func permute(strs []string) {
	m := make(map[string]int)

	for _, ch := range strs {
		val, ok := m[ch]
		if ok {
			m[ch] = m[ch] + 1
		} else {
			fmt.Println("i am here")
			m[ch] = 1

		}
	}

	charArray := make([]string, len(m))
	lengthArray := make([]int, len(m))
	index := 0
	for k, v := range m {
		fmt.Println("key and values are ", k, v)
		charArray[index] = k
		lengthArray[index] = v
		index++
	}
	resultArray := make([]string, len(strs))
	permuteUtil(charArray, lengthArray, resultArray, 0)

}

func permuteUtil(str []string, count []int, result []string, level int) {
	if level == len(result) {
		fmt.Println(result)
	}

	for i := 0; i < len(str); i++ {
		if count[i] == 0 {
			continue
		}
		result[level] = str[i]
		count[i]--
		permuteUtil(str, count, result, level+1)
		count[i]++

	}
}
func main() {
	newArr := strings.Split("gowthammmmmm", "")
	
	permute(newArr)

}
