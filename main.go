package main

import (
	"fmt"

	"com.gowtham/algorithms_go/arrays"
)

// func main() {
// 	newArr := strings.Split("gowthammmmmm", "")

// 	permute(newArr)

// }
func combinationSum(candidates []int, target int) [][]int {
	var dfs func(int, int, []int, [][]int) [][]int
	dfs = func(i, sum int, combination []int, out [][]int) [][]int {
		if sum == target {
			out = append(out, append([]int(nil), combination...))
			return out
		}
		if sum > target {
			return out
		}
		for j := i; j < len(candidates); j++ {
			out = dfs(j, sum+candidates[j], append(combination, candidates[j]), out)
		}
		return out
	}
	return dfs(0, 0, nil, nil)
}
func main() {
	nums := []int{2, 7, 6, 3, 5, 1}
	fmt.Println(arrays.CombinationSum(nums, 9))
	fmt.Println(arrays.RemoveKdigits("12345677", 3))

}
