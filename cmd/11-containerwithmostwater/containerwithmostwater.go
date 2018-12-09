package main

import (
	"fmt"
)

func main() {
	nlist := []int{1, 8, 9, 2, 5, 4, 8, 3, 9}
	fmt.Println(maxArea(nlist))
}

func maxArea(height []int) int {
	i, k, area, val := 0, len(height)-1, 0, 0

	for i < k {
		lh, rh := height[i], height[k]
		if lh <= rh {
			val = lh * (k - i)
		} else {
			val = rh * (k - i)
		}
		if val > area {
			area = val
		}

		if lh <= rh {
			for i < k && height[i] <= lh {
				i++
			}
		} else {
			for i < k && height[k] <= rh {
				k = k - 1
			}
		}
	}

	return area
}
