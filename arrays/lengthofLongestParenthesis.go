package arrays

func LongestValidParentheses(str string) int {

	stack := make([]int, 0)
	var max, left int = 0, -1

	for i, ch := range str {

		if ch == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				left = i
			} else {
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					max = getMax(max, i-left)
				} else {
					max = getMax(max, i-stack[len(stack)-1])
				}
			}
		}

	}
	return max
}

func getMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}
