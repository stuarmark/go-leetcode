package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(myAtoi("-3219"))
}

func myAtoi(str string) int {
	sign, num := 1, 0
	if str == "" {
		return 0
	}

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		} else if str[i] == '-' {
			sign = -1
			continue
		}
		if str[i] < 48 || str[i] > 57 {
			break
		}
		num = num*10 + int(str[i]-'0')

	}
	num = sign * num

	if num > math.MaxInt32 {
		return math.MaxInt32
	} else if num < math.MinInt32 {
		return math.MinInt32
	}

	return num
}
