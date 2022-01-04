package arrays

func Unique_char_string(str string) bool {
	var char_set [256]bool
	chars := []rune(str)
	for i := 0; i < len(chars)-1; i++ {
		if char_set[chars[i]] {
			return false
		}
		char_set[chars[i]] = true
	}

	return true
}
