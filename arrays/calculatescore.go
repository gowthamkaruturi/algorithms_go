package arrays

// Your last Plain Text code is saved below:
// Calculate a final score from a set of round scores

// In each round, the score is either a number, X, or S
// * Number - count as part of the final score
// * X - remove the previous round's score
// * S - worth the sum of the scores from the previous two rounds

// Examples:

// [1, 2, 4] => 7

// [1, 2, 3, X] => 3

// [1, 2, 4, S] => 13

func getSumAtLocation(arr []int64, loc int) int64 {
	var value int64 = 0
	if loc >= 0 {
		if len(arr) > loc {
			value = arr[loc]
		}
	}
	return value
}

func RemoveElemetInIndex(s []int64, index int) []int64 {
	elements := make([]int64, 0)
	if index >= 1 {
		elements = append(s[:index], s[index+1:]...)
	}
	return elements
}

func CalcuteTheScore(arr []interface{}) int64 {

	sumOfRounds := make([]int64, 0)
	for index, val := range arr {
		l := len(sumOfRounds)
		switch t := val.(type) {
		case int:
			sumOfRounds = append(sumOfRounds, getSumAtLocation(sumOfRounds, index-1)+int64(t))
		case int64:
			sumOfRounds = append(sumOfRounds, getSumAtLocation(sumOfRounds, index-1)+t)
		case string:

			if t == "x" {
				sumOfRounds = RemoveElemetInIndex(sumOfRounds, len(sumOfRounds)-1)
			}
			if t == "s" {
				if l >= 1 {
					sumOfRounds = append(sumOfRounds, sumOfRounds[len(sumOfRounds)-1]*2)
				}

			}
		}

	}

	return sumOfRounds[len(sumOfRounds)-1]
}

// func main() {
// 	arr1 := []interface{}{1, 2, 3}
// 	arr2 := []interface{}{1, 2, 3, "x"}
// 	arr3 := []interface{}{1, 2, 3, 4, "s", "s", "x"}
// 	arr4 := []interface{}{1, "s", "x", 4, "s", "s", "x"}
// 	arr5 := []interface{}{"x", "s", "x", 4, "s", "s", "x"}
// 	fmt.Println(CalcuteTheScore(arr1))
// 	fmt.Println(CalcuteTheScore(arr2))
// 	fmt.Println(CalcuteTheScore(arr3))
// 	fmt.Println(CalcuteTheScore(arr4))
// 	fmt.Println(CalcuteTheScore(arr5))
// }
