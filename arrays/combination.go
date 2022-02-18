package arrays

var result [][]int
var combination []int

func CombinationSum(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	combination = make([]int, 0)
	getResult(candidates, target, 0)
	return result
}

func getResult(candidates []int, target int, index int) {
	if index == len(candidates) || target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(combination))
		copy(temp, combination)
		result = append(result, temp)
		return
	}

	if target > 0 {

		for i := index; i < len(candidates) && candidates[i] <= target; i++ {
			combination = append(combination, candidates[i])
			getResult(candidates, target-candidates[i], i)
			combination = combination[:len(combination)-1]
		}
	}

}

// var ans [][]int
// var stack []int

func search(candidates []int, target int, index int) {
	if target > 0 {
		for index >= 0 {
			combination = append(combination, candidates[index])
			search(candidates, target-candidates[index], index)
			if len(combination) == 1 {
				combination = make([]int, 0)
			} else {
				combination = combination[:len(combination)-1]
			}
			index--
		}
		return
	} else if target == 0 {
		temp := make([]int, len(combination))
		copy(temp, combination)
		result = append(result, temp)
		return
	} else {
		return
	}
}

// func CombinationSum(candidates []int, target int) [][]int {
// 	ans = make([][]int, 0)
// 	stack = make([]int, 0)
// 	search(candidates, target, len(candidates)-1)
// 	return ans
// }
