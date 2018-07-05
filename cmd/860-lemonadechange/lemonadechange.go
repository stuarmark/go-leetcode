package main

import (
	"fmt"
)

func main() {
	billsa := []int{5, 5, 5, 10, 5, 5, 10, 20, 20, 20}
	changebool := lemonadeChange(billsa)

	fmt.Println(fmt.Sprintf("%v produce: %t", billsa, changebool))
}

func lemonadeChange(bills []int) bool {
	bill5, bill10 := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			bill5++
		case 10:
			bill10++
			bill5--
		case 20:
			if bill10 > 0 && bill5 > 0 {
				bill10--
				bill5--
			} else {
				bill5 -= 3
			}
		}
		if bill5 < 0 || bill10 < 0 {
			return false
		}

	}
	return true
}

//fmt.Println(fmt.Sprintf("bill: %d bill5: %d  bill10: %d", bill, bill5, bill10))
