package arrays

func LongestPalindrome(s string) string {
	resultLength, resultStart := 0, 0
	for i, _ := range s {
		len1 := ExpandRange(s, i, i)
		len2 := ExpandRange(s, i, i+1)
		length := getMax(len1, len2)
		if length > resultStart-resultLength {
			resultStart = i - (length-1)/2
			resultLength = i + length/2
		}

	}
	return s[resultStart : resultStart+1]

}

func ExpandRange(str string, start int, end int) int {
	l, r := start, end
	for l >= 0 && r < len(str) && str[l] == str[r] {
		l--
		r++
	}
	return r - l - 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
