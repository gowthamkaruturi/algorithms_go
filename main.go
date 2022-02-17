package main

import (
	"fmt"

	"com.gowtham/algorithms_go/arrays"
)

// func main() {
// 	newArr := strings.Split("gowthammmmmm", "")

// 	permute(newArr)

// }
func main() {
	// fmt.Println(arrays.Fiboatn(65))
	// var a [5]int
	// fmt.Println(arrays.Fibo(4, a[0:5]))
	// profits := []int{1, 6, 10, 16}
	// weights := []int{1, 2, 3, 5}
	// fmt.Println(knapsack.Knapcsacksolve(profits, weights, 7))
	// newArr := strings.Split("abc", "")
	// fmt.Print(stringalgo.PermutationOfString(newArr))

	// input := [][]int{{1, 0}, {0, 1}}
	// target := [][]int{{1, 1}, {1, 0}}
	// fmt.Println(stringalgo.FindRotation(input, target))

	// arr := []int{4, 4, 6, 6, 8}

	// fmt.Println(arrays.FirstNonRepeating(arr))

	// tree := &trees.BinaryTree{}
	// tree.Insert(100).
	// 	Insert(-20).
	// 	Insert(-50).
	// 	Insert(-15).
	// 	Insert(-60).
	// 	Insert(50).
	// 	Insert(60).
	// 	Insert(55).
	// 	Insert(85).
	// 	Insert(15).
	// 	Insert(5).
	// 	Insert(-10)

	// trees.PrintInOrder(tree.Root)
	// // arr := []int32{-4, 3, -9, 0, 4, 1}
	// // arrays.PlusMinus(arr)

	// fmt.Println(trees.KThMaxValue(tree.Root, 9))
	arr1 := []interface{}{1, 2, 3}
	arr2 := []interface{}{1, 2, 3, "x"}
	arr3 := []interface{}{1, 2, 3, 4, "s", "s", "x"}
	arr4 := []interface{}{1, "s", "x", 4, "s", "s", "x"}
	arr5 := []interface{}{"x", "s", "x", 4, "s", "s", "x"}
	fmt.Println(arrays.CalcuteTheScore(arr1))
	fmt.Println(arrays.CalcuteTheScore(arr2))
	fmt.Println(arrays.CalcuteTheScore(arr3))
	fmt.Println(arrays.CalcuteTheScore(arr4))
	fmt.Println(arrays.CalcuteTheScore(arr5))

}
