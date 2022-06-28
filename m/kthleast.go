package m

import "sort"

func KthLeastElement(elements []int, k int) int {
	kthleastElement := 0

	if k > len(elements) {
		return kthleastElement
	}
	SortTheArrayByAscending(elements)

	return elements[k-1]

}

func SortTheArrayByAscending(elements []int) {
	sort.Slice(elements, func(i, j int) bool {
		return elements[i] > elements[j]
	})

}
