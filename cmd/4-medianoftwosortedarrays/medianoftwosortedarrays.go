package main

import (
	"fmt"
)

func main() {
	// nums1 := []int{1, 2}
	// nums2 := []int{3, 4}
	nums1 := []int{1, 3}
	nums2 := []int{2}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lena := len(nums1)
	lenb := len(nums2)

	lent := lena + lenb
	max, min, ia, ib, it := 0, 0, 0, 0, 0

	for it <= lent/2 && ia < lena && ib < lenb {
		if nums1[ia] < nums2[ib] {
			max, min = min, nums1[ia]
			ia++
		} else {
			max, min = min, nums2[ib]
			ib++
		}
		it++
	}

	for it <= lent/2 && ia < lena {
		max, min = min, nums1[ia]
		ia++
		it++
	}

	for it <= lent/2 && ib < lenb {
		max, min = min, nums2[ib]
		ib++
		it++
	}

	if lent%2 == 0 {
		return float64(min+max) / 2.0
	}

	return float64(min)
}
