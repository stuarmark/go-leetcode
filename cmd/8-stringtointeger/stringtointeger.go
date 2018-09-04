package main

import (
	"fmt"
	"math"
)

func main() {

	//
	// "   +0 123"
	// "  -+12344"
	// "  +1"
	// "   +-2"
	// "9223372036854775808"  ->2147483647

	fmt.Println(myAtoi("-1233"))
}

func myAtoi(str string) int {
	sign, num, signcount, charcount := 1, 0, 0, 0
	if str == "" {
		return 0
	}

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			if charcount > 0 {
				break
			}
			continue
		} else if str[i] == '+' {
			if signcount > 0 {
				break
			}
			signcount++
			charcount++
			continue
		} else if str[i] == '-' {
			if signcount == 0 {
				sign = -1
				signcount++
				charcount++
			} else {
				break
			}
			continue
		}
		if str[i] < 48 || str[i] > 57 {
			break
		}
		num = num*10 + int(str[i]-'0')
		charcount++
		fmt.Println(num)
	}
	num = sign * num

	if num > math.MaxInt32 {
		return math.MaxInt32
	} else if num < math.MinInt32 {
		return math.MinInt32
	}

	return num
}
