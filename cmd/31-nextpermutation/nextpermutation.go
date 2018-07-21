package main

import (
	"sort"
)

func main() {
	// 	1,2,3 → 1,3,2
	// 	3,2,1 → 1,2,3
	// 	1,1,5 → 1,5,1
	numarray := []int{3, 2, 1}
	nextPermutation(numarray)
}

func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}
	index := -1

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			index = i
			break
		}
	}

	if index == -1 {
		sort.Ints(nums)
		return
	}

	index2 := index
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[index2] {
			index = i
			break
		}
	}
	nums[index2], nums[index] = nums[index], nums[index2]

	s1, s2 := index2+1, len(nums)-1

	for s1 < s2 {
		nums[s1], nums[s2] = nums[s2], nums[s1]
		s1, s2 = s1+1, s2-1
	}

	return
}
