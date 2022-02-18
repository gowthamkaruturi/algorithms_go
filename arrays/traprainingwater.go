package arrays

func trap(height []int) int {

	maxleft, maxright, totalWater, left, right := 0, 0, 0, 0, len(height)-1

	for left < right {

		if height[left] < height[right] {

			if height[left] >= maxleft {
				maxleft = height[left]
			} else {
				totalWater += maxleft - height[left]
			}
			left++
		} else {
			if height[right] >= maxright {
				maxright = height[right]
			} else {
				totalWater += maxright - height[right]
			}

		}
	}
	return totalWater
}
