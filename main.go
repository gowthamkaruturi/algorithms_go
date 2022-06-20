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
	// nums := []int{2, 7, 6, 3, 5, 1}
	// fmt.Println(arrays.CombinationSum(nums, 9))
	// fmt.Println(arrays.RemoveKdigits("12345677", 3))
	// fmt.Println(arrays.LongestPalindrome("amama"))

	// g := graphs.InitGraph(4)
	// g.AddEdge(0, 1)
	// g.AddEdge(0, 2)
	// g.AddEdge(1, 3)
	// g.AddEdge(2, 3)
	// g.PrintGraph()
	// s := set.NewSet()

	// s.Add("gowtham")
	// s.Add(1)
	// s.Add(2)
	// s.Add("gowtham")
	// s.Print()

	fmt.Println(arrays.WordSort("i am an indian i need some help with coding."))

}
