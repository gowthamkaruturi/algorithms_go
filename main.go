package main

import (
	"fmt"
	"strings"

	"com.gowtham/algorithms_go/stringalgo"
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
	newArr := strings.Split("gowthammmmmm", "")
	fmt.Print(stringalgo.PermutationOfString(newArr))

}
