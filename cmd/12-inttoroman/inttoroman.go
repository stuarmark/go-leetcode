package main

import (
	"fmt"
)

func main() {

	fmt.Println(intToRoman(2999))
}

func intToRoman(num int) string {
	m := [10]string{"", "M", "MM", "MMM"}
	c := [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"} // D 500
	x := [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"} // L 50
	i := [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return fmt.Sprintf("%s%s%s%s", m[num/1000], c[(num%1000)/100], x[(num%1000%100)/10], i[(num%1000%100%10)])
}
