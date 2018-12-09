package main

import (
	"fmt"
	"testing"
)

func TestContainerWIthMostWater(t *testing.T) {
	nlist := []int{1, 8, 9, 2, 5, 4, 8, 3, 9}
	retVal := maxArea(nlist)
	if retVal == 56 {
		fmt.Printf("The Container area is %d (7 * 8) \n", retVal)
	}
}
