package stringalgo

func PermutationOfString(s []string) []string {

	m := make(map[string]int)

	for _, ch := range s {
		val, ok := m[ch]
		if ok {
			m[ch] = val + 1
		} else {
			m[ch] = 1
		}
	}
	charArray := make([]string, len(m))
	lengthArray := make([]int, len(m))
	index := 0
	for char, count := range m {
		charArray[index] = char
		lengthArray[index] = count
		index++
	}
	resultArray := make([]string, len(s))
	var returnArray []string
	return PermutatationUtil(resultArray, charArray, lengthArray, returnArray, 0)
}

func PermutatationUtil(resultArray []string, charArray []string, lengthArray []int, returnArray []string, level int) []string {

	if level == len(resultArray) {

		returnArray = append(returnArray, resultArray...)
	}

	for i := 0; i < len(charArray); i++ {
		if lengthArray[i] == 0 {
			continue
		}
		resultArray[level] = charArray[i]
		lengthArray[i]--
		PermutatationUtil(resultArray, charArray, lengthArray, returnArray, level+1)
		lengthArray[i]++
	}
	return returnArray
}
