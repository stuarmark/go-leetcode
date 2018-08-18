package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(myAtoi("-   234"))
}

func myAtoi(str string) int {
	sign, num, signcount, numcount := 1, 0, 0, 0

	if str == "" {
		return 0
	}

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			if numcount > 0 || signcount > 0 {
				break
			}
			continue
		} else if str[i] == '+' {
			if signcount > 0 {
				break
			}
			signcount++
			continue
		} else if str[i] == '-' {
			if signcount == 0 && numcount == 0 {
				sign = -1
				signcount++
			} else if numcount > 0 || signcount > 0 {
				break
			}

			continue
		}
		if str[i] < 48 || str[i] > 57 {
			break
		}

		num = num*10 + int(str[i]-'0')
		numcount++

		if num > math.MaxInt32 && sign == -1 {
			return math.MinInt32
		} else if num > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return sign * num
}
