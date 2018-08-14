package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	fmt.Println(reverse(12999999))
}

func reverse(x int) int {
	str := strconv.Itoa(x)
	var arr []byte
	strLen := len(str)
	stop := 0

	if str[0] == '-' {
		arr = append([]byte("-"), arr...)
		stop = 1
	}

	for i := strLen - 1; i >= stop; i-- {
		arr = append(arr, byte(str[i]))
	}
	number, err := strconv.Atoi(string(arr))
	if err != nil {
		panic(err)
	}

	if number > math.MaxInt32 || number < math.MinInt32 {
		return 0
	}

	return number
}
