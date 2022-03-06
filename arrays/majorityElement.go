package arrays

func majorityElement(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}
	occurencesCount := len(nums) / 3
	result := make([]int, 0)
	majorityElement := make(map[int]int)

	for _, num := range nums {
		if count, ok := majorityElement[num]; ok {
			majorityElement[num] = count + 1
		} else {
			majorityElement[num] = 1
		}
	}

	for num, count := range majorityElement {
		if count > occurencesCount {
			result = append(result, num)
		}
	}

	return result
}
