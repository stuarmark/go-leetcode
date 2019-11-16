package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{6, 1, 4, 3}, 9))
}

func twoSum(nums []int, target int) []int {
	numsmap := make(map[int]int)
	for i, v := range nums {
		if j, k := numsmap[target-v]; k {
			return []int{j, i}
		}
		numsmap[v] = i
	}
	return nil
}
